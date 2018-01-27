// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"errors"
	"strconv"
)

// Utility and helper functions

// shortCode extracts a single-character from a string
func shortCode(s string, i int) (code string) {

	if len(s) > i {
		code = string(s[i])
	}

	return code
}

// shortCodeLookup performs lookups on single-character reference tables (maps)
// and returns the code and, if found, the descriptive label for the code.
func shortCodeLookup(codeList map[string]string, s string, i int) (code, label string) {

	code = shortCode(s, i)

	if code != "" {
		label = codeList[code]
	}

	return code, label
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
