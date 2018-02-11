// Use the results from parsing the LoC field list pages to create
// the code maps and functions for parsing/translating MARC controlfields

// 'tis ugly code that follows

package main

import (
	"fmt"
	"strings"
	//
	codegen "github.com/gsiems/go-marc21/codegen/pkg"
)

var htmlfiles = map[string]string{
	"Authority":      "input/ecadlist.html",
	"Bibliography":   "input/ecbdlist.html",
	"Classification": "input/eccdlist.html",
	"Community":      "input/eccilist.html",
	"Holdings":       "input/echdlist.html",
}

var subtagCodes = map[string]string{
	"004	DEFAULT": "",
	"006	BOOKS": "BK",
	"006	COMPUTER FILES/ELECTRONIC RESOURCES": "CF",
	"006	CONTINUING RESOURCES": "CR",
	"006	MAPS": "MP",
	"006	MIXED MATERIALS": "MX",
	"006	MUSIC": "MU",
	"006	VISUAL MATERIALS": "VM",
	"007	DEFAULT": "",
	"007	ELECTRONIC RESOURCE": "ELR",
	"007	GLOBE": "GLB",
	"007	KIT": "KIT",
	"007	MAP": "MAP",
	"007	MICROFORM": "MIC",
	"007	MOTION PICTURE": "MOP",
	"007	NONPROJECTED GRAPHIC": "NPG",
	"007	NOTATED MUSIC": "NMU",
	"007	PROJECTED GRAPHIC": "PRG",
	"007	REMOTE-SENSING IMAGE": "RSI",
	"007	SOUND RECORDING": "SOR",
	"007	TACTILE MATERIAL": "TAM",
	"007	TEXT": "TXT",
	"007	UNSPECIFIED": "UNS",
	"007	VIDEORECORDING": "VIR",
	"008	ALL MATERIALS": "",
	"008	BOOKS": "BK",
	"008	COMPUTER FILES": "CF",
	"008	CONTINUING RESOURCES": "CR",
	"008	DEFAULT": "",
	"008	MAPS": "MP",
	"008	MIXED MATERIALS": "MX",
	"008	MUSIC": "MU",
	"008	VISUAL MATERIALS": "VM",
	"009	DEFAULT": "",
}

func main() {

	fmt.Println("package marc21")
	fmt.Println()

	// Ensure that the order does not change from run to run
	fl := []string{
		"Authority",
		"Bibliography",
		"Classification",
		"Community",
		"Holdings",
	}

	for _, format := range fl {
		file := htmlfiles[format]
		formatBanner(format)
		cftags := codegen.ExtractCfStruct(file)

		for _, cftag := range validTags(cftags.Tags) {
			//tagBanner(format, cftag.Tag)

			if cftag.Tag == "006" && format == "Bibliography" {
				//tagBanner(format, cftag.Tag)
				//makeBibliography006FormOfMaterialList(format, cftag)
				continue
			}
			if cftag.Tag == "007" && (format == "Bibliography" || format == "Holdings") {
				tagBanner(format, cftag.Tag)
				make007CategoryOfMaterialList(format, cftag)
			}

			vst := validSubtags(cftag.Subtags)
			for _, cfsubtag := range vst {
				subtagBanner(format, cftag.Tag, cfsubtag.Label)

				stcode := subtagCodes[fmt.Sprintf("%s\t%s", cftag.Tag, cfsubtag.Label)]

				ve := validElements(cfsubtag.Elements)

				for _, cfelement := range ve {
					// bibliography006VMFormOfMaterial
					// bibliography007MAPCategoryOfMaterial
					if format == "Bibliography" && cftag.Tag == "006" && cfelement.CamelName == "FormOfMaterial" {
						continue
					}
					if cftag.Tag == "007" && cfelement.CamelName == "CategoryOfMaterial" {
						continue
					}

					if len(cfelement.LookupValues) > 0 && cfelement.FnType != "read" && cfelement.FnType != "range" {
						varname := strings.ToLower(format) + cftag.Tag + stcode + cfelement.CamelName
						if varname == "holdings008SpecificRetentionPolicy" {
							continue
						}
						makeLookupList(cfelement, varname)
					}
				}

				if cftag.Tag == "008" {
					//makeValidLookupFunc(format, cftag.Tag, cfsubtag)
					make008LookupFunc(format, cftag.Tag, cfsubtag)
				} else if cftag.Tag == "007" {
					//makeValidLookupFunc(format, cftag.Tag, cfsubtag)
					make007LookupFunc(format, cftag.Tag, cfsubtag)
				}
			}
		}
	}
}

