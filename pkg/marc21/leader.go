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
Note that it should be possible to generate the majority of this file
using the information in the following LOC web pages:

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

// validFields is the base for the valid fields mappings for the
// different MARC record formats
type validFields map[string]int

// ValidLeaderFields returns a list of the different lookup fields
// that are valid for the MARC record (based on record format)
func (rec Record) ValidLeaderFields() (f []string) {

	var a validFields

	switch rec.RecordFormat() {
	case Bibliography:
		a = validBibliographyFields
	case Holdings:
		a = validHoldingsFields
	case Authority:
		a = validAuthorityFields
	case Classification:
		a = validClassificationFields
	case Community:
		a = validCommunityFields
	}

	if a != nil {
		for k := range a {
			f = append(f, k)
		}
	}
	return f
}

// LookupLeaderField looks up the code and label for the specified
// lookup field
func (rec Record) LookupLeaderField(s string) (code, label string) {
	switch rec.RecordFormat() {
	case Bibliography:
		return rec.lookupBibliographyField(s)
	case Holdings:
		return rec.lookupHoldingsField(s)
	case Authority:
		return rec.lookupAuthorityField(s)
	case Classification:
		return rec.lookupClassificationField(s)
	case Community:
		return rec.lookupCommunityField(s)
	}
	return "", ""
}

