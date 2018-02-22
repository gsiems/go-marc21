package codegen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Ldr is the collection of elements that make up a leader definition
type Ldr struct {
	Elements []*LdrElement
}

// LdrElement is the definition for one or more characters that make up
// an "element"
type LdrElement struct {
	Name         string
	CamelName    string
	Offset       int
	Width        int
	wrapIndent   int
	CodeWidth    int
	FnType       string
	LookupValues []*LookupValue
}

// String should return a string that matches, excepting some whitespace
// and cleaned-up linewraps, the input data. This is primarily intended
// for testing that the parsing/extraction is working correctly.
func (ldr Ldr) String() string {
	var lines []string

	if len(ldr.Elements) > 0 {
		lines = append(lines, "     Character Positions")
	}

	for _, v2 := range ldr.Elements {

		offset := fmt.Sprintf("%02d", v2.Offset)
		if v2.Width > 1 {
			offset = fmt.Sprintf("%s-%02d", offset, v2.Offset+v2.Width-1)
		}

		lines = append(lines, fmt.Sprintf("      %s - %s", offset, v2.Name))

		// codes/labels
		for _, cv := range v2.LookupValues {
			lines = append(lines, fmt.Sprintf("         %s - %s", cv.Code, cv.Label))
		}
	}
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}

// ExtractLdrStruct extracts the structure for leader entry as
// defined in the [saved-to-disc] fields list webpage
func ExtractLdrStruct(filename string) (ldr Ldr) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	notInLdrBlock := true

	//ldr = new(Ldr)
	LdrEl := new(LdrElement)
	LdrLv := new(LookupValue)

	reElem := regexp.MustCompile("^     [ ]{0,1}[0-9][0-9]")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// We are only interested in looking at control field definitions
		notInLdrBlock = isNotInLdrBlock(line, notInLdrBlock)
		if notInLdrBlock {
			continue
		}

		if canIgnoreLdrLine(line) {
			continue
		}

		//// Check if this is a "NEW ELEMENT" line
		if reElem.MatchString(line) {

			if LdrEl.Name != "" {
				if LdrLv.Code != "" {
					LdrEl.LookupValues = append(LdrEl.LookupValues, LdrLv)
					LdrLv = new(LookupValue)
				}
				ldr.Elements = append(ldr.Elements, LdrEl)
				LdrEl = new(LdrElement)
			}

			line = strings.TrimSpace(line)

			offset, err := toInt([]byte(pluckBytes(line, 0, 2)))
			if err != nil {
				LdrEl.Name = "BAD_PARSE"
				log.Printf("BAD PARSE: %q\n", line)
				continue
			}

			width, err := toInt([]byte(pluckBytes(line, 3, 2)))
			if err == nil {
				width = width - offset + 1
				line = strings.TrimSpace(line[5:])
			} else {
				width = 1
				line = strings.TrimSpace(line[2:])
			}
			// There is usually a hyphen here, but just in case...
			if strings.Index(line, "-") == 0 {
				line = strings.TrimSpace(line[1:])
			}

			LdrEl = new(LdrElement)
			LdrEl.Offset = offset
			LdrEl.Width = width
			LdrEl.Name = line
			LdrEl.CamelName = camelizeString(line)
			LdrEl.FnType = "read"
			continue
		}

		//// At this point, we assert that this is a "LOOKUP VALUE" line

		// Determine if this line is a continuation line from the previous
		if isWrappedLine(LdrEl.wrapIndent, line) {
			il := len(LdrEl.LookupValues) - 1
			LdrEl.LookupValues[il].Label += " " + strings.TrimSpace(line)
			continue
		}

		lv := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(lv) == 1 {
			log.Printf("BAD PARSE: %q, %q\n", LdrEl.Name, line)
			continue
		}

		code := lv[0]
		label := lv[1]

		if pluckBytes(code, 0, 1) != "-" {
			// line is not like: "         - - Unknown"
			code = strings.TrimRight(code, "- ")
		}
		label = strings.TrimLeft(label, "- ")

		// On the first item in the list we want to determine how
		// "wide" the codes are and how much indentation to check for
		// when testing for line-wraps
		if len(LdrEl.LookupValues) == 0 {

			// Determine how much indent qualifies as a line-wrap
			LdrEl.wrapIndent = strings.Index(line, "-") + 1

			LdrEl.CodeWidth = calcCodeWidth(code)
			LdrEl.FnType = "lookup"
		}

		e := new(LookupValue)
		e.Code = code
		e.Label = label

		LdrEl.LookupValues = append(LdrEl.LookupValues, e)

	}

	if LdrEl.Name != "" {
		if LdrLv.Code != "" {
			LdrEl.LookupValues = append(LdrEl.LookupValues, LdrLv)
		}
		ldr.Elements = append(ldr.Elements, LdrEl)
	}

	return ldr
}

func canIgnoreLdrLine(line string) bool {
	// For the few lines that we really want to ignore
	if pluckByte(line, 0) == "#" {
		return true
	}

	// Toss any extra comments also
	if pluckBytes(line, 0, 2) == "--" {
		return true
	}

	// Empty lines.
	if len(strings.TrimSpace(line)) < 1 {
		return true
	}

	if strings.TrimSpace(strings.ToUpper(line)) == "CHARACTER POSITIONS" {
		return true
	}

	if strings.TrimSpace(strings.ToUpper(line)) == "LEADER" {
		return true
	}

	return false
}

func isNotInLdrBlock(line string, notInLdrBlock bool) bool {

	// Start looking at leader definition
	if strings.TrimSpace(line) == "LEADER" {
		return false
	}

	// Stop looking at leader definition
	if strings.TrimSpace(line) == "DIRECTORY" {
		return true
	}

	return notInLdrBlock
}