func validTags(cftags []*codegen.CfTag) (f []*codegen.CfTag) {
	for _, c := range cftags {
		if !strings.Contains(c.Label, "OBSOLETE") {
			f = append(f, c)
		}
	}
	return f
}

func validSubtags(cfsubtags []*codegen.CfSubtag) (f []*codegen.CfSubtag) {
	for _, c := range cfsubtags {
		if !strings.Contains(c.Label, "OBSOLETE") {
			f = append(f, c)
		}
	}
	return f
}

func validElements(cfe []*codegen.CfElement) (f []*codegen.CfElement) {
	for _, c := range cfe {
		if !strings.Contains(c.Name, "OBSOLETE") {
			f = append(f, c)
		}
	}
	return f
}

func formatBanner(format string) {
	fmt.Println()
	bannerLine()
	fmt.Printf("// %s\n", format)
}

func tagBanner(format, tag string) {
	bannerLine()
	fmt.Printf("// %s -- %s \n", format, tag)
}

func subtagBanner(format, tag, subtag string) {
	bannerLine()
	if subtag != "DEFAULT" {
		fmt.Printf("// %s -- %s -- %s\n", format, tag, subtag)
	} else {
		fmt.Printf("// %s -- %s\n", format, tag)
	}
}

func bannerLine() {
	fmt.Println("////////////////////////////////////////////////////////////////////////")
}

func makeBibliography006FormOfMaterialList(format string, cftag *codegen.CfTag) {

	varname := strings.ToLower(format) + cftag.Tag + "FormOfMaterial"
	fmt.Printf("var %s = map[string]string{\n", varname)

	vst := validSubtags(cftag.Subtags)
	for _, cfsubtag := range vst {

		ve := validElements(cfsubtag.Elements)
		for _, cfe := range ve {
			if cftag.Tag == "006" && cfe.CamelName == "FormOfMaterial" {
				for _, lv := range cfe.LookupValues {
					if !strings.Contains(lv.Label, "OBSOLETE") {
						fmt.Printf("\t%q: %q,\n", lv.Code, lv.Label)
					}
				}
			}
		}
	}

	fmt.Println("}")
}

func make007CategoryOfMaterialList(format string, cftag *codegen.CfTag) {

	varname := strings.ToLower(format) + cftag.Tag + "CategoryOfMaterial"
	fmt.Printf("var %s = map[string]string{\n", varname)

	vst := validSubtags(cftag.Subtags)
	for _, cfsubtag := range vst {

		ve := validElements(cfsubtag.Elements)
		for _, cfe := range ve {
			if cftag.Tag == "007" && cfe.CamelName == "CategoryOfMaterial" {
				for _, lv := range cfe.LookupValues {
					if !strings.Contains(lv.Label, "OBSOLETE") {
						fmt.Printf("\t%q: %q,\n", lv.Code, lv.Label)
					}
				}
			}
		}
	}

	fmt.Println("}")
}

func makeLookupList(cfe *codegen.CfElement, varname string) {
	fmt.Printf("var %s = map[string]string{\n", varname)

	for _, lv := range cfe.LookupValues {
		if strings.Contains(lv.Label, "OBSOLETE") {
			continue
		}

		if lv.Code == "#" {
			fmt.Printf("\t%q: %q,\n", " ", lv.Label)
		} else if lv.Code == "##" {
			fmt.Printf("\t%q: %q,\n", "  ", lv.Label)
		} else if lv.Code == "###" {
			fmt.Printf("\t%q: %q,\n", "   ", lv.Label)
		} else {
			fmt.Printf("\t%q: %q,\n", lv.Code, lv.Label)
		}
	}
	fmt.Println("}")
}

