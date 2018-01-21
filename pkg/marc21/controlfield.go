// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"fmt"
	"log"
	"strings"
)

/*
https://www.loc.gov/marc/specifications/specrecstruc.html

    Control fields in MARC 21 formats are assigned tags beginning with
    two zeroes. They are comprised of data and a field terminator; they
    do not contain indicators or subfield codes. The control number
    field is assigned tag 001 and contains the control number of the
    record. Each record contains only one control number field (with
    tag 001), which is to be located at the base address of data.
*/

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Variable control fields - The 00X fields. These fields are
    identified by a field tag in the Directory but they contain neither
    indicator positions nor subfield codes. The variable control fields
    are structurally different from the variable data fields. They may
    contain either a single data element or a series of fixed-length
    data elements identified by relative character position.
*/

// http://www.loc.gov/marc/bibliographic/bd00x.html

// CfValue contains a code and it's corresponding descriptive label
// for a controlfield entry.
type CfValue struct {
	Code  string
	Label string
}

var materialType = map[string]string{
	"BK": "Books",
	"CF": "Computer Files",
	"MP": "Maps",
	"MU": "Music",
	"CR": "Continuing Resources",
	"VM": "Visual Materials",
	"MX": "Mixed Materials",
}

// cfShortCode performs lookups on single-character reference tables (maps)
// and returns the code and, if found, the descriptive label for the code.
func cfShortCode(codeList map[string]string, b []byte, i int) (v CfValue) {

	var code, label string

	if len(b) > i {
		code = string(b[i])
		if code != "" {
			label = codeList[code]
		}
	}

	return CfValue{Code: code, Label: label}
}

// cfWideCode performs lookups on multi-character reference tables (maps)
// and returns the code and, if found, the descriptive label for the code.
func cfWideCode(codeList map[string]string, b []byte, i, w int) (v CfValue) {

	var code, label string

	if len(b) > i+w {
		code = string(b[i : i+w])
		if code != "" {
			label = codeList[code]
		}
	}

	return CfValue{Code: code, Label: label}
}

// cfPluckVal extracts one or more bytes from a byte slice and returns
// the string result.
func cfPluckVal(b []byte, i, w int) (s string) {
	if len(b) > i+w {
		s = string(b[i : i+w])
	}
	return s
}

// CfData contains the results from parsing the controlfields for a MARC
// record.
type CfData struct {
	ControlNumber           string
	ControlNumberIdentifier string
	LatestTransDateTime     string
	FileDate                string
	DateTypePubStatus       CfValue
	Date1                   string
	Date2                   string
	PlaceOfPublication      string
	Language                string
	ModifiedRecord          CfValue
	CatalogingSource        CfValue
	MaterialType            CfValue
	MaterialCharactor       CfMatlDesc
	PhysicalDescription     []CfPhysDesc
}

// ParseControlfields extracts the information encoded in the control
// fields.
func (rec Record) ParseControlfields() (c CfData) {

	// 008:
	//  00-05 - Date entered on file
	//  06 - Type of date/Publication status
	//  07-10 - Date 1
	//  11-14 - Date 2
	//  15-17 - Place of publication, production, or execution
	//  18-34 Non-common elements
	//  35-37 - Language
	//  38 - Modified record
	//  39 - Cataloging source

	c.ControlNumber = rec.Controlfield("001")
	c.ControlNumberIdentifier = rec.Controlfield("003")
	c.LatestTransDateTime = rec.Controlfield("005")

	b := []byte(rec.Controlfield("008"))
	c.FileDate = cfPluckVal(b, 0, 5)
	c.DateTypePubStatus = cfShortCode(dateTypePubStatus, b, 6)
	c.Date1 = cfPluckVal(b, 7, 4)
	c.Date2 = cfPluckVal(b, 11, 4)
	c.PlaceOfPublication = cfPluckVal(b, 15, 3)
	c.Language = cfPluckVal(b, 35, 3)
	c.ModifiedRecord = cfShortCode(modifiedRecord, b, 38)
	c.CatalogingSource = cfShortCode(catalogingSource, b, 39)

	mt, label := rec.MaterialType()
	if mt == "" {
		// MARC data can be really dorky...
		log.Printf("Error determining Material Type for ControlNumber %q (defaulting to BK/Book)\n", c.ControlNumber)
		mt = "BK"
		label = materialType[mt]
	}
	c.MaterialType = CfValue{Code: mt, Label: label}

	c.MaterialCharactor = make(CfMatlDesc)

	type fn func(md *CfMatlDesc, b []byte)

	m := map[string]fn{
		"BK": parseBookCf,
		"CF": parseComputerFilesCf,
		"MP": parseMapCf,
		"MU": parseMusicCf,
		"CR": parseContinuingResourcesCf,
		"VM": parseVisualMaterialsCf,
		"MX": parseMixedMaterialsCf,
	}

	if len(b) > 18 {
		// Call the parse function with everything 18 and beyond
		// even though the bytes after 34 aren't material specific. The
		// parse functions will ignore any extra bytes and, even though
		// some 008 fields are "short", we'd still like to get as
		// much data as possible.
		fcn, ok := m[mt]
		if ok {
			fcn(&c.MaterialCharactor, b[18:])
		}
	}

	cf006 := rec.Controlfields("006")
	for _, x := range cf006 {
		b := []byte(x.text)
		if len(b) > 1 {
			// NOTE: The Material Form is b[0] and ignoring it *could*
			// disconnect the material form from the corresponding values.
			// However the material form *should* match the leader record
			// type so that really shouldn't be an issue... should it?
			// Perhaps a check to verify that b[0] matches?
			fcn, ok := m[mt]
			if ok {
				fcn(&c.MaterialCharactor, b[1:])
			}
		}
	}

	c.PhysicalDescription = rec.parse007fields()
	return c
}

