// Copyright 2017-2018 Gregory Siems. All rights reserved.
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
	// FmtUnknown indicates that the record format is not known or not specified
	FmtUnknown = iota
	// Bibliography indicates that the record is a Bibliography (or bib-holding) record
	Bibliography
	// Holdings indicates that the record is a Holdings record
	Holdings
	// Authority indicates that the record is an Authority record
	Authority
	// Classification indicates that the record is a Classification record
	Classification
	// Community indicates that the record is a Community Information record
	Community
)

// RecordFormat indicates the high level nature of the record and is
// used to differentiate between Bibliography, Holdings, Authority,
// Classification, and Community record formats.
func (rec Record) RecordFormat() int {

	code, _ := rec.RecordType()

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
	return FmtUnknown
}

//  05 - Record status
// Valid values for Bibliography records
var recordStatus = map[string]string{
	"a": "Increase in encoding level",                     // Bib, Auth,       Class,
	"c": "Corrected or revised",                           // Bib, Auth, Hold, Class, CI
	"d": "Deleted",                                        // Bib, Auth, Hold, Class, CI
	"n": "New",                                            // Bib, Auth, Hold, Class, CI
	"p": "Increase in encoding level from prepublication", // Bib
}

// Valid values for Authority and Classification records
var recordStatusA = map[string]string{
	"a": "Increase in encoding level", // Bib, Auth,       Class,
	"c": "Corrected or revised",       // Bib, Auth, Hold, Class, CI
	"d": "Deleted",                    // Bib, Auth, Hold, Class, CI
	"n": "New",                        // Bib, Auth, Hold, Class, CI
}

// Valid values for Holdings and Community records
var recordStatusH = map[string]string{
	"c": "Corrected or revised", // Bib, Auth, Hold, Class, CI
	"d": "Deleted",              // Bib, Auth, Hold, Class, CI
	"n": "New",                  // Bib, Auth, Hold, Class, CI
}

// RecordStatus returns the one character code and label indicating
// the "05 Record status"
func (rec Record) RecordStatus() (code, label string) {

	switch rec.RecordFormat() {
	case Bibliography:
		code, label = shortCodeLookup(recordStatus, rec.Leader.Text, 5)
	case Holdings, Community:
		code, label = shortCodeLookup(recordStatusH, rec.Leader.Text, 5)
	case Authority, Classification:
		code, label = shortCodeLookup(recordStatusA, rec.Leader.Text, 5)
	}
	return code, label
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
	return shortCodeLookup(recordType, rec.Leader.Text, 6)
}

//  09 - Character coding scheme
var characterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}

// CharacterCodingScheme returns the code and label indicating the
// "09 - Character coding scheme" of the record (MARC-8 or UCS/Unicode).
func (rec Record) CharacterCodingScheme() (code, label string) {
	return shortCodeLookup(characterCodingScheme, rec.Leader.Text, 9)
}

////////////////////////////////////////////////////////////////////////
// Functions specific to Bibliography formats

//  07 - "Bibliographic level" for Bibliography records.
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
// "07 - Bibliographic level" of the Bibliography record.
func (rec Record) BibliographicLevel() (code, label string) {
	if rec.RecordFormat() == Bibliography {
		code, label = shortCodeLookup(bibliographicLevel, rec.Leader.Text, 7)
	}
	return code, label
}

//  08 - "Type of control" for Bibliography records.
var controlType = map[string]string{
	" ": "No specified type",
	"a": "Archival",
}

// ControlType returns the code and label indicating the
// "08 - Type of control" of the Bibliography record.
func (rec Record) ControlType() (code, label string) {
	if rec.RecordFormat() == Bibliography {
		code, label = shortCodeLookup(controlType, rec.Leader.Text, 8)
	}
	return code, label
}

//  17 - "Encoding level" for Bibliography records.
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

// EncodingLevel returns the code and label indicating the
// "17 - Encoding level" of the Bibliography record.
func (rec Record) EncodingLevel() (code, label string) {
	if rec.RecordFormat() == Bibliography {
		code, label = shortCodeLookup(encodingLevel, rec.Leader.Text, 17)
	}
	return code, label
}

//  18 - "Descriptive cataloging form" for Bibliography records.
var descriptiveCatalogingForm = map[string]string{
	" ": "Non-ISBD",
	"a": "AACR 2",
	"c": "ISBD punctuation omitted",
	"i": "ISBD punctuation included",
	"n": "Non-ISBD punctuation omitted",
	"u": "Unknown",
}

// CatalogingForm returns the code and label indicating the
// "18 - Descriptive cataloging form" of the Bibliography record.
func (rec Record) CatalogingForm() (code, label string) {
	if rec.RecordFormat() == Bibliography {
		code, label = shortCodeLookup(descriptiveCatalogingForm, rec.Leader.Text, 18)
	}
	return code, label
}

//  19 - "Multipart resource record level" for Bibliography records.
var multipartResourceRecordLevel = map[string]string{
	" ": "Not specified or not applicable",
	"a": "Set",
	"b": "Part with independent title",
	"c": "Part with dependent title",
}

// MultipartResourceRecordLevel returns the code and label indicating the
// "19 - Multipart resource record level" of the Bibliography record.
func (rec Record) MultipartResourceRecordLevel() (code, label string) {
	if rec.RecordFormat() == Bibliography {
		code, label = shortCodeLookup(multipartResourceRecordLevel, rec.Leader.Text, 19)
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////

// GetText returns the text for the leader
func (ldr Leader) GetText() string {
	return ldr.Text
}