/*
func makeValidLookupFunc(format, cftag string, cfsubtag *codegen.CfSubtag) {

	stcode := subtagCodes[fmt.Sprintf("%s\t%s", cftag, cfsubtag.Label)]

	funcName := strings.Join([]string{"valid", format, cftag, stcode, "Fields"}, "")

	fmt.Println()
	if cfsubtag.Label == "DEFAULT" {
		fmt.Printf(`// %s lists the valid fields for %s control field
// data for %s records data`, funcName, cftag, format)
	} else {
		fmt.Printf(`// %s lists the valid fields for %s control field
// data for %s records %s (%s) data`, funcName, cftag, format, cfsubtag.Label, stcode)
	}

	fmt.Println()
	fmt.Printf("func %s() []string {\n\n", funcName)

	fmt.Println("\tvalidFields := []string{")

	ve := validElements(cfsubtag.Elements)
	for _, e := range ve {
		fieldName := fmt.Sprintf("(%02d/%02d) %s", e.Offset, e.Width, e.Name)
		fmt.Printf("\t\t%q,\n", fieldName)
	}
	fmt.Println("\t}\n")
	fmt.Println("\treturn validFields")
	fmt.Println("}")
}
*/

func make007LookupFunc(format, cftag string, cfsubtag *codegen.CfSubtag) {

	stcode := subtagCodes[fmt.Sprintf("%s\t%s", cftag, cfsubtag.Label)]

	funcName := strings.Join([]string{"parse", format, cftag, stcode}, "")

	fmt.Println()
	if cfsubtag.Label == "DEFAULT" {
		fmt.Printf(`// %s parses the %s control field data for
// %s records data`, funcName, cftag, format)
	} else {
		fmt.Printf(`// %s parses the %s control field data for
// %s records %s (%s) data`, funcName, cftag, format, cfsubtag.Label, stcode)
	}

	fmt.Println()
	fmt.Printf("func %s(s string) (pd Cf007Desc) {\n\n", funcName)
	fmt.Println("\tpd = make(Cf007Desc)\n")

	ve := validElements(cfsubtag.Elements)
	for _, e := range ve {
		var varname string
		if e.CamelName == "CategoryOfMaterial" {
			varname = strings.ToLower(format) + cftag + e.CamelName

		} else {
			varname = strings.ToLower(format) + cftag + stcode + e.CamelName
		}

		fieldName := fmt.Sprintf("(%02d/%02d) %s", e.Offset, e.Width, e.Name)
		if len(e.LookupValues) > 0 && e.FnType == "lookup" {
			fmt.Printf("\tpd[%q] = codeLookup(%s, s, %d, %d)\n",
				fieldName, varname, e.Offset, e.Width)
		} else if e.FnType == "read" || e.FnType == "range" {
			fmt.Printf("\tpd[%q] = CodeValue{Code: pluckBytes(s, %d, %d), Label: \"\"}\n",
				fieldName, e.Offset, e.Width)
		} else if len(e.LookupValues) > 0 && e.FnType == "multi" {
			end := e.Offset + e.Width

			fmt.Println()
			if e.CodeWidth == 1 {
				j := 1
				for i := e.Offset; i < end; i++ {
					fn := fmt.Sprintf("%s - %d", fieldName, j)
					j++
					fmt.Printf("\tpd[%q] = codeLookup(%s, s, %d, %d)\n", fn, varname, i, e.CodeWidth)
				}
			} else {
				j := 1
				for i := e.Offset; i < end; i = i + e.CodeWidth {
					fn := fmt.Sprintf("%s - %d", fieldName, i)
					j++
					fmt.Printf("\tpd[%q] = codeLookup(%s, s, %d, %d)\n", fn, varname, i, e.CodeWidth)
				}
			}
			fmt.Println()
		} else if len(e.LookupValues) > 0 && e.FnType == "hybrid" {

			// Find the element that has the range. Code should have a
			// hyphen. We want the label
			for _, lv := range e.LookupValues {
				if strings.Contains(lv.Code, "-") {

					vn := fmt.Sprintf("rt%02d", e.Offset)

					fmt.Println()
					fmt.Printf("\t%s := codeLookup(%s, s, %d, %d)\n", vn, varname, e.Offset, e.CodeWidth)
					fmt.Printf("\tif %s.Code != \"\" && %s.Label == \"\" {\n", vn, vn)
					fmt.Printf("\t\t%s.Label = %q\n", vn, lv.Label)
					fmt.Println("\t}")
					fmt.Printf("\tpd[%q] = %s\n", fieldName, vn)
					fmt.Println()

					break
				}
			}
		}
	}
	fmt.Println("\n\treturn pd")
	fmt.Println("}\n")
}

