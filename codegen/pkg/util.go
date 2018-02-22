package codegen

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// LookupValue is the definition of a single lookup value for an element
type LookupValue struct {
	Code  string
	Label string
}

// camelizeString removes all non-alpha-numerics from a string and
// upper-camel cased the results
func camelizeString(s string) (c string) {
	re := regexp.MustCompile("[[:^alnum:]]+")
	c = strings.Replace(strings.Title(re.ReplaceAllString(s, " ")), " ", "", -1)
	return c
}

// pluckBytes plucks the value of one or more contiguous bytes from
// a string
func pluckBytes(b string, i, w int) (s string) {

	if w == 0 {
		return pluckByte(b, i)
	}
	if len(b) > i+w {
		s = b[i : i+w]
	}
	return s
}

// pluckByte plucks the value of a single byte from a string
func pluckByte(b string, i int) (s string) {
	if len(b) > i {
		s = string(b[i])
	}
	return s
}

// toInt converts a string "integer" to a real integer
func toInt(b []byte) (ret int, err error) {
	ret, err = strconv.Atoi(string(b))
	if err != nil {

		var digits = map[string]int{
			"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
			"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
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

// Determines if the current line is a continuation line of the previous line
func isWrappedLine(minIndent int, line string) (isWrap bool) {
	if minIndent > 0 {
		t1 := strings.TrimSpace(pluckBytes(line, 0, minIndent))
		if t1 == "" {
			return true
		}
	}
	return false
}

// calcCodeWidth determines the width of the lookup code (accounting
// for hyphenations)
func calcCodeWidth(c string) int {
	if pluckBytes(c, 0, 1) == "-" {
		return len(c)
	}
	e := strings.SplitN(c, "-", 2)
	return len(e[0])
}
