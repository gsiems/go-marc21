// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

//import "strconv"

/*
https://www.loc.gov/marc/specifications/specrecstruc.html

    A directory entry in MARC 21 is made up of a tag, length-of-field,
    and field starting position. The directory begins in character
    position 24 of the record and ends with a field terminator. It is
    of variable length and consists of a series of fixed fields,
    referred to as "entries." One entry is associated with each
    variable field (control or data) present in the record. Each
    directory entry is 12 characters in length;...
*/

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Directory - A series of entries that contain the tag, length, and
    starting location of each variable field within a record. Each
    entry is 12 character positions in length. Directory entries for
    variable control fields appear first, sequenced by the field tag in
    increasing numerical order. Entries for variable data fields
    follow, arranged in ascending order according to the first
    character of the tag. The stored sequence of the variable data
    fields in a record does not necessarily correspond to the order of
    the corresponding Directory entries. Duplicate tags are
    distinguished only by the location of the respective fields within
    the record. The Directory ends with a field terminator character
    (ASCII 1E hex).
*/

/*
http://www.loc.gov/marc/bibliographic/bddirectory.html

    CHARACTER POSITIONS

    00-02 - Tag
        Three ASCII numeric or ASCII alphabetic characters (upper case
        or lower case, but not both) that identify an associated
        variable field.

    03-06 - Field length
        Four ASCII numeric characters that specify the length of the
        variable field, including indicators, subfield codes, data, and
        the field terminator. A Field length number of less than four
        digits is right justified and unused positions contain zeros.

    07-11 - Starting character position
        Five ASCII numeric characters that specify the starting
        character position of the variable field relative to the Base
        address of data (Leader/12-16) of the record. A Starting
        character position number of less than five digits is right
        justified and unused positions contain zeros.
*/

// parseDirectory extracts the directory from the raw MARC record bytes
func parseDirectory(r []byte) (dir []*directoryEntry, err error) {

	for i := leaderLen; r[i] != fieldTerminator; i += 12 {
		var de directoryEntry

		de.tag = string(r[i : i+3])
		de.fieldLength, err = toInt(r[i+3 : i+7]) //strconv.Atoi(string(r[i+3 : i+7]))
		if err != nil {
			return nil, err
		}

		de.startingPos, err = toInt(r[i+7 : i+12]) // strconv.Atoi(string(r[i+7 : i+12]))
		if err != nil {
			return nil, err
		}

		dir = append(dir, &de)
	}
	return dir, nil
}
