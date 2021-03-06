// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"bytes"
	"errors"
	"strings"
)

/*
https://www.loc.gov/marc/specifications/specrecstruc.html

    Data fields in MARC 21 formats are assigned tags beginning with
    ASCII numeric characters other than two zeroes. Such fields contain
    indicators and subfield codes, as well as data and a field
    terminator. There are no restrictions on the number, length, or
    content of data fields other than those already stated or implied,
    e.g., those resulting from the limitation of total record length.

    Indicators are the first two characters in every variable data
    field, preceding any subfield code (delimiter plus data element
    identifier) which may be present. Each indicator is one character
    and every data field in the record includes two indicators, even if
    values have not been defined for the indicators in a particular
    field. Indicators supply additional information about the field,
    and are defined individually for each field. Indicator values are
    interpreted independently; meaning is not ascribed to the two
    indicators taken together. Indicators may be any ASCII lowercase
    alphabetic, numeric, or blank. A blank is used in an undefined
    indicator position, and may also have a defined meaning in a
    defined indicator position. The numeric character 9 is reserved for
    local definition as an indicator.

    Subfield codes identify the individual data elements within the
    field, and precede the data elements they identify. Each data field
    contains at least one subfield code. The subfield code consists of
    a delimiter (ASCII 1F (hex)) followed by a data element identifier.
    Data element identifiers defined in MARC 21 may be any ASCII
    lowercase alphabetic or numeric character.
*/

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Variable data fields - The remaining variable fields defined in the
    format. In addition to being identified by a field tag in the
    Directory, variable data fields contain two indicator positions
    stored at the beginning of each field and a two-character subfield
    code preceding each data element within the field.

    The variable data fields are grouped into blocks according to the
    first character of the tag, which with some exceptions identifies
    the function of the data within the record. The type of information
    in the field is identified by the remainder of the tag.

    Indicator positions - The first two character positions in the
    variable data fields that contain values which interpret or
    supplement the data found in the field. Indicator values are
    interpreted independently, that is, meaning is not ascribed to the
    two indicators taken together. Indicator values may be a lowercase
    alphabetic or a numeric character. A blank (ASCII SPACE),
    represented in this document as a #, is used in an undefined
    indicator position. In a defined indicator position, a blank may be
    assigned a meaning, or may mean no information provided.

    Subfield codes - Two characters that distinguish the data elements
    within a field which require separate manipulation. A subfield code
    consists of a delimiter (ASCII 1F hex), represented in this
    document as a $, followed by a data element identifier. Data
    element identifiers may be a lowercase alphabetic or a numeric
    character. Subfield codes are defined independently for each field;
    however, parallel meanings are preserved whenever possible (e.g.,
    in the 100, 400, and 600 Personal Name fields). Subfield codes are
    defined for purposes of identification, not arrangement. The order
    of subfields is generally specified by standards for the data
    content, such as cataloging rules.

Also:
    http://www.loc.gov/marc/holdings/hdintro.html
    http://www.loc.gov/marc/authority/adintro.html
    http://www.loc.gov/marc/classification/cdintro.html
    http://www.loc.gov/marc/community/ciintro.html

*/

// extractDatafields extracts the data fields/sub-fields from the raw MARC record bytes
func extractDatafields(rawRec []byte, baseAddress int, dir []*directoryEntry) (dfs []*Datafield, err error) {

	for _, de := range dir {
		if !strings.HasPrefix(de.tag, "00") {
			start := baseAddress + de.startingPos
			b := rawRec[start : start+de.fieldLength]

			if b[de.fieldLength-1] != fieldTerminator {
				return nil, errors.New("extractDatafields: Field terminator not found at end of field")
			}

			df := Datafield{
				Tag:  de.tag,
				Ind1: string(b[0]),
				Ind2: string(b[1]),
			}

			for _, t := range bytes.Split(b[2:de.fieldLength-1], []byte{delimiter}) {
				if len(t) > 0 {
					df.Subfields = append(df.Subfields, &Subfield{Code: string(t[0]), Text: string(t[1:])})
				}
			}
			dfs = append(dfs, &df)
		}
	}

	return dfs, nil
}

// GetDatafields returns datafields for the record that match the
// specified comma separated list of tags. If no tags are specified
// (empty string) then all datafields are returned
func (rec Record) GetDatafields(tags string) (dfs []*Datafield) {
	if tags == "" {
		return rec.Datafields
	}

	for _, t := range strings.Split(tags, ",") {
		for _, df := range rec.Datafields {
			if df.Tag == t {
				dfs = append(dfs, df)
			}
		}
	}
	return dfs
}

// GetSubfields returns subfields for the datafield that match the
// specified codes. If no codes are specified (empty string) then all
// subfields are returned
func (df Datafield) GetSubfields(codes string) (sfs []*Subfield) {
	if codes == "" {
		return df.Subfields
	}

	for _, c := range []byte(codes) {
		for _, sf := range df.Subfields {
			if sf.Code == string(c) {
				sfs = append(sfs, sf)
			}
		}
	}
	return sfs
}

// GetTag returns the tag for the datafield
func (df Datafield) GetTag() string {
	return df.Tag
}

// GetInd1 returns the indicator 1 value for the datafield
func (df Datafield) GetInd1() string {
	if df.Ind1 == "" {
		return " "
	}
	return df.Ind1
}

// GetInd2 returns the indicator 2 value for the datafield
func (df Datafield) GetInd2() string {
	if df.Ind2 == "" {
		return " "
	}
	return df.Ind2
}

// GetCode returns the code for the subfield
func (sf Subfield) GetCode() string {
	return sf.Code
}

// GetText returns the text for the subfield
func (sf Subfield) GetText() string {
	return sf.Text
}
