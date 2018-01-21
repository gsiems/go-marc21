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

// Collection is fo containing zero or more MARC records
type Collection struct {
	Name    xml.Name  `xml:"collection"`
	Records []*Record `xml:"record"`
}

// Leader is for containing the text string of the MARC record Leader
type Leader struct {
	text string `xml:",chardata"`
}

// Record is for containing a MARC record
type Record struct {
	leader        Leader          `xml:"leader"`
	controlfields []*Controlfield `xml:"controlfield"`
	datafields    []*Datafield    `xml:"datafield"`
}

// directoryEntry contains a single directory entry
type directoryEntry struct {
	tag         string
	startingPos int
	fieldLength int
}

// Controlfield contains a controlfield entry
type Controlfield struct {
	tag  string `xml:"tag,attr"`
	text string `xml:",chardata"`
}

// Datafield contains a datafield entry
type Datafield struct {
	tag       string      `xml:"tag,attr"`
	ind1      string      `xml:"ind1,attr"`
	ind2      string      `xml:"ind2,attr"`
	subfields []*Subfield `xml:"subfield"`
}

// Subfield contains a subfield entry
type Subfield struct {
	code string `xml:"code,attr"`
	text string `xml:",chardata"`
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

	recLen, err := strconv.Atoi(string(rawLen[0:5]))
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

	rec.leader.text = string(rawRec[:24])

	dir, err := parseDirectory(rawRec)
	if err != nil {
		return nil, err
	}

	baseDataAddress, err := strconv.Atoi(string(rawRec[12:17]))
	if err != nil {
		return nil, err
	}

	rec.controlfields, err = extractControlfields(rawRec, baseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	rec.datafields, err = extractDatafields(rawRec, baseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (rec Record) String() string {

	ret := fmt.Sprintf("LDR %s\n", rec.leader.Text())
	for _, cf := range rec.controlfields {
		ret += fmt.Sprintf("%s     %s\n", cf.Tag(), cf.Text())
	}
	for _, df := range rec.datafields {
		pre := fmt.Sprintf("%s %s%s _", df.Tag(), df.Ind1(), df.Ind2())
		for _, sf := range df.subfields {
			ret += fmt.Sprintf("%s%s%s\n", pre, sf.Code(), sf.Text())
			pre = "       _"
		}
	}
	return ret
}

// RecordAsMARC converts a Record into a MARC record byte array
func (rec Record) RecordAsMARC() (marc []byte, err error) {

	if rec.leader.text == "" {
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
	for _, cf := range rec.controlfields {

		if cf.Text() == "" {
			continue
		}

		b := []byte(cf.Text())
		b = append(b, ft...)
		cfs = append(cfs, b...)

		dir = append(dir, directoryEntry{tag: cf.Tag(), startingPos: startPos, fieldLength: len(b)})

		startPos += len(b)
	}

	// Pack the data fields/sub-fields
	for _, df := range rec.datafields {

		ind1 := df.Ind1()
		ind2 := df.Ind2()

		b := []byte(ind1)
		b = append(b, []byte(ind2)...)

		for _, sf := range df.subfields {
			b = append(b, dl...)
			b = append(b, []byte(sf.Code())...)
			b = append(b, []byte(sf.Text())...)
		}
		b = append(b, ft...)
		dfs = append(dfs, b...)

		dir = append(dir, directoryEntry{tag: df.Tag(), startingPos: startPos, fieldLength: len(b)})

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

	ldr := []byte(rec.leader.text)
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