// RecordStatus returns the one character code and label indicating
// the "05 Record status"
func (rec Record) RecordStatus() (code, label string) {
	switch rec.RecordFormat() {
	case Bibliography:
		code, label = shortCodeLookup(bibliographyRecordStatus, rec.Leader.Text, 5)
	case Holdings:
		code, label = shortCodeLookup(holdingsRecordStatus, rec.Leader.Text, 5)
	case Authority:
		code, label = shortCodeLookup(authorityRecordStatus, rec.Leader.Text, 5)
	case Classification:
		code, label = shortCodeLookup(classificationRecordStatus, rec.Leader.Text, 5)
	case Community:
		code, label = shortCodeLookup(communityRecordStatus, rec.Leader.Text, 5)
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
// Lookup data for Bibliography records

var bibliographyMaterialType = map[string]string{
	"BK": "Books",
	"CF": "Computer Files",
	"MP": "Maps",
	"MU": "Music",
	"CR": "Continuing Resources",
	"VM": "Visual Materials",
	"MX": "Mixed Materials",
}

// BibliographyMaterialType returns the code and description of the type
// of material documented by Bibliography the record.
// {"Books",  "Computer Files", "Maps", "Music", "Continuing Resources",
// "Visual Materials" or "Mixed Materials"}
func (rec Record) BibliographyMaterialType() (code, label string) {
	rectype := shortCode(rec.Leader.Text, 6)

	// the simple record type to material type mappings
	var rtmap = map[string]string{
		"c": "MU",
		"d": "MU",
		"i": "MU",
		"j": "MU",
		"e": "MP",
		"f": "MP",
		"g": "VM",
		"k": "VM",
		"o": "VM",
		"r": "VM",
		"m": "CF",
		"p": "MX",
	}

	code, ok := rtmap[rectype]
	if !ok {
		// no simple match
		biblevel := shortCode(rec.Leader.Text, 7)
		switch biblevel {
		case "a", "c", "d", "m":
			if rectype == "a" || rectype == "t" {
				code = "BK"
			}
		case "b", "i", "s":
			if rectype == "a" {
				code = "CR"
			}
		}
	}

	label = bibliographyMaterialType[code]
	return code, label
}

var validBibliographyFields = validFields{
	"RecordStatus":                 5,
	"TypeOfRecord":                 6,
	"BibliographicLevel":           7,
	"TypeOfControl":                8,
	"CharacterCodingScheme":        9,
	"EncodingLevel":                17,
	"DescriptiveCatalogingForm":    18,
	"MultipartResourceRecordLevel": 19,
}
var bibliographyFieldData = map[string]map[string]string{
	"RecordStatus":                 bibliographyRecordStatus,
	"TypeOfRecord":                 bibliographyTypeOfRecord,
	"BibliographicLevel":           bibliographyBibliographicLevel,
	"TypeOfControl":                bibliographyTypeOfControl,
	"CharacterCodingScheme":        characterCodingScheme,
	"EncodingLevel":                bibliographyEncodingLevel,
	"DescriptiveCatalogingForm":    bibliographyDescriptiveCatalogingForm,
	"MultipartResourceRecordLevel": bibliographyMultipartResourceRecordLevel,
}

var bibliographyRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"p": "Increase in encoding level from prepublication",
}
var bibliographyTypeOfRecord = map[string]string{
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
var bibliographyBibliographicLevel = map[string]string{
	"a": "Monographic component part",
	"b": "Serial component part",
	"c": "Collection",
	"d": "Subunit",
	"i": "Integrating resource",
	"m": "Monograph/Item",
	"s": "Serial",
}
var bibliographyTypeOfControl = map[string]string{
	" ": "No specified type",
	"a": "Archival",
}
var bibliographyEncodingLevel = map[string]string{
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
var bibliographyDescriptiveCatalogingForm = map[string]string{
	" ": "Non-ISBD",
	"a": "AACR 2",
	"c": "ISBD punctuation omitted",
	"i": "ISBD punctuation included",
	"n": "Non-ISBD punctuation omitted",
	"u": "Unknown",
}
var bibliographyMultipartResourceRecordLevel = map[string]string{
	" ": "Not specified or not applicable",
	"a": "Set",
	"b": "Part with independent title",
	"c": "Part with dependent title",
}

func (rec Record) lookupBibliographyField(s string) (code, label string) {
	if rec.RecordFormat() == Bibliography {
		i, ok := validBibliographyFields[s]
		if ok {
			code, label = shortCodeLookup(bibliographyFieldData[s], rec.Leader.Text, i)
		}
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////
// Lookup data for Holdings records

var validHoldingsFields = validFields{
	"RecordStatus":            5,
	"TypeOfRecord":            6,
	"CharacterCodingScheme":   9,
	"EncodingLevel":           17,
	"ItemInformationInRecord": 18,
}
var holdingsFieldData = map[string]map[string]string{
	"RecordStatus":            holdingsRecordStatus,
	"TypeOfRecord":            holdingsTypeOfRecord,
	"CharacterCodingScheme":   characterCodingScheme,
	"EncodingLevel":           holdingsEncodingLevel,
	"ItemInformationInRecord": holdingItemInformationInRecord,
}

var holdingsRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var holdingsTypeOfRecord = map[string]string{
	"u": "Unknown",
	"v": "Multipart item holdings",
	"x": "Single-part item holdings",
	"y": "Serial item holdings",
}
var holdingsEncodingLevel = map[string]string{
	"1": "Holdings level 1",
	"2": "Holdings level 2",
	"3": "Holdings level 3",
	"4": "Holdings level 4",
	"5": "Holdings level 4 with piece designation",
	"m": "Mixed level",
	"u": "Unknown",
	"z": "Other level",
}
var holdingItemInformationInRecord = map[string]string{
	"i": "Item information",
	"n": "No item information",
}

func (rec Record) lookupHoldingsField(s string) (code, label string) {
	if rec.RecordFormat() == Holdings {
		i, ok := validHoldingsFields[s]
		if ok {
			code, label = shortCodeLookup(holdingsFieldData[s], rec.Leader.Text, i)
		}
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////
// Lookup data for Authority records

var validAuthorityFields = validFields{
	"RecordStatus":          5,
	"TypeOfRecord":          6,
	"CharacterCodingScheme": 9,
	"EncodingLevel":         17,
	"PunctuationPolicy":     18,
}
var authorityFieldData = map[string]map[string]string{
	"RecordStatus":          authorityRecordStatus,
	"TypeOfRecord":          authorityTypeOfRecord,
	"CharacterCodingScheme": characterCodingScheme,
	"EncodingLevel":         authorityEncodingLevel,
	"PunctuationPolicy":     authorityPunctuationPolicy,
}

var authorityRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"o": "Obsolete",
	"s": "Deleted; heading split into two or more headings",
	"x": "Deleted; heading replaced by another heading",
}
var authorityTypeOfRecord = map[string]string{
	"z": "Authority data",
}
var authorityEncodingLevel = map[string]string{
	"n": "Complete authority record",
	"o": "Incomplete authority record",
}
var authorityPunctuationPolicy = map[string]string{
	" ": "No information provided",
	"c": "Punctuation omitted",
	"i": "Punctuation included",
	"u": "Unknown",
}

func (rec Record) lookupAuthorityField(s string) (code, label string) {
	if rec.RecordFormat() == Authority {
		i, ok := validAuthorityFields[s]
		if ok {
			code, label = shortCodeLookup(authorityFieldData[s], rec.Leader.Text, i)
		}
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////
// Lookup data for Classification records

var validClassificationFields = validFields{
	"RecordStatus":          5,
	"TypeOfRecord":          6,
	"CharacterCodingScheme": 9,
	"EncodingLevel":         17,
}
var classificationFieldData = map[string]map[string]string{
	"RecordStatus":          classificationRecordStatus,
	"TypeOfRecord":          classificationTypeOfRecord,
	"CharacterCodingScheme": characterCodingScheme,
	"EncodingLevel":         classificationEncodingLevel,
}

var classificationRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var classificationTypeOfRecord = map[string]string{
	"w": "Classification data",
}
var classificationEncodingLevel = map[string]string{
	"n": "Complete classification record",
	"o": "Incomplete classification record",
}

func (rec Record) lookupClassificationField(s string) (code, label string) {
	if rec.RecordFormat() == Classification {
		i, ok := validClassificationFields[s]
		if ok {
			code, label = shortCodeLookup(classificationFieldData[s], rec.Leader.Text, i)
		}
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////
// Lookup data for Community records

var validCommunityFields = validFields{
	"RecordStatus":          5,
	"TypeOfRecord":          6,
	"KindOfData":            7,
	"CharacterCodingScheme": 9,
}
var communityFieldData = map[string]map[string]string{
	"RecordStatus":          communityRecordStatus,
	"TypeOfRecord":          communityTypeOfRecord,
	"KindOfData":            communityKindOfData,
	"CharacterCodingScheme": characterCodingScheme,
}

var communityRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var communityTypeOfRecord = map[string]string{
	"q": "Community information",
}
var communityKindOfData = map[string]string{
	"n": "Individual",
	"o": "Organization",
	"p": "Program or service",
	"q": "Event",
	"z": "Other",
}

func (rec Record) lookupCommunityField(s string) (code, label string) {
	if rec.RecordFormat() == Community {
		i, ok := validCommunityFields[s]
		if ok {
			code, label = shortCodeLookup(communityFieldData[s], rec.Leader.Text, i)
		}
	}
	return code, label
}

////////////////////////////////////////////////////////////////////////

// GetText returns the text for the leader
func (ldr Leader) GetText() string {
	return ldr.Text
}
