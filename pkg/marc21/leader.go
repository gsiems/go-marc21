// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

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

Also:
    http://www.loc.gov/marc/holdings/hdleader.html
    http://www.loc.gov/marc/authority/adleader.html
    http://www.loc.gov/marc/classification/cdleader.html
    http://www.loc.gov/marc/community/cileader.html

While the general leader layout is the same for the different MARC formats
there are differences.

MARC 21 Bibliography
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07 - Bibliographic level
    08 - Type of control
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code count
    12-16 - Base address of data
    17 - Encoding level
    18 - Descriptive cataloging form
    19 - Multipart resource record level
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Holdings
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18 - Item information in record
    19 - Undefined character position
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Authority
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18 - Punctuation policy
    19 - Undefined
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Classification
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18-19 - Undefined character positions
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Community Information
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07 - Kind of data
    08 - Undefined character position
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17-19 - Undefined character positions
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

*/

/*
http://www.loc.gov/marc/bibliographic/ecbdlist.html
http://www.loc.gov/marc/holdings/echdlist.html
http://www.loc.gov/marc/authority/ecadlist.html
http://www.loc.gov/marc/classification/eccdlist.html
http://www.loc.gov/marc/community/eccilist.html
*/

const (
	FmtUnknown = iota
	Bibliography
	Holdings
	Authority
	Classification
	Community
)

// RecordFormat indicates the high level nature of the record and is
// used to differentiate between Bibliography, Holdings, Authority,
// Classification, and Community record formats.
func (rec Record) RecordFormat() int {

	if len(rec.Leader.Text) > 6 {
		code := string(rec.Leader.Text[6])
		switch code {
		case "q":
			return Community
		case "z":
			return Authority
		case "w":
			return Classification
		case "u", "v", "x", "y":
			return Holdings
		case "a", "c", "d", "e", "f", "g", "i", "j", "k", "m", "o", "p", "r", "t":
			return Bibliography
		}
	}
	return FmtUnknown
}

//  06 - Type of record
var recordType = map[string]string{
	// Bibliography
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
	// Holding
	"u": "Unknown",
	"v": "Multipart item holdings",
	"x": "Single-part item holdings",
	"y": "Serial item holdings",
	// Classification
	"w": "Classification data",
	// Authority
	"z": "Authority data",
	// Community
	"q": "Community information",
}

// RecordType returns the one character code and label indicating
// the "06 - Type of record" for the record. Use RecordFormat
// to determine the record format (bibliographic, holdings, etc.)
func (rec Record) RecordType() (code, label string) {
	if len(rec.Leader.Text) > 6 {
		code = string(rec.Leader.Text[6])
		label = recordType[code]
	}
	return code, label
}

//  09 - Character coding scheme
var characterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}

// CharacterCodingScheme returns the code and label indicating the
// "09 - Character coding scheme" of the record (MARC-8 or UCS/Unicode).
func (rec Record) CharacterCodingScheme() (code, label string) {
	if len(rec.Leader.Text) > 9 {
		code = string(rec.Leader.Text[9])
		label = characterCodingScheme[code]
	}
	return code, label
}

// GetText returns the text for the leader
func (ldr Leader) GetText() string {
	return ldr.Text
}
