// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

// http://www.loc.gov/marc/bibliographic/bd006.html
// http://www.loc.gov/marc/bibliographic/bd008.html

type CfMatlDesc map[string][]CfValue

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
	"  ": "Projection not specified",
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

var runningTime = map[string]string{
	"000": "Running time exceeds three characters",
	//001-999 - Running time
	"nnn": "Not applicable",
	"---": "Unknown",
	"|||": "No attempt to code",
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

var materialCategory = map[string]string{
	"a": "Map",
	"c": "Electronic resource",
	"d": "Globe",
	"f": "Tactile material",
	"g": "Projected graphic",
	"h": "Microform",
	"k": "Nonprojected graphic",
	"m": "Motion picture",
	"o": "Kit",
	"q": "Notated music",
	"r": "Remote-sensing image",
	"s": "Sound recording",
	"t": "Text",
	"v": "Videorecording",
	"z": "Unspecified",
}

// parseBookCf parses the control field data for Book (BK) material types.
func (mc CfMatlDesc) parseBookCf(b []byte) {

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

	for i := 0; i <= 3; i++ {
		mc.appendCfMatlDesc("Illustrations", cfShortCode(illustrations, b, i))
	}

	mc.appendCfMatlDesc("TargetAudience", cfShortCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 5))

	for i := 6; i <= 9; i++ {
		mc.appendCfMatlDesc("NatureOfContents", cfShortCode(natureOfContents, b, i))
	}

	mc.appendCfMatlDesc("GovernmentPublication", cfShortCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ConferencePublication", cfShortCode(conferencePublication, b, 11))
	mc.appendCfMatlDesc("Festschrift", cfShortCode(festschrift, b, 12))
	mc.appendCfMatlDesc("Index", cfShortCode(index, b, 13))
	mc.appendCfMatlDesc("LiteraryForm", cfShortCode(literaryForm, b, 15))
	mc.appendCfMatlDesc("Biography", cfShortCode(biography, b, 16))
}

// parseComputerFilesCf parses the control field data for ComputerFiles (MU) material types.
func (mc CfMatlDesc) parseComputerFilesCf(b []byte) {

	// 008:
	//  18-21 - Undefined
	//  22 - Target audience (006/05)
	//  23 - Form of item (006/06)
	//  24-25 - Undefined
	//  26 - Type of computer file (006/09)
	//  27 - Undefined
	//  28 - Government publication (006/11)
	//  29-34 - Undefined

	mc.appendCfMatlDesc("TargetAudience", cfShortCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(computerItemForm, b, 5))
	mc.appendCfMatlDesc("ComputerFileType", cfShortCode(computerFileType, b, 8))
	mc.appendCfMatlDesc("GovernmentPublication", cfShortCode(governmentPublication, b, 10))
}

// parseMapCf parses the control field data for Map (MP) material types.
func (mc CfMatlDesc) parseMapCf(b []byte) {

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
		mc.appendCfMatlDesc("Relief", cfShortCode(mapRelief, b, i))
	}

	mc.appendCfMatlDesc("Projection", cfWideCode(mapProjection, b, 4, 2))
	mc.appendCfMatlDesc("CartographicMaterialType", cfShortCode(cartographicMaterialType, b, 7))
	mc.appendCfMatlDesc("GovernmentPublication", cfShortCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 11))
	mc.appendCfMatlDesc("Index", cfShortCode(index, b, 13))

	for i := 15; i <= 16; i++ {
		mc.appendCfMatlDesc("SpecialFormat", cfShortCode(mapSpecialFormat, b, i))
	}
}

// parseMusicCf parses the control field data for Music (MU) material types.
func (mc CfMatlDesc) parseMusicCf(b []byte) {

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

	mc.appendCfMatlDesc("CompositionForm", cfWideCode(compositionForm, b, 1, 2))
	mc.appendCfMatlDesc("MusicFormat", cfShortCode(musicFormat, b, 2))
	mc.appendCfMatlDesc("MusicParts", cfShortCode(musicParts, b, 3))
	mc.appendCfMatlDesc("TargetAudience", cfShortCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 5))

	for i := 6; i <= 11; i++ {
		mc.appendCfMatlDesc("AccompanyingMatter", cfShortCode(accompanyingMatter, b, i))
	}

	for i := 12; i <= 13; i++ {
		mc.appendCfMatlDesc("RecordingLiteraryText", cfShortCode(recordingLiteraryText, b, i))
	}

	mc.appendCfMatlDesc("TranspositionAndArrangement", cfShortCode(transpositionAndArrangement, b, 15))
}

// parseContinuingResourcesCf parses the control field data for ContinuingResources (MU) material types.
func (mc CfMatlDesc) parseContinuingResourcesCf(b []byte) {

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

	mc.appendCfMatlDesc("Frequency", cfShortCode(frequency, b, 0))
	mc.appendCfMatlDesc("Regularity", cfShortCode(regularity, b, 1))
	mc.appendCfMatlDesc("ContinuingResourceType", cfShortCode(continuingResourceType, b, 3))
	mc.appendCfMatlDesc("OriginalItemForm", cfShortCode(itemForm, b, 4))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 5))
	mc.appendCfMatlDesc("NatureOfWork", cfShortCode(natureOfContents, b, 6))

	for i := 7; i <= 9; i++ {
		mc.appendCfMatlDesc("NatureOfContents", cfShortCode(natureOfContents, b, i))
	}
	mc.appendCfMatlDesc("GovernmentPublication", cfShortCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ConferencePublication", cfShortCode(conferencePublication, b, 11))
	mc.appendCfMatlDesc("OriginalScriptOfTitle", cfShortCode(originalScriptOfTitle, b, 15))
	mc.appendCfMatlDesc("EntryConvention", cfShortCode(entryConvention, b, 16))
}

// parseVisualMaterialsCf parses the control field data for VisualMaterials (MU) material types.
func (mc CfMatlDesc) parseVisualMaterialsCf(b []byte) {

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

	rt := cfWideCode(runningTime, b, 0, 3)
	if rt.Code != "" && rt.Label == "" {
		rt.Label = "Running time"
	}
	mc.appendCfMatlDesc("RunningTime", rt)

	mc.appendCfMatlDesc("TargetAudience", cfShortCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("GovernmentPublication", cfShortCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 11))
	mc.appendCfMatlDesc("VisualMaterialType", cfShortCode(visualMaterialType, b, 15))
	mc.appendCfMatlDesc("Technique", cfShortCode(technique, b, 16))
}

// parseMixedMaterialsCf parses the control field data for MixedMaterials (MU) material types.
func (mc CfMatlDesc) parseMixedMaterialsCf(b []byte) {

	// 008:
	//  18-22 - Undefined
	//  23 - Form of item (006/06)
	//  24-34 - Undefined

	mc.appendCfMatlDesc("ItemForm", cfShortCode(itemForm, b, 5))
}

func (mc CfMatlDesc) appendCfMatlDesc(k string, c CfValue) {

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
