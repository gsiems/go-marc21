// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"errors"
	"strconv"
)

// Utility and helper functions

// shortCodeLookup performs lookups on single-character reference tables (maps)
// and returns the code and, if found, the descriptive label for the code.
func shortCodeLookup(codeList map[string]string, s string, i int) (code, label string) {

	code = pluckByte(s, i)

	if code != "" {
		label = codeList[code]
	}

	return code, label
}

// pluckByte extracts a single-byte from a string and returns
// the string result.
func pluckByte(b string, i int) (s string) {

	if len(b) > i {
		s = string(b[i])
	}

	return s
}

// toInt converts a byte array of digits to its corresponding integer
// value
func toInt(b []byte) (ret int, err error) {
	ret, err = strconv.Atoi(string(b))
	if err != nil {

		var digits = map[string]int{
			"0": 0,
			"1": 1,
			"2": 2,
			"3": 3,
			"4": 4,
			"5": 5,
			"6": 6,
			"7": 7,
			"8": 8,
			"9": 9,
		}

		ret = 0
		for i := range b {
			x, ok := digits[string(b[i])]
			if !ok {
				return 0, errors.New("toInt(): Not an integer")
			}
			ret = (10 * ret) + x
		}
	}
	return ret, nil
}
