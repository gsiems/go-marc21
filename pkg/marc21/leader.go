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

var marcFormatName = map[int]string{
	FmtUnknown:     "Unknown",
	Bibliography:   "Bibliography",
	Holdings:       "Holdings",
	Authority:      "Authority",
	Classification: "Classification",
	Community:      "Community Information",
}

// RecordFormat indicates the high level nature of the record and is
// used to differentiate between Bibliography, Holdings, Authority,
// Classification, and Community record formats.
func (rec Record) RecordFormat() int {
	code := pluckByte(rec.Leader.Text, 6)
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

// RecordFormatName indicates the name of the format of the record and is
// used to differentiate between Bibliography, Holdings, Authority,
// Classification, and Community record formats.
func (rec Record) RecordFormatName() string {
	f := rec.RecordFormat()
	n := marcFormatName[f]
	return n
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
	rectype := pluckByte(rec.Leader.Text, 6)

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
		biblevel := pluckByte(rec.Leader.Text, 7)
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

////////////////////////////////////////////////////////////////////////

// GetText returns the text for the leader
func (ldr Leader) GetText() string {
	return ldr.Text
}
