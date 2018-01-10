// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package marc21 reads/parses MARC21 and MARCXML data.
//
// Enables converting between MARC21 and MARCXML data.
package marc21

import (
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
)

/*
   MARC-8 vs. UTF-8 encoding
       * leader.CharacterCodingScheme == "a" is UCS/Unicode
       * https://www.loc.gov/marc/specifications/speccharucs.html
       * https://www.loc.gov/marc/specifications/codetables.xml

*/

/*
https://www.loc.gov/marc/specifications/specrecstruc.html
*/

const (
	Delimiter        = 0x1f
	FieldTerminator  = 0x1e
	RecordTerminator = 0x1d
	LeaderLen        = 24
	MaxRecordSize    = 99999
)

type Collection struct {
	Name    xml.Name  `xml:"collection"`
	Records []*Record `xml:"record"`
}

type Leader struct {
	Text string `xml:",chardata"`
}

type Record struct {
	Leader        Leader          `xml:"leader"`
	Controlfields []*Controlfield `xml:"controlfield"`
	Datafields    []*Datafield    `xml:"datafield"`
}

type Directory struct {
	Tag         string
	StartingPos int
	FieldLength int
}

type Controlfield struct {
	Tag  string `xml:"tag,attr"`
	Text string `xml:",chardata"`
}

type Datafield struct {
	Tag       string      `xml:"tag,attr"`
	Ind1      string      `xml:"ind1,attr"`
	Ind2      string      `xml:"ind2,attr"`
	Subfields []*Subfield `xml:"subfield"`
}

type Subfield struct {
	Code string `xml:"code,attr"`
	Text string `xml:",chardata"`
}

func ParseNextRecord(r io.Reader) (rec *Record, err error) {

	rawRec, err := NextRecord(r)
	if err != nil {
		return nil, err
	}

	rec, err = ParseRecord(rawRec)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

func NextRecord(r io.Reader) (rawRec []byte, err error) {

	// Read the first 5 bytes, determine the record length and
	//    read the remainder of the record
	rawLen := make([]byte, 5)
	_, err = r.Read(rawLen)
	if err != nil {
		return nil, err
	}

	recLen, err := strconv.Atoi(string(rawLen[0:5]))
	if err != nil {
		return nil, err
	}

	// Ensure that we have a "sane" record length?
	if recLen <= LeaderLen {
		err = errors.New("MARC record is too short")
		return nil, err
	} else if recLen > MaxRecordSize {
		err = errors.New("MARC record is too long")
		return nil, err
	}

	rawRec = make([]byte, recLen)
	// ensure that the raw len is available for the leader
	copy(rawRec, rawLen)

	// Read the remainder of the record
	_, err = r.Read(rawRec[5:recLen])
	if err != nil {
		return nil, err
	}

	// The last byte should be a record terminator
	if rawRec[len(rawRec)-1] != RecordTerminator {
		return nil, errors.New("Record terminator not found at end of record")
	}

	return rawRec, nil
}

func ParseRecord(rawRec []byte) (rec *Record, err error) {

	rec = new(Record)

	rec.Leader.Text = string(rawRec[:24])

	dir, err := parseDirectory(rawRec)
	if err != nil {
		return nil, err
	}

	baseDataAddress, err := strconv.Atoi(string(rawRec[12:17]))
	if err != nil {
		return nil, err
	}

	rec.Controlfields, err = extractControlfields(rawRec, baseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	rec.Datafields, err = extractDatafields(rawRec, baseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (rec Record) String() string {

	ret := fmt.Sprintf("LDR %s\n", rec.Leader.Text)
	for _, cf := range rec.Controlfields {
		ret += fmt.Sprintf("%s     %s\n", cf.Tag, cf.Text)
	}
	for _, df := range rec.Datafields {
		pre := fmt.Sprintf("%s %s%s _", df.Tag, df.Ind1, df.Ind2)
		for _, sf := range df.Subfields {
			ret += fmt.Sprintf("%s%s%s\n", pre, sf.Code, sf.Text)
			pre = "       _"
		}
	}
	return ret
}

// RecordAsMARC converts a Record into a MARC record byte array
func (rec Record) RecordAsMARC() (marc []byte, err error) {

	if rec.Leader.Text == "" {
		err = errors.New("Record Leader is undefined.")
		return marc, err
	}

	dl := termToByte(Delimiter)
	ft := termToByte(FieldTerminator)
	rt := termToByte(RecordTerminator)

	var Dir []Directory
	var dir []byte
	var cfs []byte
	var dfs []byte
	var startPos int

	// Pack the control fields
	for _, cf := range rec.Controlfields {

		if cf.Text == "" {
			continue
		}

		b := []byte(cf.Text)
		b = append(b, ft...)
		cfs = append(cfs, b...)

		Dir = append(Dir, Directory{Tag: cf.Tag, StartingPos: startPos, FieldLength: len(b)})

		startPos += len(b)
	}

	// Pack the data fields/sub-fields
	for _, df := range rec.Datafields {

		ind1 := df.Ind1
		if ind1 == "" {
			ind1 = " "
		}
		ind2 := df.Ind2
		if ind2 == "" {
			ind2 = " "
		}

		b := []byte(ind1)
		b = append(b, []byte(ind2)...)

		for _, sf := range df.Subfields {
			b = append(b, dl...)
			b = append(b, []byte(sf.Code)...)
			b = append(b, []byte(sf.Text)...)
		}
		b = append(b, ft...)
		dfs = append(dfs, b...)

		Dir = append(Dir, Directory{Tag: df.Tag, StartingPos: startPos, FieldLength: len(b)})

		startPos += len(b)
	}

	// Generate the directory
	for _, d := range Dir {
		dir = append(dir, []byte(d.Tag)...)
		dir = append(dir, []byte(fmt.Sprintf("%04d", d.FieldLength))...)
		dir = append(dir, []byte(fmt.Sprintf("%05d", d.StartingPos))...)
	}
	dir = append(dir, ft...)

	// Build the leader
	recLen := []byte(fmt.Sprintf("%05d", 24+len(dir)+len(cfs)+len(dfs)+1))
	recBaseDataAddress := []byte(fmt.Sprintf("%05d", 24+len(dir)))

	ldr := []byte(rec.Leader.Text)
	for i := 0; i <= 4; i++ {
		ldr[i] = recLen[i]
		ldr[i+12] = recBaseDataAddress[i]
	}

	// Final assembly
	marc = append(marc, ldr...)
	marc = append(marc, dir...)
	marc = append(marc, cfs...)
	marc = append(marc, dfs...)
	marc = append(marc, rt...)

	return marc, nil
}

func termToByte(i int) (b []byte) {

	// We apparently need a 2-byte array for setting the uint16...
	x := make([]byte, 2)
	binary.LittleEndian.PutUint16(x, uint16(i))
	// ... but we don't want the trailing null byte in actual use
	b = x[:1]

	return b
}

// Field returns datafields for the record that match the specified tags
func (rec Record) Field(tags []string) (f []*Datafield) {
	for _, t := range tags {
		for _, d := range rec.Datafields {
			if d.Tag == t {
				f = append(f, d)
			}
		}
	}
	return f
}

// Subfield returns subfields for the datafield that match the specified codes
func (d Datafield) Subfield(codes []string) (sf []*Subfield) {
	for _, c := range codes {
		for _, s := range d.Subfields {
			if s.Code == c {
				sf = append(sf, s)
			}
		}
	}
	return sf
}
