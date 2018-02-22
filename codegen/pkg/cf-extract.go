// Here be some ugliness
//
// Parses a LoC MARC21 fields list page and extracts the control field
// definitions contained within

package codegen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// CfTags is a list of Control field definitions
type CfTags struct {
	Tags []*CfTag
}

// CfTag is a Control field definition
type CfTag struct {
	Tag     string
	Label   string
	Subtags []*CfSubtag
}

// CfSubtag is a Control field *subtag* definition (not that all control
// fields have subtags (in which case the subtag label is "DEFAULT"))
type CfSubtag struct {
	Label    string
	Elements []*CfElement
}

// CfElement is the definition for one or more characters that make up
// an "element"
type CfElement struct {
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
func (t CfTag) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("%s - %s", t.Tag, t.Label))

	if len(t.Subtags) == 0 {
		lines = append(lines, "")
	}

	for _, v := range t.Subtags {

		if v.Label != "DEFAULT" {
			lines = append(lines, fmt.Sprintf("   %s--%s", t.Tag, v.Label))
		}

		if len(v.Elements) > 0 {
			lines = append(lines, "     Character Positions")
		}

		for _, v2 := range v.Elements {

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
	}
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}

// ExtractCfStruct extracts the structure for control file entries as
// defined in the [saved-to-disc] fields list webpage
func ExtractCfStruct(filename string) (tags CfTags) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	notInCfBlock := true

	cfTg := new(CfTag)
	cfSt := new(CfSubtag)
	cfEl := new(CfElement)
	cfLv := new(LookupValue)

	reElem := regexp.MustCompile("^     [ ]{0,1}[0-9][0-9]")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// We are only interested in looking at control field definitions
		notInCfBlock = isNotInCfBlock(line, notInCfBlock)
		if notInCfBlock {
			continue
		}

		if canIgnoreCfLine(line) {
			continue
		}

		//// Check if this is a "NEW TAG" line
		if pluckBytes(line, 0, 2) == "00" {

			if cfTg.Tag != "" {
				if cfSt.Label != "" {
					if cfEl.Name != "" {
						cfSt.Elements = append(cfSt.Elements, cfEl)
						cfEl = new(CfElement)
					}
					cfTg.Subtags = append(cfTg.Subtags, cfSt)
					cfSt = new(CfSubtag)
				}

				tags.Tags = append(tags.Tags, cfTg)
				cfTg = new(CfTag)
			}

			tl := strings.SplitN(line, "-", 2)
			cfTg.Tag = strings.TrimSpace(tl[0])
			if len(tl) > 1 {
				cfTg.Label = strings.TrimSpace(tl[1])
			}
			continue
		}

		// Just in case we are not in a tag block
		if cfTg.Tag == "" {
			continue
		}

		//// Check if this is a "NEW SUBTAG" line
		if strings.Contains(line, cfTg.Tag+"--") {
			if cfSt.Label != "" {
				if cfEl.Name != "" {
					cfSt.Elements = append(cfSt.Elements, cfEl)
					cfEl = new(CfElement)
				}
				cfTg.Subtags = append(cfTg.Subtags, cfSt)
				cfSt = new(CfSubtag)
			}

			x := strings.SplitN(line, "--", 2)
			cfSt.Label = x[len(x)-1]
			continue
		}

		//// Check if this is a "NEW ELEMENT" line
		if reElem.MatchString(line) {

			if cfSt.Label == "" {
				cfSt.Label = "DEFAULT"
			}

			if cfEl.Name != "" {
				if cfLv.Code != "" {
					cfEl.LookupValues = append(cfEl.LookupValues, cfLv)
					cfLv = new(LookupValue)
				}
				cfSt.Elements = append(cfSt.Elements, cfEl)
				cfEl = new(CfElement)
			}

			line = strings.TrimSpace(line)

			offset, err := toInt([]byte(pluckBytes(line, 0, 2)))
			if err != nil {
				cfEl.Name = "BAD_PARSE"
				log.Printf("BAD PARSE: %s, %q, %q\n", cfTg.Label, cfSt.Label, line)
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
			// There is usually a hyphen here, but not always...
			if strings.Index(line, "-") == 0 {
				line = strings.TrimSpace(line[1:])
			}

			cfEl = new(CfElement)
			cfEl.Offset = offset
			cfEl.Width = width
			cfEl.Name = line
			cfEl.CamelName = camelizeString(line)
			cfEl.FnType = "read"
			continue
		}

		//// At this point, we assert that this is a "LOOKUP VALUE" line

		// Determine if this line is a continuation line from the previous
		if isWrappedLine(cfEl.wrapIndent, line) {
			il := len(cfEl.LookupValues) - 1
			cfEl.LookupValues[il].Label += " " + strings.TrimSpace(line)
			continue
		}

		lv := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(lv) == 1 {
			log.Printf("BAD PARSE: %s, %q, %q, %q\n", cfTg.Label, cfSt.Label, cfEl.Name, line)
			continue
		}

		// We need to check and account for lookup definitions that do
		// not match the typical " ... xx - label" pattern such as:
		// "         1-9 Number of units"     // code is a range of values
		// or
		// "         --- - Unknown"           // the dashes ARE the code
		// or
		// "         |- No attempt to code"   // no space before the separating dash

		code := lv[0]
		label := lv[1]
		if pluckBytes(code, 0, 1) != "-" {
			// line is not like: "         --- - Unknown"
			code = strings.TrimRight(code, "- ")
		}
		label = strings.TrimLeft(label, "- ")

		// On the first item in the list we want to determine how
		// "wide" the codes are and how much indentation to check for
		// when testing for line-wraps
		if len(cfEl.LookupValues) == 0 {

			// Determine how much indent qualifies as a line-wrap
			cfEl.wrapIndent = strings.Index(line, "-") + 1

			cfEl.CodeWidth = calcCodeWidth(code)

			if isRangeCode(code) {
				cfEl.FnType = "range"
			} else {
				cfEl.FnType = "lookup"
			}
		}

		if isRangeCode(code) {
			if cfEl.FnType == "lookup" || cfEl.FnType == "hybrid" {
				cfEl.FnType = "hybrid"
			} else if cfEl.CodeWidth < cfEl.Width && strings.Contains(label, "Date") {
				cfEl.FnType = "hybrid-date"
			} else {
				cfEl.FnType = "range"
			}
		} else if cfEl.FnType != "hybrid-date" {
			if cfEl.FnType == "range" {
				cfEl.FnType = "hybrid"
			} else if cfEl.CodeWidth < cfEl.Width {
				cfEl.FnType = "multi"
			} else if strings.Contains(code, "[") {
				cfEl.FnType = "hybrid"
			}
		}

		e := new(LookupValue)
		e.Code = code
		e.Label = label

		cfEl.LookupValues = append(cfEl.LookupValues, e)
	}

	if cfSt.Label != "" {
		if cfEl.Name != "" {
			cfSt.Elements = append(cfSt.Elements, cfEl)
		}
		cfTg.Subtags = append(cfTg.Subtags, cfSt)
	}
	tags.Tags = append(tags.Tags, cfTg)

	return tags
}

// isRangeCode checks the code to determine if is defining a range
func isRangeCode(code string) bool {
	if pluckBytes(code, 0, 1) != "-" {
		if strings.Index(code, "-") > 0 {
			return true
		}
	}
	return false
}

func canIgnoreCfLine(line string) bool {
	// For the few lines that we really want to ignore
	if pluckByte(line, 0) == "#" {
		return true
	}

	// Toss the few extra comments also
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

	return false
}

func isNotInCfBlock(line string, notInCfBlock bool) bool {

	// Start looking at control field definitions
	if pluckBytes(line, 0, 18) == "--Control Fields (" {
		return false
	}

	// Stop looking at control field definitions
	if pluckBytes(line, 0, 8) == "--Number" {
		return true
	}

	return notInCfBlock
}