func make008LookupFunc(format, cftag string, cfsubtag *codegen.CfSubtag) {

	stcode := subtagCodes[fmt.Sprintf("%s\t%s", cftag, cfsubtag.Label)]

	funcName := strings.Join([]string{"parse", format, cftag, stcode}, "")

	fmt.Println()
	if cfsubtag.Label == "DEFAULT" {
		fmt.Printf(`// %s parses the %s control field data for
// %s records data`, funcName, cftag, format)
	} else {
		fmt.Printf(`// %s parses the %s control field data for
// %s records %s (%s) data`, funcName, cftag, format, cfsubtag.Label, stcode)
	}

	fmt.Println()
	fmt.Printf("func %s(d *Cf008Desc, s string) {\n\n", funcName)

	ve := validElements(cfsubtag.Elements)
	for _, e := range ve {
		varname := strings.ToLower(format) + cftag + stcode + e.CamelName
		if varname == "holdings008SpecificRetentionPolicy" {
			continue
		}

		// Adjust the offset for bibliography records so that the bib
		// functions can be called using both 008 and 006 data...
		offsetAdj := 0
		if format == "Bibliography" && cfsubtag.Label != "ALL MATERIALS" {
			offsetAdj = 18
		}

		fieldName := fmt.Sprintf("(%02d/%02d) %s", e.Offset, e.Width, e.Name)
		if len(e.LookupValues) > 0 && e.FnType == "lookup" {
			fmt.Printf("\td.append(%q, codeLookup(%s, s, %d, %d))\n",
				fieldName, varname, e.Offset-offsetAdj, e.Width)
		} else if e.FnType == "read" || e.FnType == "range" {
			fmt.Printf("\td.append(%q, CodeValue{Code: pluckBytes(s, %d, %d), Label: \"\"} )\n",
				fieldName, e.Offset-offsetAdj, e.Width)
		} else if len(e.LookupValues) > 0 && e.FnType == "multi" {
			end := e.Offset - offsetAdj + e.Width

			fmt.Println()
			if e.CodeWidth == 1 {
				fmt.Printf("\tfor i := %d; i < %d; i++ {\n", e.Offset-offsetAdj, end)
			} else {
				fmt.Printf("\tfor i := %d; i < %d; i = i + %d {\n", e.Offset-offsetAdj, end, e.CodeWidth)
			}
			fmt.Printf("\t\td.append(%q, codeLookup(%s, s, i, %d))\n", fieldName, varname, e.CodeWidth)
			fmt.Println("\t}")
			fmt.Println()
		} else if len(e.LookupValues) > 0 && e.FnType == "hybrid" {

			// Find the element that has the range. Code should have a
			// hyphen (or be enclosed in square braces). We want the label
			for _, lv := range e.LookupValues {
				if strings.Contains(lv.Code, "-") || strings.Contains(lv.Code, "[") {

					vn := fmt.Sprintf("rt%02d", e.Offset)

					fmt.Println()
					fmt.Printf("\t%s := codeLookup(%s, s, %d, %d)\n", vn, varname, e.Offset-offsetAdj, e.CodeWidth)
					fmt.Printf("\tif %s.Code != \"\" && %s.Label == \"\" {\n", vn, vn)
					fmt.Printf("\t\t%s.Label = %q\n", vn, lv.Label)
					fmt.Println("\t}")
					fmt.Printf("\td.append(%q, %s)\n", fieldName, vn)
					fmt.Println()

					break
				}
			}
		} else if len(e.LookupValues) > 0 && e.FnType == "hybrid-date" {

			vn := fmt.Sprintf("rt%02d", e.Offset)

			fmt.Println()
			fmt.Printf("\t%s := codeLookup(%s, s, %d, 1)\n", vn, varname, e.Offset-offsetAdj)
			fmt.Printf("\tif %s.Label == \"\" {\n", vn)
			fmt.Printf("\t\t%s = CodeValue {Code: pluckBytes(s, %d, %d), Label: \"Date\"}\n", vn, e.Offset-offsetAdj, e.Width)
			fmt.Println("\t}")
			fmt.Printf("\td.append(%q, %s)\n", fieldName, vn)
			fmt.Println()

		}
	}

	fmt.Println("}\n")
}
