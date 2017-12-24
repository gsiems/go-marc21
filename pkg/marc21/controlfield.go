// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	//"errors"
	//"fmt"
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

// TODO: validate control fields?
// TODO: ensure there are no duplicate control fields?
// TODO: create/update control fields for new/update record?
//          005 -> last updated: yyyymmddhhmmss.f
// TODO: parse/translate 006, 007, 008

func parseControlfields(rawRec []byte, baseAddress int, dir []*directory) (cfs []Controlfield, err error) {

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
		if strings.HasPrefix(d.tag, "00") {

			var cf Controlfield
			cf.Tag = d.tag

			start := baseAddress + d.startingPos
			b := rawRec[start : start+d.fieldLength]

			if b[len(b)-1] == fieldTerminator {
				cf.Text = string(b[:len(b)-1])
			} else {
				parseError = true
			}
			cfs = append(cfs, cf)
		}
	}

	if parseError {
		log.Printf("Control fields parsing error: %s\n", cfs)
	}

	return cfs, nil
}
