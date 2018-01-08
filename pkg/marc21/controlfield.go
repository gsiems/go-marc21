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

type CfValue struct {
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

var materialType = map[string]string{
	"BK": "Books",
	"CF": "Computer Files",
	"MP": "Maps",
	"MU": "Music",
	"CR": "Continuing Resources",
	"VM": "Visual Materials",
	"MX": "Mixed Materials",
}

// 01 - Specific material designation
var pdSpecificDesignationA = map[string]string{
	"d": "Atlas",
	"g": "Diagram",
	"j": "Map",
	"k": "Profile",
	"q": "Model",
	"r": "Remote-sensing image",
	"s": "Section",
	"u": "Unspecified",
	"y": "View",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorA = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}

// 04 - Physical medium
var pdPhysicalMediumA = map[string]string{
	"a": "Paper",
	"b": "Wood",
	"c": "Stone",
	"d": "Metal",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textiles",
	"i": "Plastic",
	"j": "Glass",
	"l": "Vinyl",
	"n": "Vellum 	p - Plaster",
	"q": "Flexible base photographic, positive",
	"r": "Flexible base photographic, negative",
	"s": "Non-flexible base photographic, positive",
	"t": "Non-flexible base photographic, negative",
	"u": "Unknown",
	"v": "Leather",
	"w": "Parchment",
	"y": "Other photographic medium",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Type of reproduction
var pdReproductionTypeA = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 06 - Production/reproduction details
var pdProductionDetailsA = map[string]string{
	"a": "Photocopy, blueline print",
	"b": "Photocopy",
	"c": "Pre-production",
	"d": "Film",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Positive/negative aspect
var pdPositiveNegativeAspectA = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"n": "Not applicable",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationC = map[string]string{
	"a": "Tape cartridge",
	"b": "Chip cartridge",
	"c": "Computer optical disc cartridge",
	"d": "Computer disc, type unspecified",
	"e": "Computer disc cartridge, type unspecified",
	"f": "Tape cassette",
	"h": "Tape reel",
	"j": "Magnetic disk",
	"k": "Computer card",
	"m": "Magneto-optical disc",
	"o": "Optical disc",
	"r": "Remote",
	"s": "Standalone device",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorC = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"g": "Gray scale",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Dimensions
var pdDimensionsC = map[string]string{
	"a": "3 1/2 in.",
	"e": "12 in.",
	"g": "4 3/4 in. or 12 cm.",
	"i": "1 1/8 x 2 3/8 in.",
	"j": "3 7/8 x 2 1/2 in.",
	"n": "Not applicable",
	"o": "5 1/4 in.",
	"u": "Unknown",
	"v": "8 in.",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Sound
var pdSoundC = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 09 - File formats
var pdFileFormatsC = map[string]string{
	"a": "One file format",
	"m": "Multiple file formats",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 10 - Quality assurance target(s)
var pdQATargetsC = map[string]string{
	"a": "Absent",
	"n": "Not applicable",
	"p": "Present",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 11 - Antecedent/Source
var pdSourceC = map[string]string{
	"a": "File reproduced from original",
	"b": "File reproduced from microform",
	"c": "File reproduced from an electronic resource",
	"d": "File reproduced from an intermediate (not microform)",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 12 - Level of compression
var pdCompressionC = map[string]string{
	"a": "Uncompressed",
	"b": "Lossless",
	"d": "Lossy",
	"m": "Mixed",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 13 - Reformatting Quality
var pdReformattingQualityC = map[string]string{
	"a": "Access",
	"n": "Not applicable",
	"p": "Preservation",
	"r": "Replacement",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationD = map[string]string{
	"a": "Celestial globe",
	"b": "Planetary or lunar globe",
	"c": "Terrestrial globe",
	"e": "Earth moon globe",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorD = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}

// 04 - Physical medium
var pdPhysicalMediumD = map[string]string{
	"a": "Paper",
	"b": "Wood",
	"c": "Stone",
	"d": "Metal",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textile",
	"i": "Plastic",
	"l": "Vinyl",
	"n": "Vellum",
	"p": "Plaster",
	"u": "Unknown",
	"v": "Leather",
	"w": "Parchment",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Type of reproduction
var pdReproductionTypeD = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationF = map[string]string{
	"a": "Moon",
	"b": "Braille",
	"c": "Combination",
	"d": "Tactile, with no writing system",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03-04 - Class of braille writing
var pdBrailleWritingClassF = map[string]string{
	" ": "No specified class of braille writing",
	"a": "Literary braille",
	"b": "Format code braille",
	"c": "Mathematics and scientific braille",
	"d": "Computer braille",
	"e": "Music braille",
	"m": "Multiple braille types",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Level of contraction
var pdContractionLevelF = map[string]string{
	"a": "Uncontracted",
	"b": "Contracted",
	"m": "Combination",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 06-08 - Braille music format
var pdBrailleMusicFormatF = map[string]string{
	" ":   "No specified braille music format",
	"a":   "Bar over bar",
	"b":   "Bar by bar",
	"c":   "Line over line",
	"d":   "Paragraph",
	"e":   "Single line",
	"f":   "Section by section",
	"g":   "Line by line",
	"h":   "Open score",
	"i":   "Spanner short form scoring",
	"j":   "Short form scoring",
	"k":   "Outline",
	"l":   "Vertical score",
	"n":   "Not applicable",
	"u":   "Unknown",
	"z":   "Other",
	"|||": "No attempt to code",
}

// 09 - Special physical characteristics
var pdSpecialCharacteristicsF = map[string]string{
	"a": "Print/braille",
	"b": "Jumbo or enlarged braille",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationG = map[string]string{
	"c": "Filmstrip cartridge",
	"d": "Filmslip",
	"f": "Filmstrip, type unspecified",
	"o": "Filmstrip roll",
	"s": "Slide",
	"t": "Transparency",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorG = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Base of emulsion
var pdEmulsionBaseG = map[string]string{
	"d": "Glass",
	"e": "Synthetic",
	"j": "Safety film",
	"k": "Film base, other than safety film",
	"m": "Mixed collection",
	"o": "Paper",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Sound on medium or separate
var pdSoundLocationG = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 06 - Medium for sound
var pdSoundMediumG = map[string]string{
	" ": "No sound (silent)",
	"a": "Optical sound track on motion picture film",
	"b": "Magnetic sound track on motion picture film",
	"c": "Magnetic audio tape in cartridge",
	"d": "Sound disc",
	"e": "Magnetic audio tape on reel",
	"f": "Magnetic audio tape in cassette",
	"g": "Optical and magnetic sound track on motion picture film",
	"h": "Videotape",
	"i": "Videodisc",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Dimensions
var pdDimensionsG = map[string]string{
	"a": "Standard 8 mm. film width",
	"b": "Super 8 mm./single 8 mm. film width",
	"c": "9.5 mm. film width",
	"d": "16 mm. film width",
	"e": "28 mm. film width",
	"f": "35 mm. film width",
	"g": "70 mm. film width",
	"j": "2x2 in. or 5x5 cm. slide",
	"k": "2 1/4 x 2 1/4 in. or 6x6 cm. slide",
	"s": "4x5 in. or 10x13 cm. transparency",
	"t": "5x7 in. or 13x18 cm. transparency",
	"v": "8x10 in. or 21x26 cm. transparency",
	"w": "9x9 in. or 23x23 cm. transparency",
	"x": "10x10 in. or 26x26 cm. transparency",
	"y": "7x7 in. or 18x18 cm. transparency",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 08 - Secondary support material
var pdSecondarySupportMaterialG = map[string]string{
	" ": "No secondary support",
	"c": "Cardboard",
	"d": "Glass",
	"e": "Synthetic",
	"h": "Metal",
	"j": "Metal and glass",
	"k": "Synthetic and glass",
	"m": "Mixed collection",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationH = map[string]string{
	"a": "Aperture card",
	"b": "Microfilm cartridge",
	"c": "Microfilm cassette",
	"d": "Microfilm reel",
	"e": "Microfiche",
	"f": "Microfiche cassette",
	"g": "Microopaque",
	"h": "Microfilm slip",
	"j": "Microfilm roll",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Positive/negative aspect
var pdPositiveNegativeAspectH = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 04 - Dimensions
var pdDimensionsH = map[string]string{
	// Microfilm
	"a": "8 mm.",
	"d": "16 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"h": "105 mm.",
	// Microfiche, Microopaque, etc.
	"l": "3x5 in. or 8x13 cm.",
	"m": "4x6 in. or 11x15 cm.",
	"o": "6x9 in. or 16x23 cm.",
	// Aperture Card
	"p": "3 1/4 x 7 3/8 in. or 9x19 cm.",
	// Other
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Reduction ratio range
var pdReductionRatioRangeH = map[string]string{
	"a": "Low reduction ratio",
	"b": "Normal reduction",
	"c": "High reduction",
	"d": "Very high reduction",
	"e": "Ultra high reduction",
	"u": "Unknown",
	"v": "Reduction rate varies",
	"|": "No attempt to code",
}

// 09 - Color
var pdColorH = map[string]string{
	"b": "Black-and-white",
	"c": "Multicolored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 10 - Emulsion on film
var pdFilmEmulsionH = map[string]string{
	"a": "Silver halide",
	"b": "Diazo",
	"c": "Vesicular",
	"m": "Mixed emulsion",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 11 - Generation
var pdGenerationH = map[string]string{
	"a": "First generation (master)",
	"b": "Printing master",
	"c": "Service copy",
	"m": "Mixed generation",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 12 - Base of film
var pdFilmBaseH = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationK = map[string]string{
	"a": "Activity card",
	"c": "Collage",
	"d": "Drawing",
	"e": "Painting",
	"f": "Photomechanical print",
	"g": "Photonegative",
	"h": "Photoprint",
	"i": "Picture",
	"j": "Print",
	"k": "Poster",
	"l": "Technical drawing",
	"n": "Chart",
	"o": "Flash card",
	"p": "Postcard",
	"q": "Icon",
	"r": "Radiograph",
	"s": "Study print",
	"u": "Unspecified",
	"v": "Photograph, type unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorK = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Primary support material
var pdPrimarySupportMaterialK = map[string]string{
	"a": "Canvas",
	"b": "Bristol board",
	"c": "Cardboard/illustration board",
	"d": "Glass",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textile",
	"h": "Metal",
	"i": "Plastic",
	"l": "Vinyl",
	"m": "Mixed collection",
	"n": "Vellum",
	"o": "Paper",
	"p": "Plaster",
	"q": "Hardboard",
	"r": "Porcelain",
	"s": "Stone",
	"t": "Wood",
	"u": "Unknown",
	"v": "Leather",
	"w": "Parchment",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Secondary support material
var pdSecondarySupportMaterialK = map[string]string{
	" ": "No secondary support",
	"a": "Canvas",
	"b": "Bristol board",
	"c": "Cardboard/illustration board",
	"d": "Glass",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textile",
	"h": "Metal",
	"i": "Plastic",
	"l": "Vinyl",
	"m": "Mixed collection",
	"n": "Vellum",
	"o": "Paper",
	"p": "Plaster",
	"q": "Hardboard",
	"r": "Porcelain",
	"s": "Stone",
	"t": "Wood",
	"u": "Unknown",
	"v": "Leather",
	"w": "Parchment",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationM = map[string]string{
	"c": "Film cartridge",
	"f": "Film cassette",
	"o": "Film roll",
	"r": "Film reel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorM = map[string]string{
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Motion picture presentation format
var pdPresentationFormatM = map[string]string{
	"a": "Standard sound aperture (reduced frame)",
	"b": "Nonanamorphic (wide-screen)",
	"c": "3D",
	"d": "Anamorphic (wide-screen)",
	"e": "Other wide-screen format",
	"f": "Standard silent aperture (full frame)",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Sound on medium or separate
var pdSoundLocationM = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 06 - Medium for sound
var pdSoundMediumM = map[string]string{
	" ": "No sound (silent)",
	"a": "Optical sound track on motion picture film",
	"b": "Magnetic sound track on motion picture film",
	"c": "Magnetic audio tape in cartridge",
	"d": "Sound disc",
	"e": "Magnetic audio tape on reel",
	"f": "Magnetic audio tape in cassette",
	"g": "Optical and magnetic sound track on motion picture film",
	"h": "Videotape",
	"i": "Videodisc",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Dimensions
var pdDimensionsM = map[string]string{
	"a": "Standard 8 mm.",
	"b": "Super 8 mm./single 8 mm.",
	"c": "9.5 mm.",
	"d": "16 mm.",
	"e": "28 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 08 - Configuration of playback channels
var pdPlaybackChannelConfigM = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 09 - Production elements
var pdProductionElementsM = map[string]string{
	"a": "Workprint",
	"b": "Trims",
	"c": "Outtakes",
	"d": "Rushes",
	"e": "Mixing tracks",
	"f": "Title bands/inter-title rolls",
	"g": "Production rolls",
	"n": "Not applicable",
	"z": "Other",
	"|": "No attempt to code",
}

// 10 - Positive/negative aspect
var pdPositiveNegativeAspectM = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 11 - Generation
var pdGenerationM = map[string]string{
	"d": "Duplicate",
	"e": "Master",
	"o": "Original",
	"r": "Reference print/viewing copy",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 12 - Base of film
var pdFilmBaseM = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 13 - Refined categories of color
var pdColorCategoriesM = map[string]string{
	"a": "3 layer color",
	"b": "2 color, single strip",
	"c": "Undetermined 2 color",
	"d": "Undetermined 3 color",
	"e": "3 strip color",
	"f": "2 strip color",
	"g": "Red strip",
	"h": "Blue or green strip",
	"i": "Cyan strip",
	"j": "Magenta strip",
	"k": "Yellow strip",
	"l": "S E N 2",
	"m": "S E N 3",
	"n": "Not applicable",
	"p": "Sepia tone",
	"q": "Other tone",
	"r": "Tint",
	"s": "Tinted and toned",
	"t": "Stencil color",
	"u": "Unknown",
	"v": "Hand colored",
	"z": "Other",
	"|": "No attempt to code",
}

// 14 - Kind of color stock or print
var pdColorTypeM = map[string]string{
	"a": "Imbibition dye transfer prints",
	"b": "Three-layer stock",
	"c": "Three layer stock, low fade",
	"d": "Duplitized stock",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 15 - Deterioration stage
var pdDeteriorationStageM = map[string]string{
	// Other
	"a": "None apparent",
	"b": "Nitrate: suspicious odor",
	"c": "Nitrate: pungent odor",
	"d": "Nitrate: brownish, discoloration, fading, dusty",
	"e": "Nitrate: sticky",
	"f": "Nitrate: frothy, bubbles, blisters",
	"g": "Nitrate: congealed",
	"h": "Nitrate: powder",
	"k": "Non-nitrate: detectable deterioration",
	"l": "Non-nitrate: advanced deterioration",
	"m": "Non-nitrate: disaster",
	"|": "No attempt to code",
}

// 16 - Completeness
var pdCompletenessM = map[string]string{
	"c": "Complete",
	"i": "Incomplete",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

//17-22 - Film inspection date

// 01 - Specific material designation
var pdSpecificDesignationO = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationQ = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationR = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// 03 - Altitude of sensor
var pdSensorAltitudeR = map[string]string{
	"a": "Surface",
	"b": "Airborne",
	"c": "Spaceborne",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Attitude of sensor
var pdSensorAttitudeR = map[string]string{
	"a": "Low oblique",
	"b": "High oblique",
	"c": "Vertical",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 05 - Cloud cover
var pdCloudCoverR = map[string]string{
	"0": "0-9%",
	"1": "10-19%",
	"2": "20-29%",
	"3": "30-39%",
	"4": "40-49%",
	"5": "50-59%",
	"6": "60-69%",
	"7": "70-79%",
	"8": "80-89%",
	"9": "90-100%",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 06 - Platform construction type
var pdPlatformConstructionTypeR = map[string]string{
	"a": "Balloon",
	"b": "Aircraft--low altitude",
	"c": "Aircraft--medium altitude",
	"d": "Aircraft--high altitude",
	"e": "Manned spacecraft",
	"f": "Unmanned spacecraft",
	"g": "Land-based remote-sensing device",
	"h": "Water surface-based remote-sensing device",
	"i": "Submersible remote-sensing device",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Platform use category
var pdPlatformUseCategoryR = map[string]string{
	"a": "Meteorological",
	"b": "Surface observing",
	"c": "Space observing",
	"m": "Mixed uses",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 08 - Sensor type
var pdSensorTypeR = map[string]string{
	"a": "Active",
	"b": "Passive",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 09-10 - Data type
var pdDataTypeR = map[string]string{
	"aa": "Visible light",
	"da": "Near infrared",
	"db": "Middle infrared",
	"dc": "Far infrared",
	"dd": "Thermal infrared",
	"de": "Shortwave infrared (SWIR)",
	"df": "Reflective infrared",
	"dv": "Combinations",
	"dz": "Other infrared data",
	"ga": "Sidelooking airborne radar (SLAR)",
	"gb": "Synthetic aperture radar (SAR)-Single frequency",
	"gc": "SAR-multi-frequency (multichannel)",
	"gd": "SAR-like polarization",
	"ge": "SAR-cross polarization",
	"gf": "Infometric SAR",
	"gg": "polarmetric SAR",
	"gu": "Passive microwave mapping",
	"gz": "Other microwave data",
	"ja": "Far ultraviolet",
	"jb": "Middle ultraviolet",
	"jc": "Near ultraviolet",
	"jv": "Ultraviolet combinations",
	"jz": "Other ultraviolet data",
	"ma": "Multi-spectral, multidata",
	"mb": "Multi-temporal",
	"mm": "Combination of various data types",
	// Other
	"nn": "Not applicable",
	"pa": "Sonar--water depth",
	"pb": "Sonar--bottom topography images, sidescan",
	"pc": "Sonar--bottom topography, near-surface",
	"pd": "Sonar--bottom topography, near-bottom",
	"pe": "Seismic surveys",
	"pz": "Other acoustical data",
	"ra": "Gravity anomalies (general)",
	"rb": "Free-air",
	"rc": "Bouger",
	"rd": "Isostatic",
	"sa": "Magnetic field",
	"ta": "radiometric surveys",
	"uu": "Unknown",
	"zz": "Other",
	"||": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationS = map[string]string{
	"d": "Sound disc",
	"e": "Cylinder",
	"g": "Sound cartridge",
	"i": "Sound-track film",
	"q": "Roll",
	"r": "Remote",
	"s": "Sound cassette",
	"t": "Sound-tape reel",
	"u": "Unspecified",
	"w": "Wire recording",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Speed
var pdSpeedS = map[string]string{
	"a": "16 rpm (discs)",
	"b": "33 1/3 rpm (discs)",
	"c": "45 rpm (discs)",
	"d": "78 rpm (discs)",
	"e": "8 rpm (discs)",
	"f": "1.4 m. per second (discs)",
	"h": "120 rpm (cylinders)",
	"i": "160 rpm (cylinders)",
	"k": "15/16 ips (tapes)",
	"l": "1 7/8 ips (tapes)",
	"m": "3 3/4 ips (tapes)",
	"n": "Not applicable",
	"o": "7 1/2 ips (tapes)",
	"p": "15 ips (tapes)",
	"r": "30 ips (tape)",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Configuration of playback channels
var pdPlaybackConfigS = map[string]string{
	"m": "Monaural",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Groove width/groove pitch
var pdGrooveS = map[string]string{
	"m": "Microgroove/fine",
	"n": "Not applicable",
	"s": "Coarse/standard",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 06 - Dimensions
var pdDimensionsS = map[string]string{
	"a": "3 in. diameter",
	"b": "5 in. diameter",
	"c": "7 in. diameter",
	"d": "10 in. diameter",
	"e": "12 in. diameter",
	"f": "16 in. diameter",
	"g": "4 3/4 in. or 12 cm. diameter",
	"j": "3 7/8 x 2 1/2 in.",
	"n": "Not applicable",
	"o": "5 1/4 x 3 7/8 in.",
	"s": "2 3/4 x 4 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Tape width
var pdTapeWidthS = map[string]string{
	"l": "1/8 in.",
	"m": "1/4 in.",
	"n": "Not applicable",
	"o": "1/2 in.",
	"p": "1 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 08 - Tape Configuration
var pdTapeConfigS = map[string]string{
	"a": "Full (1) track",
	"b": "Half (2) track",
	"c": "Quarter (4) track",
	"d": "Eight track",
	"e": "Twelve track",
	"f": "Sixteen track",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 09 - Kind of disc, cylinder or tape
var pdItemType = map[string]string{
	"a": "Master tape",
	"b": "Tape duplication master",
	"d": "Disc master (negative)",
	"i": "Instantaneous (recorded on the spot)",
	"m": "Mass-produced",
	"n": "Not applicable",
	"r": "Mother (positive)",
	"s": "Stamper (negative)",
	"t": "Test pressing",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 10 - Kind of material
var pdMaterialTypeS = map[string]string{
	"a": "Lacquer coating",
	"b": "Cellulose nitrate",
	"c": "Acetate tape with ferrous oxide",
	"g": "Glass with lacquer",
	"i": "Aluminum with lacquer",
	"l": "Metal",
	"m": "Plastic with metal",
	"n": "Not applicable",
	"p": "Plastic",
	"r": "Paper with lacquer or ferrous oxide",
	"s": "Shellac",
	"w": "Wax",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 11 - Kind of cutting
var pdCuttingTypeS = map[string]string{
	"h": "Hill-and-dale cutting",
	"l": "Lateral or combined cutting",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 12 - Special playback characteristics
var pdSpecialPlaybackS = map[string]string{
	"a": "NAB standard",
	"b": "CCIR standard",
	"c": "Dolby-B encoded",
	"d": "dbx encoded",
	"e": "Digital recording",
	"f": "Dolby-A encoded",
	"g": "Dolby-C encoded",
	"h": "CX encoded",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 13 - Capture and storage technique
var pdCaptureStorageS = map[string]string{
	"a": "Acoustical capture, direct storage",
	"b": "Direct storage, not acoustical",
	"d": "Digital storage",
	"e": "Analog electrical storage",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationT = map[string]string{
	"a": "Regular print",
	"b": "Large print",
	"c": "Braille",
	"d": "Loose-leaf",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 01 - Specific material designation
var pdSpecificDesignationV = map[string]string{
	"c": "Videocartridge",
	"d": "Videodisc",
	"f": "Videocassette",
	"r": "Videoreel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// 03 - Color
var pdColorV = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 04 - Videorecording format
var pdFormatV = map[string]string{
	"a": "Beta (1/2 in., videocassette)",
	"b": "VHS (1/2 in., videocassette)",
	"c": "U-matic (3/4 in., videocasstte)",
	"d": "EIAJ (1/2 in., reel)",
	"e": "Type C (1 in., reel)",
	"f": "Quadruplex (1 in. or 2 in., reel)",
	"g": "Laserdisc",
	"h": "CED (Capacitance Electronic Disc) videodisc",
	"i": "Betacam (1/2 in., videocassette)",
	"j": "Betacam SP (1/2 in., videocassette)",
	"k": "Super-VHS (1/2 in., videocassette)",
	"m": "M-II (1/2 in., videocassette)",
	"o": "D-2 (3/4 in., videocassette)",
	"p": "8 mm.",
	"q": "Hi-8 mm.",
	"s": "Blu-ray disc",
	"u": "Unknown",
	"v": "DVD",
	"z": "Other",
	"|": "No attempt to code",
}

// 05 - Sound on medium or separate
var pdSoundLocationV = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}

// 06 - Medium for sound
var pdSoundMediumV = map[string]string{
	" ": "No sound (silent)",
	"a": "Optical sound track on motion picture film",
	"b": "Magnetic sound track on motion picture film",
	"c": "Magnetic audio tape in cartridge",
	"d": "Sound disc",
	"e": "Magnetic audio tape on reel",
	"f": "Magnetic audio tape in cassette",
	"g": "Optical and magnetic sound track on motion picture film",
	"h": "Videotape",
	"i": "Videodisc",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 07 - Dimensions
var pdDimensionsV = map[string]string{
	"a": "8 mm.",
	"m": "1/4 in.",
	"o": "1/2 in.",
	"p": "1 in.",
	"q": "2 in.",
	"r": "3/4 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// 08 - Configuration of playback channels
var pdPlaybackChannelsV = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

var pdSpecificDesignationZ = map[string]string{
	"m": "Multiple physical forms",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
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

func lookupCfCode(codeList map[string]string, b []byte, i int) (v CfValue) {

	var code, label string

	if len(b) > i {
		code = string(b[i])
		if code != "" {
			label, _ = codeList[code]
		}
	}

	return CfValue{Code: code, Label: label}
}

type CfMatlDesc map[string][]CfValue

type CfPhysDesc map[string]CfValue

type CfData struct {
	ControlNumber           string
	ControlNumberIdentifier string
	LatestTransDateTime     string
	FileDate                string
	DateTypePubStatus       CfValue
	Date1                   string
	Date2                   string
	PlaceOfPublication      string
	Language                string
	ModifiedRecord          CfValue
	CatalogingSource        CfValue
	MaterialType            CfValue
	MaterialCharactor       CfMatlDesc
	PhysicalDescription     []CfPhysDesc
}

func (rec Record) parseControlfields() (c CfData) {

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

	mt, label := rec.MaterialType()

	c.MaterialType = CfValue{Code: mt, Label: label}

	c.MaterialCharactor = make(CfMatlDesc)

	switch mt {
	case "BK":
		c.MaterialCharactor.parseBookCf(b[18:35])
	case "CF":
		c.MaterialCharactor.parseComputerFilesCf(b[18:35])
	case "MP":
		c.MaterialCharactor.parseMapCf(b[18:35])
	case "MU":
		c.MaterialCharactor.parseMusicCf(b[18:35])
	case "CR":
		c.MaterialCharactor.parseContinuingResourcesCf(b[18:35])
	case "VM":
		c.MaterialCharactor.parseVisualMaterialsCf(b[18:35])
	case "MX":
		c.MaterialCharactor.parseMixedMaterialsCf(b[18:35])
	}

	cf006 := rec.getCFs("006")
	for _, x := range cf006 {
		b := []byte(x)

		// NOTE: The Material Form is b[0] and ignoring it *could*
		// disconnect the material form from the corresponding values.
		// However the material form *should* match the leader record
		// type so that really shouldn't be an issue... should it?
		// Perhaps a check to verify that b[0] matches?

		switch mt {
		case "BK":
			c.MaterialCharactor.parseBookCf(b[1:])
		case "CF":
			c.MaterialCharactor.parseComputerFilesCf(b[1:])
		case "MP":
			c.MaterialCharactor.parseMapCf(b[1:])
		case "MU":
			c.MaterialCharactor.parseMusicCf(b[1:])
		case "CR":
			c.MaterialCharactor.parseContinuingResourcesCf(b[1:])
		case "VM":
			c.MaterialCharactor.parseVisualMaterialsCf(b[1:])
		case "MX":
			c.MaterialCharactor.parseMixedMaterialsCf(b[1:])
		}
	}

	// Each 007 needs to hang as a unit.
	cf007 := rec.getCFs("007")
	for _, x := range cf007 {
		b := []byte(x)
		pd := make(CfPhysDesc)

		pd["MaterialCategory"] = lookupCfCode(materialCategory, b, 0)

		switch string(b[0]) {
		case "a":
			pd.parsePdA(b) // "Map",
		case "c":
			pd.parsePdC(b) // "Electronic resource",
		case "d":
			pd.parsePdD(b) // "Globe",
		case "f":
			pd.parsePdF(b) // "Tactile material",
		case "g":
			pd.parsePdG(b) // "Projected graphic",
		case "h":
			pd.parsePdH(b) // "Microform",
		case "k":
			pd.parsePdK(b) // "Nonprojected graphic",
		case "m":
			pd.parsePdM(b) // "Motion picture",
		case "o":
			pd.parsePdO(b) // "Kit",
		case "q":
			pd.parsePdQ(b) // "Notated music",
		case "r":
			pd.parsePdR(b) // "Remote-sensing image",
		case "s":
			pd.parsePdS(b) // "Sound recording",
		case "t":
			pd.parsePdT(b) // "Text",
		case "v":
			pd.parsePdV(b) // "Videorecording",
		case "z":
			pd.parsePdZ(b) // "Unspecified",
		}

		c.PhysicalDescription = append(c.PhysicalDescription, pd)
	}

	return c
}

func (pd CfPhysDesc) parsePdA(b []byte) {
	// Map (007/00=a)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Physical medium
	// 05 - Type of reproduction
	// 06 - Production/reproduction details
	// 07 - Positive/negative aspect
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationA, b, 1)
	pd["Color"] = lookupCfCode(pdColorA, b, 3)
	pd["PhysicalMedium"] = lookupCfCode(pdPhysicalMediumA, b, 4)
	pd["ReproductionType"] = lookupCfCode(pdReproductionTypeA, b, 5)
	pd["ProductionReproductionDetails"] = lookupCfCode(pdProductionDetailsA, b, 6)
	pd["PositiveNegativeAspect"] = lookupCfCode(pdPositiveNegativeAspectA, b, 7)
}

func (pd CfPhysDesc) parsePdC(b []byte) {
	// Electronic resource (007/00=c)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Dimensions
	// 05 - Sound
	// 06-08 - Image bit depth
	// 09 - File formats
	// 10 - Quality assurance targets
	// 11 - Antecedent/source
	// 12 - Level of compression
	// 13 - Reformatting quality
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationC, b, 1)
	pd["Color"] = lookupCfCode(pdColorC, b, 3)
	pd["Dimensions"] = lookupCfCode(pdDimensionsC, b, 4)
	pd["Sound"] = lookupCfCode(pdSoundC, b, 5)

	ibd := string(b[6:9])
	var label string
	switch ibd {
	case "mmm":
		label = "Multiple"
	case "nnn":
		label = "Not applicable"
	case "---":
		label = "Unknown"
	case "|||":
		label = "No attempt to code"
	default:
		label = "Exact bit depth"
	}

	pd["ImageBitDepth"] = CfValue{Code: ibd, Label: label}
	pd["FileFormats"] = lookupCfCode(pdFileFormatsC, b, 9)
	pd["QualityAssuranceTarget"] = lookupCfCode(pdQATargetsC, b, 10)
	pd["AntecedentSource"] = lookupCfCode(pdSourceC, b, 11)
	pd["CompressionLevel"] = lookupCfCode(pdCompressionC, b, 12)
	pd["ReformattingQuality"] = lookupCfCode(pdReformattingQualityC, b, 13)
}

func (pd CfPhysDesc) parsePdD(b []byte) {
	// Globe (007/00=d)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Physical medium
	// 05 - Type of reproduction
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationD, b, 1)
	pd["Color"] = lookupCfCode(pdColorD, b, 3)
	pd["PhysicalMedium"] = lookupCfCode(pdPhysicalMediumD, b, 4)
	pd["ReproductionType"] = lookupCfCode(pdReproductionTypeD, b, 5)
}

func (pd CfPhysDesc) parsePdF(b []byte) {
	// Tactile material (007/00=f)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03-04 - Class of braille writing
	// 05 - Level of contraction
	// 06-08 - Braille music format
	// 09 - Special physical characteristics

	// TODO:
	// 03-04 - Class of braille writing
	//var pdBrailleWritingClassF = map[string]string{
	// 06-08 - Braille music format
	//var pdBrailleMusicFormatF = map[string]string{

	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationF, b, 1)
	//	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationD, b, 1)
	pd["ContractionLevel"] = lookupCfCode(pdContractionLevelF, b, 5)
	//	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationD, b, 1)
	pd["SpecialPhysicalCharacteristics"] = lookupCfCode(pdSpecialCharacteristicsF, b, 9)
}

func (pd CfPhysDesc) parsePdG(b []byte) {
	// Projected graphic (007/00=g)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Base of emulsion
	// 05 - Sound on medium or separate
	// 06 - Medium for sound
	// 07 - Dimensions
	// 08 - Secondary support material
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationG, b, 1)
	pd["Color"] = lookupCfCode(pdColorG, b, 3)
	pd["EmulsionBase"] = lookupCfCode(pdEmulsionBaseG, b, 4)
	pd["SoundLocation"] = lookupCfCode(pdSoundLocationG, b, 5)
	pd["SoundMedium"] = lookupCfCode(pdSoundMediumG, b, 6)
	pd["Dimensions"] = lookupCfCode(pdDimensionsG, b, 7)
	pd["SecondarySupportMaterial"] = lookupCfCode(pdSecondarySupportMaterialG, b, 8)
}

func (pd CfPhysDesc) parsePdH(b []byte) {
	// Microform (007/00=h)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Positive/negative aspect
	// 04 - Dimensions
	// 05 - Reduction ratio range
	// 06-08 - Reduction ratio
	// 09 - Color
	// 10 - Emulsion on film
	// 11 - Generation
	// 12 - Base of film
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationH, b, 1)
	pd["PositiveNegativeAspect"] = lookupCfCode(pdPositiveNegativeAspectH, b, 3)
	pd["Dimensions"] = lookupCfCode(pdDimensionsH, b, 4)
	pd["ReductionRatioRange"] = lookupCfCode(pdReductionRatioRangeH, b, 5)
	pd["ReductionRatio"] = CfValue{Code: string(b[6:9]), Label: "Reduction ratio"}
	pd["Color"] = lookupCfCode(pdColorH, b, 9)
	pd["FilmEmulsion"] = lookupCfCode(pdFilmEmulsionH, b, 10)
	pd["Generation"] = lookupCfCode(pdGenerationH, b, 11)
	pd["FilmBase"] = lookupCfCode(pdFilmBaseH, b, 12)
}

func (pd CfPhysDesc) parsePdK(b []byte) {
	// Nonprojected graphic (007/00=k)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Primary support material
	// 05 - Secondary support material
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationK, b, 1)
	pd["Color"] = lookupCfCode(pdColorK, b, 3)
	pd["PrimarySupportMaterial"] = lookupCfCode(pdPrimarySupportMaterialK, b, 4)
	pd["SecondarySupportMaterial"] = lookupCfCode(pdSecondarySupportMaterialK, b, 5)
}

func (pd CfPhysDesc) parsePdM(b []byte) {
	// Motion picture (007/00=m)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Motion picture presentation format
	// 05 - Sound on medium or separate
	// 06 - Medium for sound
	// 07 - Dimensions
	// 08 - Configuration of playback channels
	// 09 - Production elements
	// 10 - Positive/negative aspect
	// 11 - Generation
	// 12 - Base of film
	// 13 - Refined categories of color
	// 14 - Kind of color stock or print
	// 15 - Deterioration stage
	// 16 - Completeness
	// 17-22 - Film inspection date
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationM, b, 1)
	pd["Color"] = lookupCfCode(pdColorM, b, 3)
	pd["PresentationFormat"] = lookupCfCode(pdPresentationFormatM, b, 4)
	pd["SoundLocation"] = lookupCfCode(pdSoundLocationM, b, 5)
	pd["SoundMedium"] = lookupCfCode(pdSoundMediumM, b, 6)
	pd["Dimensions"] = lookupCfCode(pdDimensionsM, b, 7)
	pd["PlaybackChannelConfig"] = lookupCfCode(pdPlaybackChannelConfigM, b, 8)
	pd["ProductionElements"] = lookupCfCode(pdProductionElementsM, b, 9)
	pd["PositiveNegativeAspect"] = lookupCfCode(pdPositiveNegativeAspectM, b, 10)
	pd["Generation"] = lookupCfCode(pdGenerationM, b, 11)
	pd["FilmBase"] = lookupCfCode(pdFilmBaseM, b, 12)
	pd["RefinedColorCategories"] = lookupCfCode(pdColorCategoriesM, b, 13)
	pd["ColorType"] = lookupCfCode(pdColorTypeM, b, 14)
	pd["DeteriorationStage"] = lookupCfCode(pdDeteriorationStageM, b, 15)
	pd["Completeness"] = lookupCfCode(pdCompletenessM, b, 16)

	pd["FilmInspectionDate"] = CfValue{Code: string(b[17:]), Label: "Film inspection date"}
}

func (pd CfPhysDesc) parsePdO(b []byte) {
	// Kit (007/00=o)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationO, b, 1)
}

func (pd CfPhysDesc) parsePdQ(b []byte) {
	// Notated music (007/00=q)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationQ, b, 1)
}

func (pd CfPhysDesc) parsePdR(b []byte) {
	// Remote-sensing image (007/00=r)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Altitude of sensor
	// 04 - Attitude of sensor
	// 05 - Cloud cover
	// 06 - Platform construction type
	// 07 - Platform use category
	// 08 - Sensor type
	// 09-10 - Data type

	// 09-10 - Data type
	// var pdDataTypeR = map[string]string{
	// "aa": "Visible light",

	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationR, b, 1)
	pd["SensorAltitude"] = lookupCfCode(pdSensorAltitudeR, b, 3)
	pd["SensorAttitudeR"] = lookupCfCode(pdSensorAttitudeR, b, 4)
	pd["CloudCover"] = lookupCfCode(pdCloudCoverR, b, 5)
	pd["PlatformConstructionType"] = lookupCfCode(pdPlatformConstructionTypeR, b, 6)
	pd["PlatformUseCategory"] = lookupCfCode(pdPlatformUseCategoryR, b, 7)
	pd["SensorType"] = lookupCfCode(pdSensorTypeR, b, 8)

	// TODO: two char code
	//pd["DataType"] = lookupCfCode(pdDataTypeR, b, 1)

}

func (pd CfPhysDesc) parsePdS(b []byte) {
	// Sound recording (007/00=s)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Speed
	// 04 - Configuration of playback channels
	// 05 - Groove width/groove pitch
	// 06 - Dimensions
	// 07 - Tape width
	// 08 - Tape configuration
	// 09 - Kind of disc, cylinder, or tape
	// 10 - Kind of material
	// 11 - Kind of cutting
	// 12 - Special playback characteristics
	// 13 - Capture and storage technique
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationS, b, 1)
	pd["Speed"] = lookupCfCode(pdSpeedS, b, 3)
	pd["PlaybackChannelConfig"] = lookupCfCode(pdPlaybackConfigS, b, 4)
	pd["GrooveWidthPitch"] = lookupCfCode(pdGrooveS, b, 5)
	pd["Dimensions"] = lookupCfCode(pdDimensionsS, b, 6)
	pd["TapeWidth"] = lookupCfCode(pdTapeWidthS, b, 7)
	pd["TapeConfiguration"] = lookupCfCode(pdTapeConfigS, b, 8)
	pd["DiscCylinderOrTapeType"] = lookupCfCode(pdItemType, b, 9)
	pd["MaterialType"] = lookupCfCode(pdMaterialTypeS, b, 10)
	pd["CuttingType"] = lookupCfCode(pdCuttingTypeS, b, 11)
	pd["SpecialPlaybackCharacteristics"] = lookupCfCode(pdSpecialPlaybackS, b, 12)
	pd["CaptureAndStorageTechnique"] = lookupCfCode(pdCaptureStorageS, b, 13)
}

func (pd CfPhysDesc) parsePdT(b []byte) {
	// Text (007/00=t)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationT, b, 1)
}

func (pd CfPhysDesc) parsePdV(b []byte) {
	// Videorecording (007/00=v)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Videorecording format
	// 05 - Sound on medium or separate
	// 06 - Medium for sound
	// 07 - Dimensions
	// 08 - Configuration of playback channels
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationV, b, 1)
	pd["Color"] = lookupCfCode(pdColorV, b, 3)
	pd["VideorecordingFormat"] = lookupCfCode(pdFormatV, b, 4)
	pd["SoundLocation"] = lookupCfCode(pdSoundLocationV, b, 5)
	pd["SoundMedium"] = lookupCfCode(pdSoundMediumV, b, 6)
	pd["Dimensions"] = lookupCfCode(pdDimensionsV, b, 7)
	pd["PlaybackChannelConfig"] = lookupCfCode(pdPlaybackChannelsV, b, 8)
}

func (pd CfPhysDesc) parsePdZ(b []byte) {
	// Unspecified (007/00=z)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = lookupCfCode(pdSpecificDesignationZ, b, 1)
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
		mc.appendCfMatlDesc("Illustrations", lookupCfCode(illustrations, b, i))
	}

	mc.appendCfMatlDesc("TargetAudience", lookupCfCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 5))

	for i := 6; i <= 9; i++ {
		mc.appendCfMatlDesc("NatureOfContents", lookupCfCode(natureOfContents, b, i))
	}

	mc.appendCfMatlDesc("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ConferencePublication", lookupCfCode(conferencePublication, b, 11))
	mc.appendCfMatlDesc("Festschrift", lookupCfCode(festschrift, b, 12))
	mc.appendCfMatlDesc("Index", lookupCfCode(index, b, 13))
	mc.appendCfMatlDesc("LiteraryForm", lookupCfCode(literaryForm, b, 15))
	mc.appendCfMatlDesc("Biography", lookupCfCode(biography, b, 16))
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

	mc.appendCfMatlDesc("TargetAudience", lookupCfCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(computerItemForm, b, 5))
	mc.appendCfMatlDesc("ComputerFileType", lookupCfCode(computerFileType, b, 8))
	mc.appendCfMatlDesc("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
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
		mc.appendCfMatlDesc("Relief", lookupCfCode(mapRelief, b, i))
	}

	// TODO: this is wrong, it's a two char code
	//for i := 4; i <= 5; i++ {
	//	mc.appendCfMatlDesc("Projection", lookupCfCode(mapProjection, b, i))
	//}

	mc.appendCfMatlDesc("CartographicMaterialType", lookupCfCode(cartographicMaterialType, b, 7))
	mc.appendCfMatlDesc("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 11))
	mc.appendCfMatlDesc("Index", lookupCfCode(index, b, 13))

	for i := 15; i <= 16; i++ {
		mc.appendCfMatlDesc("SpecialFormat", lookupCfCode(mapSpecialFormat, b, i))
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

	// TODO: this is wrong: it's a two char code
	//for i := 0; i <= 1; i++ {
	//	mc.appendCfMatlDesc("CompositionForm", lookupCfCode(compositionForm, b, i))
	//}

	mc.appendCfMatlDesc("MusicFormat", lookupCfCode(musicFormat, b, 2))
	mc.appendCfMatlDesc("MusicParts", lookupCfCode(musicParts, b, 3))
	mc.appendCfMatlDesc("TargetAudience", lookupCfCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 5))

	for i := 6; i <= 11; i++ {
		mc.appendCfMatlDesc("AccompanyingMatter", lookupCfCode(accompanyingMatter, b, i))
	}

	for i := 12; i <= 13; i++ {
		mc.appendCfMatlDesc("RecordingLiteraryText", lookupCfCode(recordingLiteraryText, b, i))
	}

	mc.appendCfMatlDesc("TranspositionAndArrangement", lookupCfCode(transpositionAndArrangement, b, 15))
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

	mc.appendCfMatlDesc("Frequency", lookupCfCode(frequency, b, 0))
	mc.appendCfMatlDesc("Regularity", lookupCfCode(regularity, b, 1))
	mc.appendCfMatlDesc("ContinuingResourceType", lookupCfCode(continuingResourceType, b, 3))
	mc.appendCfMatlDesc("OriginalItemForm", lookupCfCode(itemForm, b, 4))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 5))
	mc.appendCfMatlDesc("NatureOfWork", lookupCfCode(natureOfContents, b, 6))

	for i := 7; i <= 9; i++ {
		mc.appendCfMatlDesc("NatureOfContents", lookupCfCode(natureOfContents, b, i))
	}
	mc.appendCfMatlDesc("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ConferencePublication", lookupCfCode(conferencePublication, b, 11))
	mc.appendCfMatlDesc("OriginalScriptOfTitle", lookupCfCode(originalScriptOfTitle, b, 15))
	mc.appendCfMatlDesc("EntryConvention", lookupCfCode(entryConvention, b, 16))
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
	mc.appendCfMatlDesc("RunningTime", CfValue{Code: rt, Label: rl})

	mc.appendCfMatlDesc("TargetAudience", lookupCfCode(targetAudience, b, 4))
	mc.appendCfMatlDesc("GovernmentPublication", lookupCfCode(governmentPublication, b, 10))
	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 11))
	mc.appendCfMatlDesc("VisualMaterialType", lookupCfCode(visualMaterialType, b, 15))
	mc.appendCfMatlDesc("Technique", lookupCfCode(technique, b, 16))
}

// parseMixedMaterialsCf parses the control field data for MixedMaterials (MU) material types.
func (mc CfMatlDesc) parseMixedMaterialsCf(b []byte) {

	// 008:
	//  18-22 - Undefined
	//  23 - Form of item (006/06)
	//  24-34 - Undefined

	mc.appendCfMatlDesc("ItemForm", lookupCfCode(itemForm, b, 5))
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