// MaterialType returns the code and description of the type of material
// documented by the record. {"Books", "Computer Files", "Maps", "Music",
// "Continuing Resources", "Visual Materials" or "Mixed Materials"}
func (rec Record) MaterialType() (code, label string) {
	rt, _ := rec.RecordType()
	bl, _ := rec.BibliographicLevel()

	switch rt {
	case "c", "d", "i", "j":
		code = "MU"
	case "e", "f":
		code = "MP"
	case "g", "k", "o", "r":
		code = "VM"
	case "m":
		code = "CF"
	case "p":
		code = "MX"
	case "t":
		switch bl {
		case "a", "c", "d", "m":
			code = "BK"
		}
	case "a":
		switch bl {
		case "a", "c", "d", "m":
			code = "BK"
		case "b", "i", "s":
			code = "CR"
		}
	}
	label = materialType[code]
	return code, label
}

// extractControlfields extracts the control fields from the raw MARC record bytes
func extractControlfields(rawRec []byte, baseAddress int, dir []*directoryEntry) (cfs []*Controlfield, err error) {

	// There are records where the 003 and 007 fields are dorky (this
	// may happen to other fields also??) where the first byte is a
	// terminator character and the directory indicates that the field
	// is longer.
	//
	// The directory for one record that has this issue looks like:
	//      001 0 11
	//      003 11 12
	//      005 12 17
	//      007 29 18
	//      008 30 41
	// where we can see that the directory has the 003 and 005 tags
	// overlapping and the 007 and 008 tags overlapping with no actual
	// data for either dorked-up 003/007 tag. Since the remainder of the
	// record appears to be good we don't want to fail, but we do want
	// to bring attention to the data issue.
	var parseErrorTags []string
	var controlNumber string

	for _, d := range dir {
		if strings.HasPrefix(d.tag, "00") {

			start := baseAddress + d.startingPos
			b := rawRec[start : start+d.fieldLength]

			if b[len(b)-1] == fieldTerminator {
				if d.tag == "001" {
					if controlNumber == "" {
						controlNumber = string(b[:len(b)-1])
					} else {
						parseErrorTags = append(parseErrorTags, d.tag)
					}
				}
				cfs = append(cfs, &Controlfield{tag: d.tag, text: string(b[:len(b)-1])})
			} else {
				parseErrorTags = append(parseErrorTags, d.tag)
			}
		}
	}

	if len(parseErrorTags) > 0 {
		badTags := strings.Join(parseErrorTags, ", ")
		log.Printf("Control fields extraction error for ControlNumber %q (fields: %s)\n", controlNumber, badTags)
	}

	return cfs, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (cf Controlfield) String() string {
	return fmt.Sprintf("{%s: '%s'}", cf.tag, cf.text)
}

// Controlfields returns the unique set of controlfields for the record
// that match the specified tags. If no tags are specified (empty string)
// then all controlfields are returned.
func (rec Record) Controlfields(tags string) (f []*Controlfield) {
	if tags == "" {
		return rec.controlfields
	}

	uniq := make(map[string]bool)
	for _, t := range strings.Split(tags, ",") {
		for _, c := range rec.controlfields {
			if c.tag == t {
				ck := strings.Join([]string{c.tag, c.text}, ":")
				_, ok := uniq[ck]
				if !ok {
					f = append(f, c)
				}
				uniq[ck] = true
			}
		}
	}
	return f
}

// Controlfield returns the text value of the first control field for the
// record that matches the specified (presumably non-repeating) tag.
func (rec Record) Controlfield(tag string) string {
	for _, cf := range rec.controlfields {
		if cf.tag == tag {
			return cf.text
		}
	}
	return ""
}

func (cf Controlfield) Tag() string {
	return cf.tag
}

func (cf Controlfield) Text() string {
	return cf.text
}
