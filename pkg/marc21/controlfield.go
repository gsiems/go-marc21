// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"fmt"
	"log"
	"strings"
)

/*
https://www.loc.gov/marc/specifications/specrecstruc.html

    Control fields in MARC 21 formats are assigned tags beginning with
    two zeroes. They are comprised of data and a field terminator; they
    do not contain indicators or subfield codes. The control number
    field is assigned tag 001 and contains the control number of the
    record. Each record contains only one control number field (with
    tag 001), which is to be located at the base address of data.
*/

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Variable control fields - The 00X fields. These fields are
    identified by a field tag in the Directory but they contain neither
    indicator positions nor subfield codes. The variable control fields
    are structurally different from the variable data fields. They may
    contain either a single data element or a series of fixed-length
    data elements identified by relative character position.
*/

//http://www.loc.gov/marc/bibliographic/bd00x.html

// TODO: validate control fields?
// TODO: ensure there are no duplicate control (001, 003, 005, 006, 008) fields?
// TODO: create/update control fields for new/update record?
//          005 -> last updated: yyyymmddhhmmss.f
// TODO: parse/translate 006, 007, 008

type CFValue struct {
	Code  string
	Label string
}

var dateTypePubStatus = map[string]string{
	"b": "No dates given; B.C. date involved",
	"c": "Continuing resource currently published",
	"d": "Continuing resource ceased publication",
	"e": "Detailed date",
	"i": "Inclusive dates of collection",
	"k": "Range of years of bulk of collection",
	"m": "Multiple dates",
	"n": "Dates unknown",
	"p": "Date of distribution/release/issue and production/recording session when different ",
	"q": "Questionable date",
	"r": "Reprint/reissue date and original date",
	"s": "Single known date/probable date",
	"t": "Publication date and copyright date",
	"u": "Continuing resource status unknown",
	"|": "No attempt to code ",
}

var illustrations = map[string]string{
	" ": "No illustrations",
	"a": "Illustrations",
	"b": "Maps",
	"c": "Portraits",
	"d": "Charts",
	"e": "Plans",
	"f": "Plates",
	"g": "Music",
	"h": "Facsimiles",
	"i": "Coats of arms",
	"j": "Genealogical tables",
	"k": "Forms",
	"l": "Samples",
	"m": "Phonodisc, phonowire, etc.",
	"o": "Photographs",
	"p": "Illuminations",
	"|": "No attempt to code",
}

var targetAudience = map[string]string{
	" ": "Unknown or not specified",
	"a": "Preschool",
	"b": "Primary",
	"c": "Pre-adolescent",
	"d": "Adolescent",
	"e": "Adult",
	"f": "Specialized",
	"g": "General",
	"j": "Juvenile",
	"|": "No attempt to code",
}

var itemForm = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"e": "Newspaper format",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}

var frequency = map[string]string{
	" ": "No determinable frequency",
	"a": "Annual",
	"b": "Bimonthly",
	"c": "Semiweekly",
	"d": "Daily",
	"e": "Biweekly",
	"f": "Semiannual",
	"g": "Biennial",
	"h": "Triennial",
	"i": "Three times a week",
	"j": "Three times a month",
	"k": "Continuously updated",
	"m": "Monthly",
	"q": "Quarterly",
	"s": "Semimonthly",
	"t": "Three times a year",
	"u": "Unknown",
	"w": "Weekly",
	"z": "Other",
	"|": "No attempt to code",
}

var regularity = map[string]string{
	"n": "Normalized irregular",
	"r": "Regular",
	"u": "Unknown",
	"x": "Completely irregular",
	"|": "No attempt to code",
}

var continuingResourceType = map[string]string{
	" ": "None of the following",
	"d": "Updating database",
	"l": "Updating loose-leaf",
	"m": "Monographic series",
	"n": "Newspaper",
	"p": "Periodical",
	"w": "Updating Web site",
	"|": "No attempt to code",
}

