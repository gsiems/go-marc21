// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import "strconv"

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
*/

/*
http://www.loc.gov/marc/bibliographic/bdleader.html

    Character Positions
    00-04 - Record length

    05 - Record status
        a - Increase in encoding level
        c - Corrected or revised
        d - Deleted
        n - New
        p - Increase in encoding level from prepublication

    06 - Type of record
        a - Language material
        c - Notated music
        d - Manuscript notated music
        e - Cartographic material
        f - Manuscript cartographic material
        g - Projected medium
        i - Nonmusical sound recording
        j - Musical sound recording
        k - Two-dimensional nonprojectable graphic
        m - Computer file
        o - Kit
        p - Mixed materials
        r - Three-dimensional artifact or naturally occurring object
        t - Manuscript language material

    07 - Bibliographic level
        a - Monographic component part
        b - Serial component part
        c - Collection
        d - Subunit
        i - Integrating resource
        m - Monograph/Item
        s - Serial

    08 - Type of control
        # - No specified type
        a - Archival

    09 - Character coding scheme
        # - MARC-8
        a - UCS/Unicode

    10 - Indicator count
        2 - Number of character positions used for indicators

    11 - Subfield code count
        2 - Number of character positions used for a subfield code

    12-16 - Base address of data
        [number] - Length of Leader and Directory

    17 - Encoding level
        # - Full level
        1 - Full level, material not examined
        2 - Less-than-full level, material not examined
        3 - Abbreviated level
        4 - Core level
        5 - Partial (preliminary) level
        7 - Minimal level
        8 - Prepublication level
        u - Unknown
        z - Not applicable

    18 - Descriptive cataloging form
        # - Non-ISBD
        a - AACR 2
        c - ISBD punctuation omitted
        i - ISBD punctuation included
        n - Non-ISBD punctuation omitted
        u - Unknown

    19 - Multipart resource record level
        # - Not specified or not applicable
        a - Set
        b - Part with independent title
        c - Part with dependent title

    20 - Length of the length-of-field portion
        4 - Number of characters in the length-of-field portion of a
            Directory entry

    21 - Length of the starting-character-position portion
        5 - Number of characters in the starting-character-position
            portion of a Directory entry

    22 - Length of the implementation-defined portion
        0 - Number of characters in the implementation-defined portion
            of a Directory entry

    23 - Undefined
        0 - Undefined
*/

/*
Example:

    01166cam  2200313   450000100
    01166   | RecordLength           int   |
    c       | RecordStatus           byte  | c - Corrected or revised
    a       | RecordType             byte  | a - Language material
    m       | BibliographicLevel     byte  | a - Monographic component part
            | ControlType            byte  | # - No specified type
            | CharacterCodingScheme  byte  | # - MARC-8
    2       | IndicatorCount         byte  | 2 - Number of character positions used for indicators
    2       | SubfieldCodeCount      byte  | 2 - Number of character positions used for a subfield code
    00313   | BaseDataAddress        int   |
            | EncodingLevel          byte  | # - Full level
            | CatalogingForm         byte  | # - Non-ISBD
            | MultipartLevel         byte  | # - Not specified or not applicable
    4       | LenOfLengthOfField     byte  | 4 - Number of characters in the length-of-field portion of a Directory entry
    5       | LenOfStartCharPosition byte  | 5 - Number of characters in the starting-character-position portion of a Directory entry
    0       | LenOfImplementDefined  byte  | 0 - Number of characters in the implementation-defined portion of a Directory entry
    0       | Undefined              byte  | 0 - Undefined
*/

// TODO: Do we want to/have use for leader validation?

// parseLeader extracts the leader information from the raw MARC record bytes
func parseLeader(b []byte) (l *Leader, err error) {
	l = new(Leader)

	l.RecordLength, err = strconv.Atoi(string(b[0:4]))
	if err != nil {
		return nil, err
	}
	l.RecordStatus = b[5]
	l.RecordType = b[6]
	l.BibliographicLevel = b[7]
	l.ControlType = b[8]
	l.CharacterCodingScheme = b[9]
	l.IndicatorCount = b[10]
	l.SubfieldCodeCount = b[11]
	l.BaseDataAddress, err = strconv.Atoi(string(b[12:17]))
	if err != nil {
		return nil, err
	}
	l.EncodingLevel = b[17]
	l.CatalogingForm = b[18]
	l.MultipartLevel = b[19]
	l.LenOfLengthOfField = b[20]
	l.LenOfStartCharPosition = b[21]
	l.LenOfImplementDefined = b[22]
	l.Undefined = b[23]

	return l, nil
}
