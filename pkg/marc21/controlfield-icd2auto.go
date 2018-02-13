package marc21

// The sole purpose of this file is to provide an interface to the
// controlfield-auto.go contents.
// * Any code not needed for that purpose does not go here.
// * Any code that is only needed for that purpose does go here.

// CodeValue contains a code and it's corresponding descriptive label
// for a controlfield entry.
type CodeValue struct {
	Code   string
	Label  string
	Offset int
	Width  int
}

// Cf008Desc is the structure for holding the description from the
// parsing of the 008 field
type Cf008Desc map[string][]CodeValue

// Cf007Desc is the structure for holding the description from the
// parsing of a 007 field
type Cf007Desc map[string]CodeValue

func (mc Cf008Desc) append(k string, c CodeValue) {

	// Ensure that the code/label hasn't already been appended (no duplicate code/labels)
	exists := false
	for _, x := range mc[k] {
		if x.Code == c.Code {
			exists = true
			break
		}
	}

	if !exists {
		mc[k] = append(mc[k], c)
	}
}

func codeLookup(codeList map[string]string, b string, i, w int) (code, label string) {

	code = pluckBytes(b, i, w)

	if code != "" {
		label = codeList[code]
	}

	return code, label
}

// Parse006 parses the 006 controlfield for a record and returns a,
// hopefully, human readable translation of the field contents.
func (rec Record) Parse006() (d Cf008Desc) {

	d = make(Cf008Desc)

	type fn func(md *Cf008Desc, s string)

	m := map[string]fn{
		"BK": parseBibliography008BK,
		"CF": parseBibliography008CF,
		"MP": parseBibliography008MP,
		"MU": parseBibliography008MU,
		"CR": parseBibliography008CR,
		"VM": parseBibliography008VM,
		"MX": parseBibliography008MX,
	}

	if rec.RecordFormat() == Bibliography {

		c, _ := rec.BibliographyMaterialType()

		// Ref: http://www.loc.gov/marc/bibliographic/bd006.html
		//
		// "Except for code s (Serial/Integrating resource), the codes
		// in field 006/00 correspond to those in Leader/06 (Type of
		// record). For each occurrence of field 006, the codes
		// defined for character positions 01-17 will be the same as
		// those defined in the corresponding field 008, character
		// positions 18-34."

		// For bibliography 008 fields pass subslice of s -- s[18:]
		// and for bibliography 006 fields pass subslice of s -- s[1:]
		// This way the same parsing functions can parse both 008 and
		// 006 control fields

		cf := rec.GetControlfields("006")

		for _, cf6 := range cf {
			//c6 := pluckByte(cf6.Text, 0)
			//if c6 == c || c6 == "s" || c == "s" {

			code, label := codeLookup(bibliography006FormOfMaterial, cf6.Text, 0, 1)
			d.append("(00/01) Form of material", CodeValue{Code: code, Label: label, Offset: 0, Width: 1})

			fcn, ok := m[c]
			if ok {
				fcn(&d, cf6.Text[1:])
			}
			//}
		}
	}

	return d
}

// Parse008 parses the 008 controlfield for a record and returns a,
// hopefully, human readable translation of the field contents.
func (rec Record) Parse008() (d Cf008Desc) {

	s := rec.GetControlfield("008")

	d = make(Cf008Desc)

	type fn func(md *Cf008Desc, s string)

	m := map[string]fn{
		"BK": parseBibliography008BK,
		"CF": parseBibliography008CF,
		"MP": parseBibliography008MP,
		"MU": parseBibliography008MU,
		"CR": parseBibliography008CR,
		"VM": parseBibliography008VM,
		"MX": parseBibliography008MX,
	}

	switch rec.RecordFormat() {
	case Bibliography:
		parseBibliography008(&d, s)

		c, _ := rec.BibliographyMaterialType()

		// For bibliography 008 fields pass subslice of s -- s[18:]
		// and for bibliography 006 fields pass subslice of s -- s[1:]
		// This way the same parsing functions can parse both 008 and
		// 006 control fields

		fcn, ok := m[c]
		if ok {
			fcn(&d, s[18:])
		}

		cf := rec.GetControlfields("006")
		for _, cf6 := range cf {
			// IIF the form of material matches?
			if pluckByte(cf6.Text, 0) == c || pluckByte(cf6.Text, 0) == "s" {
				fcn, ok := m[c]
				if ok {
					fcn(&d, cf6.Text[1:])
				}
			}
		}

	case Holdings:
		parseHoldings008(&d, s)
	case Authority:
		parseAuthority008(&d, s)
	case Classification:
		parseClassification008(&d, s)
	case Community:
		parseCommunity008(&d, s)
	}

	return d
}

// Parse007 parses the 007 controlfields for a record and returns a,
// hopefully, human readable translation for the field contents.
func (rec Record) Parse007() (d []Cf007Desc) {

	cf := rec.GetControlfields("007")
	rf := rec.RecordFormat()

	// TODO: How much of this could/should be auto-generated?

	type fn func(s string) (d Cf007Desc)

	mb := map[string]fn{
		"a": parseBibliography007MAP,
		"c": parseBibliography007ELR,
		"d": parseBibliography007GLB,
		"f": parseBibliography007TAM,
		"g": parseBibliography007PRG,
		"h": parseBibliography007MIC,
		"k": parseBibliography007NPG,
		"m": parseBibliography007MOP,
		"o": parseBibliography007KIT,
		"q": parseBibliography007NMU,
		"r": parseBibliography007RSI,
		"s": parseBibliography007SOR,
		"t": parseBibliography007TXT,
		"v": parseBibliography007VIR,
		"z": parseBibliography007UNS,
	}

	mh := map[string]fn{
		"a": parseHoldings007MAP,
		"c": parseHoldings007ELR,
		"d": parseHoldings007GLB,
		"f": parseHoldings007TAM,
		"g": parseHoldings007PRG,
		"h": parseHoldings007MIC,
		"k": parseHoldings007NPG,
		"m": parseHoldings007MOP,
		"o": parseHoldings007KIT,
		"q": parseHoldings007NMU,
		"r": parseHoldings007RSI,
		"s": parseHoldings007SOR,
		"t": parseHoldings007TXT,
		"v": parseHoldings007VIR,
		"z": parseHoldings007UNS,
	}

	for _, s := range cf {

		cm := pluckByte(s.Text, 0)

		switch rf {
		case Bibliography:
			fcn, ok := mb[cm]
			if ok {
				d = append(d, fcn(s.Text))
			}

		case Holdings:
			fcn, ok := mh[cm]
			if ok {
				d = append(d, fcn(s.Text))
			}

		case Community:
			d = append(d, parseCommunity007(s.Text))
		}

	}

	return d
}
