// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"fmt"
	"strconv"
)

/*
https://www.loc.gov/marc/specifications/specrecstruc.html

    The leader is the first field in the record and has a fixed length
    of 24 octets (character positions 0-23). Only ASCII graphic
    characters are allowed in the Leader.
*/

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Leader - Data elements that primarily provide information for the
    processing of the record. The data elements contain numbers or
    coded values and are identified by relative character position. The
    Leader is fixed in length at 24 character positions and is the
    first field of a MARC record.
*/

// http://www.loc.gov/marc/bibliographic/bdleader.html
//
//  Character Positions
//  00-04 - Record length

//  05 - Record status
var recordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"p": "Increase in encoding level from prepublication",
}

func (rec Record) RecordStatus() (code, label string) {
	if len(rec.Leader.Text) > 5 {
		code = string(rec.Leader.Text[5])
		label, _ = recordStatus[code]
	}
	return code, label
}

//  06 - Type of record
var recordType = map[string]string{
	"a": "Language material",
	"c": "Notated music",
	"d": "Manuscript notated music",
	"e": "Cartographic material",
	"f": "Manuscript cartographic material",
	"g": "Projected medium",
	"i": "Nonmusical sound recording",
	"j": "Musical sound recording",
	"k": "Two-dimensional nonprojectable graphic",
	"m": "Computer file",
	"o": "Kit",
	"p": "Mixed materials",
	"r": "Three-dimensional artifact or naturally occurring object",
	"t": "Manuscript language material",
}

// RecordType returns the one character code and label indicating
// the type of content and material documented by the record.
func (rec Record) RecordType() (code, label string) {
	if len(rec.Leader.Text) > 6 {
		code = string(rec.Leader.Text[6])
		label, _ = recordType[code]
	}
	return code, label
}

//  07 - Bibliographic level
var bibliographicLevel = map[string]string{
	"a": "Monographic component part",
	"b": "Serial component part",
	"c": "Collection",
	"d": "Subunit",
	"i": "Integrating resource",
	"m": "Monograph/Item",
	"s": "Serial",
}

// BibliographicLevel returns the code and label indicating the
// bibliographic level of the record.
func (rec Record) BibliographicLevel() (code, label string) {
	if len(rec.Leader.Text) > 7 {
		code = string(rec.Leader.Text[7])
		label, _ = bibliographicLevel[code]
	}
	return code, label
}

//  08 - Type of control
var controlType = map[string]string{
	" ": "No specified type",
	"a": "Archival",
}

func (rec Record) ControlType() (code, label string) {
	if len(rec.Leader.Text) > 8 {
		code = string(rec.Leader.Text[8])
		label, _ = controlType[code]
	}
	return code, label
}

//  09 - Character coding scheme
var characterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}

func (rec Record) CharacterCodingScheme() (code, label string) {
	if len(rec.Leader.Text) > 9 {
		code = string(rec.Leader.Text[9])
		label, _ = characterCodingScheme[code]
	}
	return code, label
}

//  10 - Indicator count
//      2 - Number of character positions used for indicators
//
//  11 - Subfield code count
//      2 - Number of character positions used for a subfield code
//
//  12-16 - Base address of data
//      [number] - Length of Leader and Directory
//
//  17 - Encoding level
var encodingLevel = map[string]string{
	" ": "Full level",
	"1": "Full level, material not examined",
	"2": "Less-than-full level, material not examined",
	"3": "Abbreviated level",
	"4": "Core level",
	"5": "Partial (preliminary) level",
	"7": "Minimal level",
	"8": "Prepublication level",
	"u": "Unknown",
	"z": "Not applicable",
}

func (rec Record) EncodingLevel() (code, label string) {
	if len(rec.Leader.Text) > 17 {
		code = string(rec.Leader.Text[17])
		label, _ = encodingLevel[code]
	}
	return code, label
}

//  18 - Descriptive cataloging form
var descriptiveCatalogingForm = map[string]string{
	" ": "Non-ISBD",
	"a": "AACR 2",
	"c": "ISBD punctuation omitted",
	"i": "ISBD punctuation included",
	"n": "Non-ISBD punctuation omitted",
	"u": "Unknown",
}

func (rec Record) CatalogingForm() (code, label string) {
	if len(rec.Leader.Text) > 18 {
		code = string(rec.Leader.Text[18])
		label, _ = descriptiveCatalogingForm[code]
	}
	return code, label
}

//  19 - Multipart resource record level
var multipartResourceRecordLevel = map[string]string{
	" ": "Not specified or not applicable",
	"a": "Set",
	"b": "Part with independent title",
	"c": "Part with dependent title",
}

func (rec Record) MultipartResourceRecordLevel() (code, label string) {
	if len(rec.Leader.Text) > 19 {
		code = string(rec.Leader.Text[19])
		label, _ = multipartResourceRecordLevel[code]
	}
	return code, label
}

