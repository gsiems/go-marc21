// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package marc21 reads/parses MARC21 and MARCXML data.
//
// Enables converting between MARC21 and MARCXML data.
package marc21

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
)

/*
PRIMARY GOALS:

    Read MARC21 records
    Read MARCXML records/collections

    Write MARC21 records
    Write MARCXML records/collections
    Write "Pretty-print" text (compatible with perl MARC::Record->as_formatted() output)

    Convert between MARC21 and MarcXML

    MARC-8 vs. UTF-8 encoding

Someday, maybe, secondary goals:

    Functions for "human-friendly" extraction of titles, authors, etc.
    Perform [some] validation
    Read/write MARCMaker?? https://www.loc.gov/marc/makrbrkr.html
    Read "Pretty-print" text
    Edit records

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

type Collection struct {
	Name    xml.Name `xml:"collection"`
	Records []Record `xml:"record"`
}

type Leader struct {
	RecordLength           int
	RecordStatus           byte
	RecordType             byte
	BibliographicLevel     byte
	ControlType            byte
	CharacterCodingScheme  byte
	IndicatorCount         byte
	SubfieldCodeCount      byte
	BaseDataAddress        int
	EncodingLevel          byte
	CatalogingForm         byte
	MultipartLevel         byte
	LenOfLengthOfField     byte
	LenOfStartCharPosition byte
	LenOfImplementDefined  byte
	Undefined              byte
}

type LeaderRaw struct {
	Text string `xml:",chardata"`
}

type Record struct {
	Leader        *Leader
	LeaderRaw     LeaderRaw      `xml:"leader"`
	Controlfields []Controlfield `xml:"controlfield"`
	Datafields    []Datafield    `xml:"datafield"`
}

type Controlfield struct {
	Tag  string `xml:"tag,attr"`
	Text string `xml:",chardata"`
}

type Datafield struct {
	Tag       string     `xml:"tag,attr"`
	Ind1      string     `xml:"ind1,attr"`
	Ind2      string     `xml:"ind2,attr"`
	Subfields []Subfield `xml:"subfield"`
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

func ParseRecord(rawRec []byte) (rec *Record, err error) {

	rec = new(Record)

	rec.LeaderRaw.Text = string(rawRec[:24])
	rec.Leader, err = parseLeader(rawRec)
	if err != nil {
		return nil, err
	}

	dir, err := parseDirectory(rawRec)
	if err != nil {
		return nil, err
	}

	rec.Controlfields, err = parseControlfields(rawRec, rec.Leader.BaseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	rec.Datafields, err = parseDatafields(rawRec, rec.Leader.BaseDataAddress, dir)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (rec Record) String() string {

	ret := fmt.Sprintf("LDR %s\n", rec.LeaderRaw.Text)
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
