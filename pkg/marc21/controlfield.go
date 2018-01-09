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

//http://www.loc.gov/marc/bibliographic/bd00x.html

type CfValue struct {
	Code  string
	Label string
}

func cfShortCode(codeList map[string]string, b []byte, i int) (v CfValue) {

	var code, label string

	if len(b) > i {
		code = string(b[i])
		if code != "" {
			label, _ = codeList[code]
		}
	}

	return CfValue{Code: code, Label: label}
}

func cfWideCode(codeList map[string]string, b []byte, i, w int) (v CfValue) {

	var code, label string

	if len(b) > i+w {
		code = string(b[i : i+w])
		if code != "" {
			label, _ = codeList[code]
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

	c.ControlNumber = rec.getCF("001")
	c.ControlNumberIdentifier = rec.getCF("003")
	c.LatestTransDateTime = rec.getCF("005")

	b := []byte(rec.getCF("008"))
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
		label = "Books"
	}
	c.MaterialType = CfValue{Code: mt, Label: label}

	c.MaterialCharactor = make(CfMatlDesc)

	if len(b) > 18 {
		// Call the parse function with everything 18 and beyond
		// even though the bytes after 34 aren't material specific. The
		// parse functions will ignore any extra bytes and, even though
		// some 008 fields are "short", we'd still like to get as
		// much data as possible.
		switch mt {
		case "BK":
			c.MaterialCharactor.parseBookCf(b[18:])
		case "CF":
			c.MaterialCharactor.parseComputerFilesCf(b[18:])
		case "MP":
			c.MaterialCharactor.parseMapCf(b[18:])
		case "MU":
			c.MaterialCharactor.parseMusicCf(b[18:])
		case "CR":
			c.MaterialCharactor.parseContinuingResourcesCf(b[18:])
		case "VM":
			c.MaterialCharactor.parseVisualMaterialsCf(b[18:])
		case "MX":
			c.MaterialCharactor.parseMixedMaterialsCf(b[18:])
		}
	}

	cf006 := rec.getCFs("006")
	for _, x := range cf006 {
		b := []byte(x)
		if len(b) > 1 {

			// NOTE: The Material Form is b[0] and ignoring it *could*
			// disconnect the material form from the corresponding values.
			// However the material form *should* match the leader record
			// type so that really shouldn't be an issue... should it?
			// Perhaps a check to verify that b[0] matches?

			switch mt {
			case "BK":
				c.MaterialCharactor.parseBookCf(b[1:])
			case "CF":
				c.MaterialCharactor.parseComputerFilesCf(b[1:])
			case "MP":
				c.MaterialCharactor.parseMapCf(b[1:])
			case "MU":
				c.MaterialCharactor.parseMusicCf(b[1:])
			case "CR":
				c.MaterialCharactor.parseContinuingResourcesCf(b[1:])
			case "VM":
				c.MaterialCharactor.parseVisualMaterialsCf(b[1:])
			case "MX":
				c.MaterialCharactor.parseMixedMaterialsCf(b[1:])
			}
		}
	}

	// Each 007 needs to hang as a unit.
	cf007 := rec.getCFs("007")
	for _, x := range cf007 {
		b := []byte(x)
		if len(b) > 1 {
			pd := make(CfPhysDesc)

			pd["MaterialCategory"] = cfShortCode(materialCategory, b, 0)

			switch string(b[0]) {
			case "a":
				pd.parsePdA(b) // "Map",
			case "c":
				pd.parsePdC(b) // "Electronic resource",
			case "d":
				pd.parsePdD(b) // "Globe",
			case "f":
				pd.parsePdF(b) // "Tactile material",
			case "g":
				pd.parsePdG(b) // "Projected graphic",
			case "h":
				pd.parsePdH(b) // "Microform",
			case "k":
				pd.parsePdK(b) // "Nonprojected graphic",
			case "m":
				pd.parsePdM(b) // "Motion picture",
			case "o":
				pd.parsePdO(b) // "Kit",
			case "q":
				pd.parsePdQ(b) // "Notated music",
			case "r":
				pd.parsePdR(b) // "Remote-sensing image",
			case "s":
				pd.parsePdS(b) // "Sound recording",
			case "t":
				pd.parsePdT(b) // "Text",
			case "v":
				pd.parsePdV(b) // "Videorecording",
			case "z":
				pd.parsePdZ(b) // "Unspecified",
			}

			c.PhysicalDescription = append(c.PhysicalDescription, pd)
		}
	}
	return c
}

// getCF returns the control field for the specified non-repeating tag
func (rec Record) getCF(tag string) string {
	for _, cf := range rec.Controlfields {
		if cf.Tag == tag {
			return cf.Text
		}
	}
	return ""
}

// getCFs returns the unique control field(s) for the specified repeating tag
func (rec Record) getCFs(tag string) (t []string) {

	uniq := make(map[string]bool)

	for _, cf := range rec.Controlfields {
		if cf.Tag == tag {
			uniq[cf.Text] = true
		}
	}

	for k := range uniq {
		t = append(t, k)
	}

	return t
}

// MaterialType returns the code and description of the type of material
// documented by the record. {"Books", "Computer Files", "Maps", "Music",
// "Continuing Resources", "Visual Materials" or "Mixed Materials"}
func (rec Record) MaterialType() (code, label string) {
	rt, _ := rec.RecordType()
	bl, _ := rec.BibliographicLevel()

	switch rt {
	case "c", "d", "i", "j":
		return "MU", "Music"
	case "e", "f":
		return "MP", "Maps"
	case "g", "k", "o", "r":
		return "VM", "Visual Materials"
	case "m":
		return "CF", "Computer Files"
	case "p":
		return "MX", "Mixed Materials"
	case "t":
		switch bl {
		case "a", "c", "d", "m":
			return "BK", "Books"
		}
	case "a":
		switch bl {
		case "a", "c", "d", "m":
			return "BK", "Books"
		case "b", "i", "s":
			return "CR", "Continuing Resources"
		}
	}
	return code, label
}

// extractControlfields extracts the control fields from the raw MARC record bytes
func extractControlfields(rawRec []byte, baseAddress int, dir []*Directory) (cfs []*Controlfield, err error) {

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
	parseError := false

	for _, d := range dir {
		if strings.HasPrefix(d.Tag, "00") {

			start := baseAddress + d.StartingPos
			b := rawRec[start : start+d.FieldLength]

			if b[len(b)-1] == FieldTerminator {
				cfs = append(cfs, &Controlfield{Tag: d.Tag, Text: string(b[:len(b)-1])})
			} else {
				parseError = true
			}
		}
	}

	if parseError {
		log.Printf("Control fields extraction error: %s\n", cfs)
	}

	return cfs, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (cf Controlfield) String() string {
	return fmt.Sprintf("{%s: '%s'}", cf.Tag, cf.Text)
}
