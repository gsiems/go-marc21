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
// http://www.loc.gov/marc/holdings/hd00x.html
// http://www.loc.gov/marc/authority/ad00x.html
// http://www.loc.gov/marc/classification/cd00x.html
// http://www.loc.gov/marc/community/ci00x.html

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
				cfs = append(cfs, &Controlfield{Tag: d.tag, Text: string(b[:len(b)-1])})
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
	return fmt.Sprintf("{%s: '%s'}", cf.Tag, cf.Text)
}

// GetControlfields returns the unique set of controlfields for the
// record that match the specified tags. If no tags are specified
// (empty string) then all controlfields are returned
func (rec Record) GetControlfields(tags string) (cfs []*Controlfield) {
	if tags == "" {
		return rec.Controlfields
	}

	uniq := make(map[string]bool)
	for _, t := range strings.Split(tags, ",") {
		for _, cf := range rec.Controlfields {
			if cf.Tag == t {
				ck := strings.Join([]string{cf.Tag, cf.Text}, ":")
				_, ok := uniq[ck]
				if !ok {
					cfs = append(cfs, cf)
				}
				uniq[ck] = true
			}
		}
	}
	return cfs
}

// GetControlfield returns the text value of the first control field for the
// record that matches the specified (presumably non-repeating) tag
func (rec Record) GetControlfield(tag string) string {
	for _, cf := range rec.Controlfields {
		if cf.Tag == tag {
			return cf.Text
		}
	}
	return ""
}

// GetTag returns the tag for the controlfield
func (cf Controlfield) GetTag() string {
	return cf.Tag
}

// GetText returns the text for the controlfield
func (cf Controlfield) GetText() string {
	return cf.Text
}
