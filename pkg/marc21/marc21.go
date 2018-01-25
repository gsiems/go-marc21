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
	delimiter        = 0x1f
	fieldTerminator  = 0x1e
	recordTerminator = 0x1d
	leaderLen        = 24
	maxRecordSize    = 99999
)

// Collection is for containing zero or more MARC records
type Collection struct {
	Name    xml.Name  `xml:"collection"`
	Records []*Record `xml:"record"`
}

// Leader is for containing the text string of the MARC record Leader
type Leader struct {
	Text string `xml:",chardata"`
}

// Record is for containing a MARC record
type Record struct {
	Leader        Leader          `xml:"leader"`
	Controlfields []*Controlfield `xml:"controlfield"`
	Datafields    []*Datafield    `xml:"datafield"`
}

// directoryEntry contains a single directory entry
type directoryEntry struct {
	tag         string
	startingPos int
	fieldLength int
}

// Controlfield contains a controlfield entry
type Controlfield struct {
	Tag  string `xml:"tag,attr"`
	Text string `xml:",chardata"`
}

// Datafield contains a datafield entry
type Datafield struct {
	Tag       string      `xml:"tag,attr"`
	Ind1      string      `xml:"ind1,attr"`
	Ind2      string      `xml:"ind2,attr"`
	Subfields []*Subfield `xml:"subfield"`
}

// Subfield contains a subfield entry
type Subfield struct {
	Code string `xml:"code,attr"`
	Text string `xml:",chardata"`
}

// ParseNextRecord reads the next MARC record and returns the parsed
// record stucture
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

// NextRecord reads the next MARC record and returns the unparsed bytes
func NextRecord(r io.Reader) (rawRec []byte, err error) {

	// Read the first 5 bytes, determine the record length and
	//    read the remainder of the record
	rawLen := make([]byte, 5)
	_, err = r.Read(rawLen)
	if err != nil {
		return nil, err
	}

	recLen, err := toInt(rawLen[0:5])
	if err != nil {
		return nil, err
	}

	// Ensure that we have a "sane" record length?
	if recLen <= leaderLen {
		err = errors.New("MARC record is too short")
		return nil, err
	} else if recLen > maxRecordSize {
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
	if rawRec[len(rawRec)-1] != recordTerminator {
		return nil, errors.New("Record terminator not found at end of record")
	}

	return rawRec, nil
}

// ParseRecord takes the bytes for a MARC record and returns the parsed
// record stucture
func ParseRecord(rawRec []byte) (rec *Record, err error) {

	rec = new(Record)

	rec.Leader.Text = string(rawRec[:24])

	dir, err := parseDirectory(rawRec)
	if err != nil {
		return nil, err
	}

	baseDataAddress, err := toInt(rawRec[12:17])
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
		ret += fmt.Sprintf("%s     %s\n", cf.GetTag(), cf.GetText())
	}
	for _, df := range rec.Datafields {
		pre := fmt.Sprintf("%s %s%s _", df.GetTag(), df.GetInd1(), df.GetInd2())
		for _, sf := range df.Subfields {
			ret += fmt.Sprintf("%s%s%s\n", pre, sf.GetCode(), sf.GetText())
			pre = "       _"
		}
	}
	return ret
}

// RecordAsMARC converts a Record into a MARC record byte array
func (rec Record) RecordAsMARC() (marc []byte, err error) {

	if rec.Leader.Text == "" {
		err = errors.New("record Leader is undefined")
		return marc, err
	}

	dl := termAsByte(delimiter)
	ft := termAsByte(fieldTerminator)
	rt := termAsByte(recordTerminator)

	var dir []directoryEntry
	var rawDir []byte
	var cfs []byte
	var dfs []byte
	var startPos int

	// Pack the control fields
	for _, cf := range rec.Controlfields {

		if cf.GetText() == "" {
			continue
		}

		b := []byte(cf.GetText())
		b = append(b, ft...)
		cfs = append(cfs, b...)

		dir = append(dir, directoryEntry{tag: cf.GetTag(), startingPos: startPos, fieldLength: len(b)})

		startPos += len(b)
	}

	// Pack the data fields/sub-fields
	for _, df := range rec.Datafields {

		ind1 := df.GetInd1()
		ind2 := df.GetInd2()

		b := []byte(ind1)
		b = append(b, []byte(ind2)...)

		for _, sf := range df.Subfields {
			b = append(b, dl...)
			b = append(b, []byte(sf.GetCode())...)
			b = append(b, []byte(sf.GetText())...)
		}
		b = append(b, ft...)
		dfs = append(dfs, b...)

		dir = append(dir, directoryEntry{tag: df.GetTag(), startingPos: startPos, fieldLength: len(b)})

		startPos += len(b)
	}

	// Generate the directory
	for _, de := range dir {
		rawDir = append(rawDir, []byte(de.tag)...)
		rawDir = append(rawDir, []byte(fmt.Sprintf("%04d", de.fieldLength))...)
		rawDir = append(rawDir, []byte(fmt.Sprintf("%05d", de.startingPos))...)
	}
	rawDir = append(rawDir, ft...)

	// Build the leader
	recLen := []byte(fmt.Sprintf("%05d", 24+len(rawDir)+len(cfs)+len(dfs)+1))
	recBaseDataAddress := []byte(fmt.Sprintf("%05d", 24+len(rawDir)))

	ldr := []byte(rec.Leader.Text)
	for i := 0; i <= 4; i++ {
		ldr[i] = recLen[i]
		ldr[i+12] = recBaseDataAddress[i]
	}

	// Final assembly
	marc = append(marc, ldr...)
	marc = append(marc, rawDir...)
	marc = append(marc, cfs...)
	marc = append(marc, dfs...)
	marc = append(marc, rt...)

	return marc, nil
}

// termAsByte converts a terminator/delimiter value to a byte
func termAsByte(i int) (b []byte) {

	// We apparently need a 2-byte array for setting the uint16...
	x := make([]byte, 2)
	binary.LittleEndian.PutUint16(x, uint16(i))
	// ... but we don't want the trailing null byte in actual use
	b = x[:1]

	return b
}

// toInt converts a byte array of digits to its corresponding integer
// value
func toInt(b []byte) (ret int, err error) {
	ret, err = strconv.Atoi(string(b))
	if err != nil {

		var digits = map[string]int{
			"0": 0,
			"1": 1,
			"2": 2,
			"3": 3,
			"4": 4,
			"5": 5,
			"6": 6,
			"7": 7,
			"8": 8,
			"9": 9,
		}

		ret = 0
		for i := range b {
			x, ok := digits[string(b[i])]
			if !ok {
				return 0, errors.New("toInt(): Not an integer")
			}
			ret = (10 * ret) + x
		}
	}
	return ret, nil
}