//  20 - Length of the length-of-field portion
//      4 - Number of characters in the length-of-field portion of a
//          Directory entry
//
//  21 - Length of the starting-character-position portion
//      5 - Number of characters in the starting-character-position
//          portion of a Directory entry
//
//  22 - Length of the implementation-defined portion
//      0 - Number of characters in the implementation-defined portion
//          of a Directory entry
//
//  23 - Undefined
//      0 - Undefined

/*
Example:

    01166cam  2200313   450000100
    01166   | RecordLength           int   |
    c       | RecordStatus           byte  | c - Corrected or revised
    a       | RecordType             byte  | a - Language material
    m       | BibliographicLevel     byte  | a - Monographic component part
            | ControlType            byte  | # - No specified type
            | CharacterCodingScheme  byte  | # - MARC-8
    2       | IndicatorCount         byte  | 2 - Number of character positions used for indicators
    2       | SubfieldCodeCount      byte  | 2 - Number of character positions used for a subfield code
    00313   | BaseDataAddress        int   |
            | EncodingLevel          byte  | # - Full level
            | CatalogingForm         byte  | # - Non-ISBD
            | MultipartLevel         byte  | # - Not specified or not applicable
    4       | LenOfLengthOfField     byte  | 4 - Number of characters in the length-of-field portion of a Directory entry
    5       | LenOfStartCharPosition byte  | 5 - Number of characters in the starting-character-position portion of a Directory entry
    0       | LenOfImplementDefined  byte  | 0 - Number of characters in the implementation-defined portion of a Directory entry
    0       | Undefined              byte  | 0 - Undefined
*/

// TODO: Do we want to/have use for leader validation?

// parseLeader extracts the leader information from the raw MARC record bytes
func parseLeader(b []byte) (l *ParsedLeader, err error) {
	l = new(ParsedLeader)

	l.RecordLength, err = strconv.Atoi(string(b[0:5]))
	if err != nil {
		return nil, err
	}
	l.RecordStatus = b[5]
	l.RecordType = b[6]
	l.BibliographicLevel = b[7]
	l.ControlType = b[8]
	l.CharacterCodingScheme = b[9]
	l.IndicatorCount = b[10]
	l.SubfieldCodeCount = b[11]
	l.BaseDataAddress, err = strconv.Atoi(string(b[12:17]))
	if err != nil {
		return nil, err
	}
	l.EncodingLevel = b[17]
	l.CatalogingForm = b[18]
	l.MultipartLevel = b[19]
	l.LenOfLengthOfField = b[20]
	l.LenOfStartCharPosition = b[21]
	l.LenOfImplementDefined = b[22]
	l.Undefined = b[23]

	return l, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (ldr Leader) String() string {

	b := []byte(ldr.Text)
	ret := fmt.Sprintf("LDR %s\n", ldr.Text)

	code := string(b[0:5])
	ret += fmt.Sprintf("    %s: ( RecordLength )\n", code)

	code = string(b[5])
	label, _ := recordStatus[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( RecordStatus = %q )\n", code, label)

	code = string(b[6])
	label, _ = recordType[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( RecordType = %q )\n", code, label)

	code = string(b[7])
	label, _ = bibliographicLevel[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( BibliographicLevel = %q )\n", code, label)

	code = string(b[8])
	label, _ = controlType[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( ControlType = %q )\n", code, label)

	code = string(b[9])
	label, _ = characterCodingScheme[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( CharacterCodingScheme = %q )\n", code, label)

	code = string(b[10])
	ret += fmt.Sprintf("        %s: ( IndicatorCount )\n", code)

	code = string(b[11])
	ret += fmt.Sprintf("        %s: ( SubfieldCodeCount )\n", code)

	code = string(b[12:17])
	ret += fmt.Sprintf("    %s: ( BaseAddressOfData )\n", code)

	code = string(b[17])
	label, _ = encodingLevel[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( EncodingLevel = %q )\n", code, label)

	code = string(b[18])
	label, _ = descriptiveCatalogingForm[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( DescriptiveCatalogingForm = %q )\n", code, label)

	code = string(b[19])
	label, _ = multipartResourceRecordLevel[code]
	if code == " " {
		code = "#"
	}
	ret += fmt.Sprintf("        %s: ( MultipartResourceRecordLevel = %q )\n", code, label)

	ret += fmt.Sprintln("        4: ( LengthOfLengthOfField )")
	ret += fmt.Sprintln("        5: ( LengthOfStartingCharacterPosition )")
	ret += fmt.Sprintln("        0: ( LengthOfImplementationDefined )")
	ret += fmt.Sprintln("        0: ( Undefined )")

	return ret
}