var originalScriptOfTitle = map[string]string{
	" ": "No alphabet or script given/No key title",
	"a": "Basic Roman",
	"b": "Extended Roman",
	"c": "Cyrillic",
	"d": "Japanese",
	"e": "Chinese",
	"f": "Arabic",
	"g": "Greek",
	"h": "Hebrew",
	"i": "Thai",
	"j": "Devanagari",
	"k": "Korean",
	"l": "Tamil",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var entryConvention = map[string]string{
	"0": "Successive entry",
	"1": "Latest entry",
	"2": "Integrated entry",
	"|": "No attempt to code",
}

var natureOfContents = map[string]string{
	" ": "No specified nature of contents",
	"a": "Abstracts/summaries",
	"b": "Bibliographies",
	"c": "Catalogs",
	"d": "Dictionaries",
	"e": "Encyclopedias",
	"f": "Handbooks",
	"g": "Legal articles",
	"i": "Indexes",
	"j": "Patent document",
	"k": "Discographies",
	"l": "Legislation",
	"m": "Theses",
	"n": "Surveys of literature in a subject area",
	"o": "Reviews",
	"p": "Programmed texts",
	"q": "Filmographies",
	"r": "Directories",
	"s": "Statistics",
	"t": "Technical reports",
	"u": "Standards/specifications",
	"v": "Legal cases and case notes",
	"w": "Law reports and digests",
	"y": "Yearbooks",
	"z": "Treaties",
	"2": "Offprints",
	"5": "Calendars",
	"6": "Comics/graphic novels",
	"|": "No attempt to code",
}

var mapRelief = map[string]string{
	" ": "No relief shown",
	"a": "Contours",
	"b": "Shading",
	"c": "Gradient and bathymetric tints",
	"d": "Hachures",
	"e": "Bathymetry/soundings",
	"f": "Form lines",
	"g": "Spot heights",
	"i": "Pictorially",
	"j": "Land forms",
	"k": "Bathymetry/isolines",
	"m": "Rock drawings",
	"z": "Other",
	"|": "No attempt to code",
}

var mapProjection = map[string]string{
	"##": "Projection not specified",
	"aa": "Aitoff",
	"ab": "Gnomic",
	"ac": "Lambert's azimuthal equal area",
	"ad": "Orthographic",
	"ae": "Azimuthal equidistant",
	"af": "Stereographic",
	"ag": "General vertical near-sided",
	"am": "Modified stereographic for Alaska",
	"an": "Chamberlin trimetric",
	"ap": "Polar stereographic",
	"au": "Azimuthal, specific type unknown",
	"az": "Azimuthal, other",
	"ba": "Gall",
	"bb": "Goode's homolographic",
	"bc": "Lambert's cylindrical equal area",
	"bd": "Mercator",
	"be": "Miller",
	"bf": "Mollweide",
	"bg": "Sinusoidal",
	"bh": "Transverse Mercator",
	"bi": "Gauss-Kruger",
	"bj": "Equirectangular",
	"bk": "Krovak",
	"bl": "Cassini-Soldner",
	"bo": "Oblique Mercator",
	"br": "Robinson",
	"bs": "Space oblique Mercator",
	"bu": "Cylindrical, specific type unknown",
	"bz": "Cylindrical, other",
	"ca": "Albers equal area",
	"cb": "Bonne",
	"cc": "Lambert's conformal conic",
	"ce": "Equidistant conic",
	"cp": "Polyconic",
	"cu": "Conic, specific type unknown",
	"cz": "Conic, other",
	"da": "Armadillo",
	"db": "Butterfly",
	"dc": "Eckert",
	"dd": "Goode's homolosine",
	"de": "Miller's bipolar oblique conformal conic",
	"df": "Van Der Grinten",
	"dg": "Dimaxion",
	"dh": "Cordiform",
	"dl": "Lambert conformal",
	"zz": "Other",
	"||": "No attempt to code",
}

var cartographicMaterialType = map[string]string{
	"a": "Single map",
	"b": "Map series",
	"c": "Map serial",
	"d": "Globe",
	"e": "Atlas",
	"f": "Separate supplement to another work",
	"g": "Bound as part of another work",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var mapSpecialFormat = map[string]string{
	" ": "No specified special format characteristics",
	"e": "Manuscript",
	"j": "Picture card, post card",
	"k": "Calendar",
	"l": "Puzzle",
	"n": "Game",
	"o": "Wall map",
	"p": "Playing cards",
	"r": "Loose-leaf",
	"z": "Other",
	"|": "No attempt to code",
}

var compositionForm = map[string]string{
	"an": "Anthems",
	"bd": "Ballads",
	"bg": "Bluegrass music",
	"bl": "Blues",
	"bt": "Ballets",
	"ca": "Chaconnes",
	"cb": "Chants, Other religions",
	"cc": "Chant, Christian",
	"cg": "Concerti grossi",
	"ch": "Chorales",
	"cl": "Chorale preludes",
	"cn": "Canons and rounds",
	"co": "Concertos",
	"cp": "Chansons, polyphonic",
	"cr": "Carols",
	"cs": "Chance compositions",
	"ct": "Cantatas",
	"cy": "Country music",
	"cz": "Canzonas",
	"df": "Dance forms",
	"dv": "Divertimentos, serenades, cassations, divertissements, and notturni",
	"fg": "Fugues",
	"fl": "Flamenco",
	"fm": "Folk music",
	"ft": "Fantasias",
	"gm": "Gospel music",
	"hy": "Hymns",
	"jz": "Jazz",
	"mc": "Musical revues and comedies",
	"md": "Madrigals",
	"mi": "Minuets",
	"mo": "Motets",
	"mp": "Motion picture music",
	"mr": "Marches",
	"ms": "Masses",
	"mu": "Multiple forms",
	"mz": "Mazurkas",
	"nc": "Nocturnes",
	"nn": "Not applicable",
	"op": "Operas",
	"or": "Oratorios",
	"ov": "Overtures",
	"pg": "Program music",
	"pm": "Passion music",
	"po": "Polonaises",
	"pp": "Popular music",
	"pr": "Preludes",
	"ps": "Passacaglias",
	"pt": "Part-songs",
	"pv": "Pavans",
	"rc": "Rock music",
	"rd": "Rondos",
	"rg": "Ragtime music",
	"ri": "Ricercars",
	"rp": "Rhapsodies",
	"rq": "Requiems",
	"sd": "Square dance music",
	"sg": "Songs",
	"sn": "Sonatas",
	"sp": "Symphonic poems",
	"st": "Studies and exercises",
	"su": "Suites",
	"sy": "Symphonies",
	"tc": "Toccatas",
	"tl": "Teatro lirico",
	"ts": "Trio-sonatas",
	"uu": "Unknown",
	"vi": "Villancicos",
	"vr": "Variations",
	"wz": "Waltzes",
	"za": "Zarzuelas",
	"zz": "Other",
	"||": "No attempt to code",
}

var musicFormat = map[string]string{
	"a": "Full score",
	"b": "Miniature or study score",
	"c": "Accompaniment reduced for keyboard",
	"d": "Voice score with accompaniment omitted",
	"e": "Condensed score or piano-conductor score",
	"g": "Close score",
	"h": "Chorus score",
	"i": "Condensed score",
	"j": "Performer-conductor part",
	"k": "Vocal score",
	"l": "Score",
	"m": "Multiple score formats",
	"n": "Not applicable",
	"p": "Piano score",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var musicParts = map[string]string{
	" ": "No parts in hand or not specified",
	"d": "Instrumental and vocal parts",
	"e": "Instrumental parts",
	"f": "Vocal parts",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

var accompanyingMatter = map[string]string{
	" ": "No accompanying matter",
	"a": "Discography",
	"b": "Bibliography",
	"c": "Thematic index",
	"d": "Libretto or text",
	"e": "Biography of composer or author",
	"f": "Biography of performer or history of ensemble",
	"g": "Technical and/or historical information on instruments",
	"h": "Technical information on music",
	"i": "Historical information",
	"k": "Ethnological information",
	"r": "Instructional materials",
	"s": "Music",
	"z": "Other",
	"|": "No attempt to code",
}

var recordingLiteraryText = map[string]string{
	" ": "Item is a music sound recording",
	"a": "Autobiography",
	"b": "Biography",
	"c": "Conference proceedings",
	"d": "Drama",
	"e": "Essays",
	"f": "Fiction",
	"g": "Reporting",
	"h": "History",
	"i": "Instruction",
	"j": "Language instruction",
	"k": "Comedy",
	"l": "Lectures, speeches",
	"m": "Memoirs",
	"n": "Not applicable",
	"o": "Folktales",
	"p": "Poetry",
	"r": "Rehearsals",
	"s": "Sounds",
	"t": "Interviews",
	"z": "Other",
	"|": "No attempt to code",
}

var transpositionAndArrangement = map[string]string{
	" ": "Not arrangement or transposition or not specified",
	"a": "Transposition",
	"b": "Arrangement",
	"c": "Both transposed and arranged",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

var governmentPublication = map[string]string{
	" ": " Not a government publication",
	"a": " Autonomous or semi-autonomous component",
	"c": " Multilocal",
	"f": " Federal/national",
	"i": " International intergovernmental",
	"l": " Local",
	"m": " Multistate",
	"o": " Government publication-level undetermined",
	"s": " State, provincial, territorial, dependent, etc.",
	"u": " Unknown if item is government publication",
	"z": " Other",
	"|": " No attempt to code",
}

var conferencePublication = map[string]string{
	"0": "Not a conference publication",
	"1": "Conference publication",
	"|": "No attempt to code",
}

var festschrift = map[string]string{
	"0": "Not a festschrift",
	"1": "Festschrift",
	"|": "No attempt to code",
}

var index = map[string]string{
	"0": "No index",
	"1": "Index present",
	"|": "No attempt to code",
}

var literaryForm = map[string]string{
	"0": "Not fiction (not further specified)",
	"1": "Fiction (not further specified)",
	"d": "Dramas",
	"e": "Essays",
	"f": "Novels",
	"h": "Humor, satires, etc.",
	"i": "Letters",
	"j": "Short stories",
	"m": "Mixed forms",
	"p": "Poetry",
	"s": "Speeches",
	"u": "Unknown",
	"|": "No attempt to code",
}

var biography = map[string]string{
	" ": "No biographical material",
	"a": "Autobiography",
	"b": "Individual biography",
	"c": "Collective biography",
	"d": "Contains biographical information",
	"|": "No attempt to code",
}

var visualMaterialType = map[string]string{
	"a": "Art original",
	"b": "Kit",
	"c": "Art reproduction",
	"d": "Diorama",
	"f": "Filmstrip",
	"g": "Game",
	"i": "Picture",
	"k": "Graphic",
	"l": "Technical drawing",
	"m": "Motion picture",
	"n": "Chart",
	"o": "Flash card",
	"p": "Microscope slide",
	"q": "Model",
	"r": "Realia",
	"s": "Slide",
	"t": "Transparency",
	"v": "Videorecording",
	"w": "Toy",
	"z": "Other",
	"|": "No attempt to code",
}

var technique = map[string]string{
	"a": "Animation",
	"c": "Animation and live action",
	"l": "Live action",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var modifiedRecord = map[string]string{
	" ": "Not modified",
	"d": "Dashed-on information omitted",
	"o": "Completely romanized/printed cards romanized",
	"r": "Completely romanized/printed cards in script",
	"s": "Shortened",
	"x": "Missing characters",
	"|": "No attempt to code",
}

var catalogingSource = map[string]string{
	" ": "National bibliographic agency",
	"c": "Cooperative cataloging program",
	"d": "Other",
	"u": "Unknown",
	"|": "No attempt to code",
}

var computerItemForm = map[string]string{
	" ": "Unknown or not specified",
	"o": "Online",
	"q": "Direct electronic",
	"|": "No attempt to code",
}

var computerFileType = map[string]string{
	"a": "Numeric data",
	"b": "Computer program",
	"c": "Representational",
	"d": "Document",
	"e": "Bibliographic data",
	"f": "Font",
	"g": "Game",
	"h": "Sound",
	"i": "Interactive multimedia",
	"j": "Online system or service",
	"m": "Combination",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var materialType = map[string]string{
	"BK": "Books",
	"CF": "Computer Files",
	"MP": "Maps",
	"MU": "Music",
	"CR": "Continuing Resources",
	"VM": "Visual Materials",
	"MX": "Mixed Materials",
}

func lookupCfCode(codeList map[string]string, b []byte, i int) (v CFValue) {

	var code, label string

	if len(b) > i {
		code = string(b[i])
		if code != "" {
			label, _ = codeList[code]
		}
	}

	return CFValue{Code: code, Label: label}
}

// Order matters so:
//      type CfData map[string]map[string]string
// won't work
//      type CfData map[string][]CFValue
// might...

type CfData map[string][]CFValue

type CfCommon struct {
	ControlNumber           string
	ControlNumberIdentifier string
	LatestTransDateTime     string
	FileDate                string
	DateTypePubStatus       CFValue
	Date1                   string
	Date2                   string
	PlaceOfPublication      string
	Language                string
	ModifiedRecord          CFValue
	CatalogingSource        CFValue
	MaterialType            CFValue
}

func (rec Record) parseControlfields() (c CfCommon, md CfData) {

	// 008:
	//  00-05 - Date entered on file
	//  06 - Type of date/Publication status
	//  07-10 - Date 1
	//  11-14 - Date 2
	//  15-17 - Place of publication, production, or execution
	//  18-34 Non-common elements
	//  35-37 - Language
	//  38 - Modified record
	//  39 - Cataloging source

	c.ControlNumber = rec.getCF("001")
	c.ControlNumberIdentifier = rec.getCF("003")
	c.LatestTransDateTime = rec.getCF("005")

	b := []byte(rec.getCF("008"))
	c.FileDate = string(b[0:6])
	c.DateTypePubStatus = lookupCfCode(dateTypePubStatus, b, 6)
	c.Date1 = string(b[7:11])
	c.Date2 = string(b[11:15])
	c.PlaceOfPublication = string(b[15:18])

	c.Language = string(b[35:38])
	c.ModifiedRecord = lookupCfCode(modifiedRecord, b, 38)
	c.CatalogingSource = lookupCfCode(catalogingSource, b, 39)

	code, label := rec.MaterialType()

	c.MaterialType = CFValue{Code: code, Label: label}

	md = make(CfData)

	switch code {
	case "BK":
		md.parseBookCf(b[18:35])
	case "CF":
		md.parseComputerFilesCf(b[18:35])
	case "MP":
		md.parseMapCf(b[18:35])
	case "MU":
		md.parseMusicCf(b[18:35])
	case "CR":
		md.parseContinuingResourcesCf(b[18:35])
	case "VM":
		md.parseVisualMaterialsCf(b[18:35])
	case "MX":
		md.parseMixedMaterialsCf(b[18:35])
	}

	cf006 := rec.getCFs("006")
	for _, x := range cf006 {
		b := []byte(x)

		// NOTE: The Material Form is b[0] and ignoring it *could*
		// disconnect the material form from the corresponding values.
		// However the material form *should* match the leader record type
		// so that really shouldn't be an issue... should it?
		// Perhaps a check to verify that b[0] matches?

		switch code {
		case "BK":
			md.parseBookCf(b[1:])
		case "CF":
			md.parseComputerFilesCf(b[1:])
		case "MP":
			md.parseMapCf(b[1:])
		case "MU":
			md.parseMusicCf(b[1:])
		case "CR":
			md.parseContinuingResourcesCf(b[1:])
		case "VM":
			md.parseVisualMaterialsCf(b[1:])
		case "MX":
			md.parseMixedMaterialsCf(b[1:])
		}
	}

	return c, md
}

// parseBookCf parses the control field data for Book (BK) material types.
func (md CfData) parseBookCf(b []byte) {

	// 008 (NR):
	//  18-21 - Illustrations (006/01-04)
	//  22 - Target audience (006/05)
	//  23 - Form of item (006/06)
	//  24-27 - Nature of contents (006/07-10)
	//  28 - Government publication (006/11)
	//  29 - Conference publication (006/12)
	//  30 - Festschrift (006/13)
	//  31 - Index (006/14)
	//  32 - Undefined
	//  33 - Literary form (006/16)
	//  34 - Biography (006/17)

	// 006 (R):
	//  00 - Form of material
	//  01-04 - Illustrations
	//  05 - Target audience
	//  06 - Form of item
	//  07-10 - Nature of contents
	//  11 - Government publication
	//  12 - Conference publication
	//  13 - Festschrift
	//  14 - Index
	//  15 - Undefined
	//  16 - Literary form
	//  17 - Biography

	for i := 0; i <= 3; i++ {
		md.appendCfData("Illustrations", lookupCfCode(illustrations, b, i))
	}

	md.appendCfData("TargetAudience", lookupCfCode(targetAudience, b, 4))
	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 5))

	for i := 6; i <= 9; i++ {
		md.appendCfData("NatureOfContents", lookupCfCode(natureOfContents, b, i))
	}

	md.appendCfData("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	md.appendCfData("ConferencePublication", lookupCfCode(conferencePublication, b, 11))
	md.appendCfData("Festschrift", lookupCfCode(festschrift, b, 12))
	md.appendCfData("Index", lookupCfCode(index, b, 13))
	md.appendCfData("LiteraryForm", lookupCfCode(literaryForm, b, 15))
	md.appendCfData("Biography", lookupCfCode(biography, b, 16))
}

// parseComputerFilesCf parses the control field data for ComputerFiles (MU) material types.
func (md CfData) parseComputerFilesCf(b []byte) {

	// 008:
	//  18-21 - Undefined
	//  22 - Target audience (006/05)
	//  23 - Form of item (006/06)
	//  24-25 - Undefined
	//  26 - Type of computer file (006/09)
	//  27 - Undefined
	//  28 - Government publication (006/11)
	//  29-34 - Undefined

	md.appendCfData("TargetAudience", lookupCfCode(targetAudience, b, 4))
	md.appendCfData("ItemForm", lookupCfCode(computerItemForm, b, 5))
	md.appendCfData("ComputerFileType", lookupCfCode(computerFileType, b, 8))
	md.appendCfData("GovernmentPublication", lookupCfCode(governmentPublication, b, 9))
}

// parseMapCf parses the control field data for Map (MP) material types.
func (md CfData) parseMapCf(b []byte) {

	// 008:
	//  18-21 - Relief (006/01-04)
	//  22-23 - Projection (006/05-06)
	//  24 - Undefined
	//  25 - Type of cartographic material (006/08)
	//  26-27 - Undefined
	//  28 - Government publication (006/11)
	//  29 - Form of item (006/12)
	//  30 - Undefined
	//  31 - Index (006/14)
	//  23 - Undefined
	//  33-34 - Special format characteristics (006/16-17)

	for i := 0; i <= 3; i++ {
		md.appendCfData("Relief", lookupCfCode(mapRelief, b, i))
	}

	for i := 4; i <= 5; i++ {
		md.appendCfData("Projection", lookupCfCode(mapProjection, b, i))
	}

	md.appendCfData("CartographicMaterialType", lookupCfCode(cartographicMaterialType, b, 7))
	md.appendCfData("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 11))
	md.appendCfData("Index", lookupCfCode(index, b, 13))

	for i := 15; i <= 16; i++ {
		md.appendCfData("SpecialFormat", lookupCfCode(mapSpecialFormat, b, i))
	}
}

// parseMusicCf parses the control field data for Music (MU) material types.
func (md CfData) parseMusicCf(b []byte) {

	// 008:
	//  18-19 - Form of composition (006/01-02)
	//  20 - Format of music (006/03)
	//  21 - Music parts (006/04)
	//  22 - Target audience (006/05)
	//  23 - Form of item (006/06)
	//  24-29 - Accompanying matter (006/07-12)
	//  30-31 - Literary text for sound recordings (006/13-14)
	//  32 - Undefined
	//  33 - Transposition and arrangement (006/16)
	//  34 - Undefined

	for i := 0; i <= 1; i++ {
		md.appendCfData("CompositionForm", lookupCfCode(compositionForm, b, i))
	}

	md.appendCfData("MusicFormat", lookupCfCode(musicFormat, b, 2))
	md.appendCfData("MusicParts", lookupCfCode(musicParts, b, 3))
	md.appendCfData("TargetAudience", lookupCfCode(targetAudience, b, 4))
	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 5))

	for i := 6; i <= 11; i++ {
		md.appendCfData("AccompanyingMatter", lookupCfCode(accompanyingMatter, b, i))
	}

	for i := 12; i <= 13; i++ {
		md.appendCfData("RecordingLiteraryText", lookupCfCode(recordingLiteraryText, b, i))
	}

	md.appendCfData("TranspositionAndArrangement", lookupCfCode(transpositionAndArrangement, b, 15))
}

// parseContinuingResourcesCf parses the control field data for ContinuingResources (MU) material types.
func (md CfData) parseContinuingResourcesCf(b []byte) {

	// 008:
	//  18 - Frequency (006/01)
	//  19 - Regularity (006/02)
	//  20 - Undefined
	//  21 - Type of continuing resource (006/04)
	//  22 - Form of original item (006/05)
	//  23 - Form of item (006/06)
	//  24 - Nature of entire work (006/07)
	//  25-27 - Nature of contents (006/08-10)
	//  28 - Government publication (006/11)
	//  29 - Conference publication (006/12)
	//  30-32 - Undefined
	//  33 - Original alphabet or script of title (006/16)
	//  34 - Entry convention (006/17)

	md.appendCfData("Frequency", lookupCfCode(frequency, b, 0))
	md.appendCfData("Regularity", lookupCfCode(regularity, b, 1))
	md.appendCfData("ContinuingResourceType", lookupCfCode(continuingResourceType, b, 3))
	md.appendCfData("OriginalItemForm", lookupCfCode(itemForm, b, 4))
	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 5))
	md.appendCfData("NatureOfWork", lookupCfCode(natureOfContents, b, 6))

	for i := 7; i <= 9; i++ {
		md.appendCfData("NatureOfContents", lookupCfCode(natureOfContents, b, i))
	}
	md.appendCfData("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	md.appendCfData("ConferencePublication", lookupCfCode(conferencePublication, b, 11))
	md.appendCfData("OriginalScriptOfTitle", lookupCfCode(originalScriptOfTitle, b, 15))
	md.appendCfData("EntryConvention", lookupCfCode(entryConvention, b, 16))
}

// parseVisualMaterialsCf parses the control field data for VisualMaterials (MU) material types.
func (md CfData) parseVisualMaterialsCf(b []byte) {

	// 008:
	//  18-20 - Running time for motion pictures and videorecordings (006/01-03)
	//      000 - Running time exceeds three characters
	//      001-999 - Running time
	//      nnn - Not applicable
	//      --- - Unknown
	//      ||| - No attempt to code
	//  21 - Undefined
	//  22 - Target audience (006/05)
	//  23-27 - Undefined
	//  28 - Government publication (006/11)
	//  29 - Form of item (006/12)
	//  30-32 - Undefined
	//  33 - Type of visual material (006/16)
	//  34 - Technique (006/17)

	rt := string(b[0:3])
	var rl string
	switch rt {
	case "000":
		rl = "Running time exceeds three characters"
	case "nnn":
		rl = "Not applicable"
	case "---":
		rl = "Unknown"
	case "|||":
		rl = "No attempt to code"
	default:
		rl = "Running time"
	}
	md.appendCfData("RunningTime", CFValue{Code: rt, Label: rl})

	md.appendCfData("TargetAudience", lookupCfCode(targetAudience, b, 4))
	md.appendCfData("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 11))
	md.appendCfData("VisualMaterialType", lookupCfCode(visualMaterialType, b, 15))
	md.appendCfData("Technique", lookupCfCode(technique, b, 16))
}

// parseMixedMaterialsCf parses the control field data for MixedMaterials (MU) material types.
func (md CfData) parseMixedMaterialsCf(b []byte) {

	// 008:
	//  18-22 - Undefined
	//  23 - Form of item (006/06)
	//  24-34 - Undefined

	md.appendCfData("ItemForm", lookupCfCode(itemForm, b, 5))
}

func (md CfData) appendCfData(k string, c CFValue) {

	// Ensure that the code/label hasn't already been appended (no duplicate code/labels)
	exists := false
	for _, x := range md[k] {
		if x.Code == c.Code {
			exists = true
			break
		}
	}

	if !exists {
		md[k] = append(md[k], c)
	}
}

// getCF returns the control field for the specified non-repeating tag
func (rec Record) getCF(tag string) string {
	for _, cf := range rec.Controlfields {
		if cf.Tag == tag {
			return cf.Text
		}
	}
	return ""
}

// getCFs returns the unique control field(s) for the specified repeating tag
func (rec Record) getCFs(tag string) (t []string) {

	uniq := make(map[string]bool)

	for _, cf := range rec.Controlfields {
		if cf.Tag == tag {
			uniq[cf.Text] = true
		}
	}

	for k := range uniq {
		t = append(t, k)
	}

	return t
}

// MaterialType returns the code and description of the type of material
// documented by the record. {"Books", "Computer Files", "Maps", "Music",
// "Continuing Resources", "Visual Materials" or "Mixed Materials"}
func (rec Record) MaterialType() (code, label string) {
	rt, _ := rec.RecordType()
	bl, _ := rec.BibliographicLevel()

	switch rt {
	case "c", "d", "i", "j":
		return "MU", "Music"
	case "e", "f":
		return "MP", "Maps"
	case "g", "k", "o", "r":
		return "VM", "Visual Materials"
	case "m":
		return "CF", "Computer Files"
	case "p":
		return "MX", "Mixed Materials"
	case "t":
		switch bl {
		case "a", "c", "d", "m":
			return "BK", "Books"
		}
	case "a":
		switch bl {
		case "a", "c", "d", "m":
			return "BK", "Books"
		case "b", "i", "s":
			return "CR", "Continuing Resources"
		}
	}
	return code, label
}

// extractControlfields extracts the control fields from the raw MARC record bytes
func extractControlfields(rawRec []byte, baseAddress int, dir []*Directory) (cfs []*Controlfield, err error) {

	// There are records where the 003 and 007 fields are dorky (this
	// may happen to other fields also??) where the first byte is a
	// terminator character and the directory indicates that the field
	// is longer.
	//
	// The directory for one record that has this issue looks like:
	//      001 0 11
	//      003 11 12
	//      005 12 17
	//      007 29 18
	//      008 30 41
	// where we can see that the directory has the 003 and 005 tags
	// overlapping and the 007 and 008 tags overlapping with no actual
	// data for either dorked-up 003/007 tag. Since the remainder of the
	// record appears to be good we don't want to fail, but we do want
	// to bring attention to the data issue.
	parseError := false

	for _, d := range dir {
		if strings.HasPrefix(d.Tag, "00") {

			start := baseAddress + d.StartingPos
			b := rawRec[start : start+d.FieldLength]

			if b[len(b)-1] == FieldTerminator {
				cfs = append(cfs, &Controlfield{Tag: d.Tag, Text: string(b[:len(b)-1])})
			} else {
				parseError = true
			}
		}
	}

	if parseError {
		log.Printf("Control fields extraction error: %s\n", cfs)
	}

	return cfs, nil
}

// Implement the Stringer interface for "Pretty-printing"
func (cf Controlfield) String() string {
	return fmt.Sprintf("{%s: '%s'}", cf.Tag, cf.Text)
}
