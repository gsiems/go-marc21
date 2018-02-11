package marc21

////////////////////////////////////////////////////////////////////////
// Authority
////////////////////////////////////////////////////////////////////////
// Authority -- 008
var authority008DirectOrIndirectGeographicSubdivision = map[string]string{
	" ": "Not subdivided geographically",
	"d": "Subdivided geographically--direct",
	"i": "Subdivided geographically--indirect",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008RomanizationScheme = map[string]string{
	"a": "International standard",
	"b": "National standard",
	"c": "National library association standard",
	"d": "National library or bibliographic agency standard",
	"e": "Local standard",
	"f": "Standard of unknown origin",
	"g": "Conventional romanization or conventional form of name in language of cataloging agency",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008LanguageOfCatalog = map[string]string{
	" ": "No information provided",
	"b": "English and French",
	"e": "English only",
	"f": "French only",
	"|": "No attempt to code",
}
var authority008KindOfRecord = map[string]string{
	"a": "Established heading",
	"b": "Untraced reference",
	"c": "Traced reference",
	"d": "Subdivision",
	"e": "Node label",
	"f": "Established heading and subdivision",
	"g": "Reference and subdivision",
}
var authority008DescriptiveCatalogingRules = map[string]string{
	"a": "Earlier rules",
	"b": "AACR 1",
	"c": "AACR 2",
	"d": "AACR 2 compatible heading",
	"n": "Not applicable",
	"z": "Other",
	"|": "No attempt to code",
}
var authority008SubjectHeadingSystemThesaurus = map[string]string{
	"a": "Library of Congress Subject Headings",
	"b": "LC subject headings for children's literature",
	"c": "Medical Subject Headings",
	"d": "National Agricultural Library subject authority file",
	"k": "Canadian Subject Headings",
	"n": "Not applicable",
	"r": "Art and Architecture Thesaurus",
	"s": "Sears List of Subject Headings",
	"v": "R&eacute;pertoire de vedettes-mati&egrave;re",
	"z": "Other",
	"|": "No attempt to code",
}
var authority008TypeOfSeries = map[string]string{
	"a": "Monographic series",
	"b": "Multipart item",
	"c": "Series-like phrase",
	"n": "Not applicable",
	"z": "Other",
	"|": "No attempt to code",
}
var authority008NumberedOrUnnumberedSeries = map[string]string{
	"a": "Numbered",
	"b": "Unnumbered",
	"c": "Numbering varies",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008HeadingUseMainOrAddedEntry = map[string]string{
	"a": "Appropriate",
	"b": "Not appropriate",
	"|": "No attempt to code",
}
var authority008HeadingUseSubjectAddedEntry = map[string]string{
	"a": "Appropriate",
	"b": "Not appropriate",
	"|": "No attempt to code",
}
var authority008HeadingUseSeriesAddedEntry = map[string]string{
	"a": "Appropriate",
	"b": "Not appropriate",
	"|": "No attempt to code",
}
var authority008TypeOfSubjectSubdivision = map[string]string{
	"a": "Topical",
	"b": "Form",
	"c": "Chronological",
	"d": "Geographic",
	"e": "Language",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008TypeOfGovernmentAgency = map[string]string{
	" ": "Not a government agency",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government agency--type undetermined",
	"s": "State, provincial, territorial, dependent, etc.",
	"u": "Unknown if heading is government agency",
	"z": "Other",
	"|": "No attempt to code",
}
var authority008ReferenceEvaluation = map[string]string{
	"a": "Tracings are consistent with the heading",
	"b": "Tracings are not necessarily consistent with the heading",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008RecordUpdateInProcess = map[string]string{
	"a": "Record can be used",
	"b": "Record is being updated",
	"|": "No attempt to code",
}
var authority008UndifferentiatedPersonalName = map[string]string{
	"a": "Differentiated personal name",
	"b": "Undifferentiated personal name",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008LevelOfEstablishment = map[string]string{
	"a": "Fully established",
	"b": "Memorandum",
	"c": "Provisional",
	"d": "Preliminary",
	"n": "Not applicable",
	"|": "No attempt to code",
}
var authority008ModifiedRecord = map[string]string{
	" ": "Not modified",
	"s": "Shortened",
	"x": "Missing characters",
	"|": "No attempt to code",
}
var authority008CatalogingSource = map[string]string{
	" ": "National bibliographic agency",
	"c": "Cooperative cataloging program",
	"d": "Other",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseAuthority008 parses the 008 control field data for
// Authority records data
func parseAuthority008(d *Cf008Desc, s string) {

	d.append("(00/06) Date entered on file", CodeValue{Code: pluckBytes(s, 0, 6), Label: ""})
	d.append("(06/01) Direct or indirect geographic subdivision", codeLookup(authority008DirectOrIndirectGeographicSubdivision, s, 6, 1))
	d.append("(07/01) Romanization scheme", codeLookup(authority008RomanizationScheme, s, 7, 1))
	d.append("(08/01) Language of catalog", codeLookup(authority008LanguageOfCatalog, s, 8, 1))
	d.append("(09/01) Kind of record", codeLookup(authority008KindOfRecord, s, 9, 1))
	d.append("(10/01) Descriptive cataloging rules", codeLookup(authority008DescriptiveCatalogingRules, s, 10, 1))
	d.append("(11/01) Subject heading system/thesaurus", codeLookup(authority008SubjectHeadingSystemThesaurus, s, 11, 1))
	d.append("(12/01) Type of series", codeLookup(authority008TypeOfSeries, s, 12, 1))
	d.append("(13/01) Numbered or unnumbered series", codeLookup(authority008NumberedOrUnnumberedSeries, s, 13, 1))
	d.append("(14/01) Heading use--main or added entry", codeLookup(authority008HeadingUseMainOrAddedEntry, s, 14, 1))
	d.append("(15/01) Heading use--subject added entry", codeLookup(authority008HeadingUseSubjectAddedEntry, s, 15, 1))
	d.append("(16/01) Heading use--series added entry", codeLookup(authority008HeadingUseSeriesAddedEntry, s, 16, 1))
	d.append("(17/01) Type of subject subdivision", codeLookup(authority008TypeOfSubjectSubdivision, s, 17, 1))
	d.append("(18/10) Undefined character positions", CodeValue{Code: pluckBytes(s, 18, 10), Label: ""})
	d.append("(28/01) Type of government agency", codeLookup(authority008TypeOfGovernmentAgency, s, 28, 1))
	d.append("(29/01) Reference evaluation", codeLookup(authority008ReferenceEvaluation, s, 29, 1))
	d.append("(30/01) Undefined character position", CodeValue{Code: pluckBytes(s, 30, 1), Label: ""})
	d.append("(31/01) Record update in process", codeLookup(authority008RecordUpdateInProcess, s, 31, 1))
	d.append("(32/01) Undifferentiated personal name", codeLookup(authority008UndifferentiatedPersonalName, s, 32, 1))
	d.append("(33/01) Level of establishment", codeLookup(authority008LevelOfEstablishment, s, 33, 1))
	d.append("(34/04) Undefined character positions", CodeValue{Code: pluckBytes(s, 34, 4), Label: ""})
	d.append("(38/01) Modified record", codeLookup(authority008ModifiedRecord, s, 38, 1))
	d.append("(39/01) Cataloging source", codeLookup(authority008CatalogingSource, s, 39, 1))
}

////////////////////////////////////////////////////////////////////////
// Bibliography
////////////////////////////////////////////////////////////////////////
// Bibliography -- 007
var bibliography007CategoryOfMaterial = map[string]string{
	"a": "Map",
	"c": "Computer file",
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

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- MAP
var bibliography007MAPSpecificMaterialDesignation = map[string]string{
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
var bibliography007MAPColor = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}
var bibliography007MAPPhysicalMedium = map[string]string{
	"a": "Paper",
	"b": "Wood",
	"c": "Stone",
	"d": "Metal",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textile",
	"i": "Plastic",
	"j": "Glass",
	"l": "Vinyl",
	"n": "Vellum",
	"p": "Plaster",
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
var bibliography007MAPTypeOfReproduction = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MAPProductionReproductionDetails = map[string]string{
	"a": "Photocopy, blueline print",
	"b": "Photocopy",
	"c": "Pre-production",
	"d": "Film",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MAPPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"n": "Not applicable",
	"|": "No attempt to code",
}

// parseBibliography007MAP parses the 007 control field data for
// Bibliography records MAP (MAP) data
func parseBibliography007MAP(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007MAPSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007MAPColor, s, 3, 1)
	pd["(04/01) Physical medium"] = codeLookup(bibliography007MAPPhysicalMedium, s, 4, 1)
	pd["(05/01) Type of reproduction"] = codeLookup(bibliography007MAPTypeOfReproduction, s, 5, 1)
	pd["(06/01) Production/reproduction details"] = codeLookup(bibliography007MAPProductionReproductionDetails, s, 6, 1)
	pd["(07/01) Positive/negative aspect"] = codeLookup(bibliography007MAPPositiveNegativeAspect, s, 7, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- ELECTRONIC RESOURCE
var bibliography007ELRSpecificMaterialDesignation = map[string]string{
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
var bibliography007ELRColor = map[string]string{
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
var bibliography007ELRDimensions = map[string]string{
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
var bibliography007ELRSound = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007ELRImageBitDepth = map[string]string{
	"001-999": "Exact bit depth",
	"mmm":     "Multiple",
	"nnn":     "Not applicable",
	"---":     "Unknown",
	"|||":     "No attempt to code",
}
var bibliography007ELRFileFormats = map[string]string{
	"a": "One",
	"m": "Multiple",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007ELRQualityAssuranceTargetS = map[string]string{
	"a": "Absent",
	"n": "Not applicable",
	"p": "Present",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007ELRAntecedentSource = map[string]string{
	"a": "File reproduced from original",
	"b": "File reproduced from microform",
	"c": "File reproduced from an electronic resource",
	"d": "File reproduced from an intermediate (not microform)",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007ELRLevelOfCompression = map[string]string{
	"a": "Uncompressed",
	"b": "Lossless",
	"d": "Lossy",
	"m": "Mixed",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007ELRReformattingQuality = map[string]string{
	"a": "Access",
	"n": "Not applicable",
	"p": "Preservation",
	"r": "Replacement",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseBibliography007ELR parses the 007 control field data for
// Bibliography records ELECTRONIC RESOURCE (ELR) data
func parseBibliography007ELR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007ELRSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007ELRColor, s, 3, 1)
	pd["(04/01) Dimensions"] = codeLookup(bibliography007ELRDimensions, s, 4, 1)
	pd["(05/01) Sound"] = codeLookup(bibliography007ELRSound, s, 5, 1)

	rt06 := codeLookup(bibliography007ELRImageBitDepth, s, 6, 3)
	if rt06.Code != "" && rt06.Label == "" {
		rt06.Label = "Exact bit depth"
	}
	pd["(06/03) Image bit depth"] = rt06

	pd["(09/01) File formats"] = codeLookup(bibliography007ELRFileFormats, s, 9, 1)
	pd["(10/01) Quality assurance target(s)"] = codeLookup(bibliography007ELRQualityAssuranceTargetS, s, 10, 1)
	pd["(11/01) Antecedent/source"] = codeLookup(bibliography007ELRAntecedentSource, s, 11, 1)
	pd["(12/01) Level of compression"] = codeLookup(bibliography007ELRLevelOfCompression, s, 12, 1)
	pd["(13/01) Reformatting quality"] = codeLookup(bibliography007ELRReformattingQuality, s, 13, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- GLOBE
var bibliography007GLBSpecificMaterialDesignation = map[string]string{
	"a": "Celestial globe",
	"b": "Planetary or lunar globe",
	"c": "Terrestrial globe",
	"e": "Earth moon globe",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007GLBColor = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}
var bibliography007GLBPhysicalMedium = map[string]string{
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
var bibliography007GLBTypeOfReproduction = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007GLB parses the 007 control field data for
// Bibliography records GLOBE (GLB) data
func parseBibliography007GLB(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007GLBSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007GLBColor, s, 3, 1)
	pd["(04/01) Physical medium"] = codeLookup(bibliography007GLBPhysicalMedium, s, 4, 1)
	pd["(05/01) Type of reproduction"] = codeLookup(bibliography007GLBTypeOfReproduction, s, 5, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- TACTILE MATERIAL
var bibliography007TAMSpecificMaterialDesignation = map[string]string{
	"a": "Moon",
	"b": "Braille",
	"c": "Combination",
	"d": "Tactile, with no writing system",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007TAMClassOfBrailleWriting = map[string]string{
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
var bibliography007TAMLevelOfContraction = map[string]string{
	"a": "Uncontracted",
	"b": "Contracted",
	"m": "Combination",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007TAMBrailleMusicFormat = map[string]string{
	" ": "No specified braille music format",
	"a": "Bar over bar",
	"b": "Bar by bar",
	"c": "Line over line",
	"d": "Paragraph",
	"e": "Single line",
	"f": "Section by section",
	"g": "Line by line",
	"h": "Open score",
	"i": "Spanner short form scoring",
	"j": "Short form scoring",
	"k": "Outline",
	"l": "Vertical score",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007TAMSpecificPhysicalCharacteristics = map[string]string{
	"a": "Print/braille",
	"b": "Jumbo or enlarged braille",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007TAM parses the 007 control field data for
// Bibliography records TACTILE MATERIAL (TAM) data
func parseBibliography007TAM(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007TAMSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}

	pd["(03/02) Class of braille writing - 1"] = codeLookup(bibliography007TAMClassOfBrailleWriting, s, 3, 1)
	pd["(03/02) Class of braille writing - 2"] = codeLookup(bibliography007TAMClassOfBrailleWriting, s, 4, 1)

	pd["(05/01) Level of contraction"] = codeLookup(bibliography007TAMLevelOfContraction, s, 5, 1)

	pd["(06/03) Braille music format - 1"] = codeLookup(bibliography007TAMBrailleMusicFormat, s, 6, 1)
	pd["(06/03) Braille music format - 2"] = codeLookup(bibliography007TAMBrailleMusicFormat, s, 7, 1)
	pd["(06/03) Braille music format - 3"] = codeLookup(bibliography007TAMBrailleMusicFormat, s, 8, 1)

	pd["(09/01) Specific physical characteristics"] = codeLookup(bibliography007TAMSpecificPhysicalCharacteristics, s, 9, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- PROJECTED GRAPHIC
var bibliography007PRGSpecificMaterialDesignation = map[string]string{
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
var bibliography007PRGColor = map[string]string{
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
var bibliography007PRGBaseOfEmulsion = map[string]string{
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
var bibliography007PRGSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007PRGMediumForSound = map[string]string{
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
var bibliography007PRGDimensions = map[string]string{
	"a": "Standard 8 mm.",
	"b": "Super 8 mm./single 8 mm.",
	"c": "9.5 mm.",
	"d": "16 mm.",
	"e": "28 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"j": "2x2 in. or 5x5 cm.",
	"k": "2 1/4 x 2 1/4 in. or 6x6 cm.",
	"s": "4x5 in. or 10x13 cm.",
	"t": "5x7 in. or 13x18 cm.",
	"u": "Unknown",
	"v": "8x10 in. or 21x26 cm.",
	"w": "9x9 in. or 23x23 cm.",
	"x": "10x10 in. or 26x26 cm.",
	"y": "7x7 in. or 18x18 cm.",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007PRGSecondarySupportMaterial = map[string]string{
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

// parseBibliography007PRG parses the 007 control field data for
// Bibliography records PROJECTED GRAPHIC (PRG) data
func parseBibliography007PRG(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007PRGSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007PRGColor, s, 3, 1)
	pd["(04/01) Base of emulsion"] = codeLookup(bibliography007PRGBaseOfEmulsion, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(bibliography007PRGSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(bibliography007PRGMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(bibliography007PRGDimensions, s, 7, 1)
	pd["(08/01) Secondary support material"] = codeLookup(bibliography007PRGSecondarySupportMaterial, s, 8, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- MICROFORM
var bibliography007MICSpecificMaterialDesignation = map[string]string{
	"a": "Aperture card",
	"b": "Microfilm cartridge",
	"c": "Microfilm cassette",
	"d": "Microfilm reel",
	"e": "Microfiche",
	"f": "Microfiche cassette",
	"g": "Microopaque",
	"h": "Microfilm slip",
	"j": "Mircrofilm roll",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MICPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007MICDimensions = map[string]string{
	"a": "8 mm.",
	"d": "16 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"h": "105 mm.",
	"l": "3x5 in. or 8x13 cm.",
	"m": "4x6 in. or 11x15 cm.",
	"o": "6x9 in. or 16x23 cm.",
	"p": "3 1/4 x 7 3/8 in. or 9x19 cm.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MICReductionRatioRange = map[string]string{
	"a": "Low reduction",
	"b": "Normal reduction",
	"c": "High reduction",
	"d": "Very high reduction",
	"e": "Ultra high reduction",
	"u": "Unknown",
	"v": "Reduction rate varies",
	"|": "No attempt to code",
}
var bibliography007MICColor = map[string]string{
	"b": "Black-and-white (or monochrome)",
	"c": "Multicolored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MICEmulsionOnFilm = map[string]string{
	"a": "Silver halide",
	"b": "Diazo",
	"c": "Vesicular",
	"m": "Mixed emulsion",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MICGeneration = map[string]string{
	"a": "First generation (master)",
	"b": "Printing master",
	"c": "Service copy",
	"m": "Mixed generation",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007MICBaseOfFilm = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007MIC parses the 007 control field data for
// Bibliography records MICROFORM (MIC) data
func parseBibliography007MIC(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007MICSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Positive/negative aspect"] = codeLookup(bibliography007MICPositiveNegativeAspect, s, 3, 1)
	pd["(04/01) Dimensions"] = codeLookup(bibliography007MICDimensions, s, 4, 1)
	pd["(05/01) Reduction ratio range"] = codeLookup(bibliography007MICReductionRatioRange, s, 5, 1)
	pd["(06/03) Reduction ratio"] = CodeValue{Code: pluckBytes(s, 6, 3), Label: ""}
	pd["(09/01) Color"] = codeLookup(bibliography007MICColor, s, 9, 1)
	pd["(10/01) Emulsion on film"] = codeLookup(bibliography007MICEmulsionOnFilm, s, 10, 1)
	pd["(11/01) Generation"] = codeLookup(bibliography007MICGeneration, s, 11, 1)
	pd["(12/01) Base of film"] = codeLookup(bibliography007MICBaseOfFilm, s, 12, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- NONPROJECTED GRAPHIC
var bibliography007NPGSpecificMaterialDesignation = map[string]string{
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
var bibliography007NPGColor = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007NPGPrimarySupportMaterial = map[string]string{
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
var bibliography007NPGSecondarySupportMaterial = map[string]string{
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

// parseBibliography007NPG parses the 007 control field data for
// Bibliography records NONPROJECTED GRAPHIC (NPG) data
func parseBibliography007NPG(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007NPGSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007NPGColor, s, 3, 1)
	pd["(04/01) Primary support material"] = codeLookup(bibliography007NPGPrimarySupportMaterial, s, 4, 1)
	pd["(05/01) Secondary support material"] = codeLookup(bibliography007NPGSecondarySupportMaterial, s, 5, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- MOTION PICTURE
var bibliography007MOPSpecificMaterialDesignation = map[string]string{
	"c": "Film cartridge",
	"f": "Film cassette",
	"o": "Film roll",
	"r": "Film reel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPColor = map[string]string{
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPMotionPicturePresentationFormat = map[string]string{
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
var bibliography007MOPSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007MOPMediumForSound = map[string]string{
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
var bibliography007MOPDimensions = map[string]string{
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
var bibliography007MOPConfigurationOfPlaybackChannels = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPProductionElements = map[string]string{
	"a": "Workprint",
	"b": "Trims",
	"c": "Outtakes",
	"d": "Rushes",
	"e": "Mixing tracks",
	"f": "Title bands/intertitle rolls",
	"g": "Production rolls",
	"n": "Not applicable",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPGeneration = map[string]string{
	"d": "Duplicate",
	"e": "Master",
	"o": "Original",
	"r": "Reference print/viewing copy",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPBaseOfFilm = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPRefinedCategoriesOfColor = map[string]string{
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
var bibliography007MOPKindOfColorStockOrPrint = map[string]string{
	"a": "Imbibition dye transfer prints",
	"b": "Three layer stock",
	"c": "Three layer stock, low fade",
	"d": "Duplitized stock",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007MOPDeteriorationStage = map[string]string{
	"a": "None apparent",
	"b": "Nitrate: suspicious odor",
	"c": "Nitrate: pungent odor",
	"d": "Nitrate: brownish, discoloration, fading, dusty",
	"e": "Nitrate: sticky",
	"f": "Nitrate: frothy, bubbles, blisters",
	"g": "Nitrate: congealed",
	"h": "Nitrate: powder",
	"k": "Non-nitrate: detectable deterioration (diacetate odor)",
	"l": "Non-nitrate: advanced deterioration",
	"m": "Non-nitrate: disaster",
	"|": "No attempt to code",
}
var bibliography007MOPCompleteness = map[string]string{
	"c": "Complete",
	"i": "Incomplete",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseBibliography007MOP parses the 007 control field data for
// Bibliography records MOTION PICTURE (MOP) data
func parseBibliography007MOP(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007MOPSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007MOPColor, s, 3, 1)
	pd["(04/01) Motion picture presentation format"] = codeLookup(bibliography007MOPMotionPicturePresentationFormat, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(bibliography007MOPSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(bibliography007MOPMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(bibliography007MOPDimensions, s, 7, 1)
	pd["(08/01) Configuration of playback channels"] = codeLookup(bibliography007MOPConfigurationOfPlaybackChannels, s, 8, 1)
	pd["(09/01) Production elements"] = codeLookup(bibliography007MOPProductionElements, s, 9, 1)
	pd["(10/01) Positive/negative aspect"] = codeLookup(bibliography007MOPPositiveNegativeAspect, s, 10, 1)
	pd["(11/01) Generation"] = codeLookup(bibliography007MOPGeneration, s, 11, 1)
	pd["(12/01) Base of film"] = codeLookup(bibliography007MOPBaseOfFilm, s, 12, 1)
	pd["(13/01) Refined categories of color"] = codeLookup(bibliography007MOPRefinedCategoriesOfColor, s, 13, 1)
	pd["(14/01) Kind of color stock or print"] = codeLookup(bibliography007MOPKindOfColorStockOrPrint, s, 14, 1)
	pd["(15/01) Deterioration stage"] = codeLookup(bibliography007MOPDeteriorationStage, s, 15, 1)
	pd["(16/01) Completeness"] = codeLookup(bibliography007MOPCompleteness, s, 16, 1)
	pd["(17/06) Film inspection date"] = CodeValue{Code: pluckBytes(s, 17, 6), Label: ""}

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- KIT
var bibliography007KITSpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// parseBibliography007KIT parses the 007 control field data for
// Bibliography records KIT (KIT) data
func parseBibliography007KIT(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007KITSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- NOTATED MUSIC
var bibliography007NMUSpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// parseBibliography007NMU parses the 007 control field data for
// Bibliography records NOTATED MUSIC (NMU) data
func parseBibliography007NMU(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007NMUSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- REMOTE-SENSING IMAGE
var bibliography007RSISpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}
var bibliography007RSIAltitudeOfSensor = map[string]string{
	"a": "Surface",
	"b": "Airborne",
	"c": "Spaceborne",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007RSIAttitudeOfSensor = map[string]string{
	"a": "Low oblique",
	"b": "High oblique",
	"c": "Vertical",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007RSICloudCover = map[string]string{
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
var bibliography007RSIPlatformConstructionType = map[string]string{
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
var bibliography007RSIPlatformUseCategory = map[string]string{
	"a": "Meteorological",
	"b": "Surface observing",
	"c": "Space observing",
	"m": "Mixed uses",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007RSISensorType = map[string]string{
	"a": "Active",
	"b": "Passive",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007RSIDataType = map[string]string{
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
	"nn": "Not applicable",
	"pa": "Sonar--water depth",
	"pb": "Sonar--bottom topography images, sidescan",
	"pc": "Sonar--bottom topography, near surface",
	"pd": "Sonar--bottom topography, near bottom",
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

// parseBibliography007RSI parses the 007 control field data for
// Bibliography records REMOTE-SENSING IMAGE (RSI) data
func parseBibliography007RSI(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007RSISpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Altitude of sensor"] = codeLookup(bibliography007RSIAltitudeOfSensor, s, 3, 1)
	pd["(04/01) Attitude of sensor"] = codeLookup(bibliography007RSIAttitudeOfSensor, s, 4, 1)
	pd["(05/01) Cloud cover"] = codeLookup(bibliography007RSICloudCover, s, 5, 1)
	pd["(06/01) Platform construction type"] = codeLookup(bibliography007RSIPlatformConstructionType, s, 6, 1)
	pd["(07/01) Platform use category"] = codeLookup(bibliography007RSIPlatformUseCategory, s, 7, 1)
	pd["(08/01) Sensor type"] = codeLookup(bibliography007RSISensorType, s, 8, 1)
	pd["(09/02) Data type"] = codeLookup(bibliography007RSIDataType, s, 9, 2)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- SOUND RECORDING
var bibliography007SORSpecificMaterialDesignation = map[string]string{
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
var bibliography007SORSpeed = map[string]string{
	"a": "16 rpm",
	"b": "33 1/3 rpm",
	"c": "45 rpm",
	"d": "78 rpm",
	"e": "8 rpm",
	"f": "1.4 m. per sec.",
	"h": "120 rpm",
	"i": "160 rpm",
	"k": "15/16 ips",
	"l": "1 7/8 ips",
	"m": "3 3/4 ips",
	"n": "Not applicable",
	"o": "7 1/2 ips",
	"p": "15 ips",
	"r": "30 ips",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORConfigurationOfPlaybackChannels = map[string]string{
	"m": "Monaural",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORGrooveWidthGroovePitch = map[string]string{
	"m": "Microgroove/fine",
	"n": "Not applicable",
	"s": "Coarse/standard",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORDimensions = map[string]string{
	"a": "3 in.",
	"b": "5 in.",
	"c": "7 in.",
	"d": "10 in.",
	"e": "12 in.",
	"f": "16 in.",
	"g": "4 3/4 in. or 12 cm.",
	"j": "3 7/8 x 2 1/2 in.",
	"o": "5 1/4 x 3 7/8 in.",
	"n": "Not applicable",
	"s": "2 3/4 x 4 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORTapeWidth = map[string]string{
	"l": "1/8 in.",
	"m": "1/4 in.",
	"n": "Not applicable",
	"o": "1/2 in.",
	"p": "1 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORTapeConfiguration = map[string]string{
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
var bibliography007SORKindOfDiscCylinderOrTape = map[string]string{
	"a": "Master tape",
	"b": "Tape duplication master",
	"d": "Disc master (negative)",
	"i": "Instantaneous (recorded on the spot)",
	"m": "Mass produced",
	"n": "Not applicable",
	"r": "Mother (positive)",
	"s": "Stamper (negative)",
	"t": "Test pressing",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORKindOfMaterial = map[string]string{
	"a": "Lacquer coating",
	"b": "Cellulose nitrate",
	"c": "Acetate tape with ferrous oxide",
	"g": "Glass with lacquer",
	"i": "Aluminum with lacquer",
	"r": "Paper with lacquer or ferrous oxide",
	"l": "Metal",
	"m": "Plastic with metal",
	"n": "Not applicable",
	"p": "Plastic",
	"s": "Shellac",
	"u": "Unknown",
	"w": "Wax",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007SORKindOfCutting = map[string]string{
	"h": "Hill-and-dale cutting",
	"l": "Lateral or combined cutting",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007SORSpecialPlaybackCharacteristics = map[string]string{
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
var bibliography007SORCaptureAndStorageTechnique = map[string]string{
	"a": "Acoustical capture, direct storage",
	"b": "Direct storage, not acoustical",
	"d": "Digital storage",
	"e": "Analog electrical storage",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007SOR parses the 007 control field data for
// Bibliography records SOUND RECORDING (SOR) data
func parseBibliography007SOR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007SORSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Speed"] = codeLookup(bibliography007SORSpeed, s, 3, 1)
	pd["(04/01) Configuration of playback channels"] = codeLookup(bibliography007SORConfigurationOfPlaybackChannels, s, 4, 1)
	pd["(05/01) Groove width/groove pitch"] = codeLookup(bibliography007SORGrooveWidthGroovePitch, s, 5, 1)
	pd["(06/01) Dimensions"] = codeLookup(bibliography007SORDimensions, s, 6, 1)
	pd["(07/01) Tape width"] = codeLookup(bibliography007SORTapeWidth, s, 7, 1)
	pd["(08/01) Tape configuration"] = codeLookup(bibliography007SORTapeConfiguration, s, 8, 1)
	pd["(09/01) Kind of disc, cylinder or tape"] = codeLookup(bibliography007SORKindOfDiscCylinderOrTape, s, 9, 1)
	pd["(10/01) Kind of material"] = codeLookup(bibliography007SORKindOfMaterial, s, 10, 1)
	pd["(11/01) Kind of cutting"] = codeLookup(bibliography007SORKindOfCutting, s, 11, 1)
	pd["(12/01) Special playback characteristics"] = codeLookup(bibliography007SORSpecialPlaybackCharacteristics, s, 12, 1)
	pd["(13/01) Capture and storage technique"] = codeLookup(bibliography007SORCaptureAndStorageTechnique, s, 13, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- TEXT
var bibliography007TXTSpecificMaterialDesignation = map[string]string{
	"a": "Regular print",
	"b": "Large print",
	"c": "Braille",
	"d": "Text in looseleaf binder",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007TXT parses the 007 control field data for
// Bibliography records TEXT (TXT) data
func parseBibliography007TXT(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007TXTSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- VIDEORECORDING
var bibliography007VIRSpecificMaterialDesignation = map[string]string{
	"c": "Videocartridge",
	"d": "Videodisc",
	"f": "Videocassette",
	"r": "Videoreel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007VIRColor = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography007VIRVideorecordingFormat = map[string]string{
	"a": "Beta (1/2 in., videocassette)",
	"b": "VHS (1/2 in., videocassette)",
	"c": "U-matic (3/4 in., videocassette)",
	"d": "EIAJ (1/2 in. reel)",
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
var bibliography007VIRSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography007VIRMediumForSound = map[string]string{
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
var bibliography007VIRDimensions = map[string]string{
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
var bibliography007VIRConfigurationOfPlaybackChannels = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007VIR parses the 007 control field data for
// Bibliography records VIDEORECORDING (VIR) data
func parseBibliography007VIR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007VIRSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(bibliography007VIRColor, s, 3, 1)
	pd["(04/01) Videorecording format"] = codeLookup(bibliography007VIRVideorecordingFormat, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(bibliography007VIRSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(bibliography007VIRMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(bibliography007VIRDimensions, s, 7, 1)
	pd["(08/01) Configuration of playback channels"] = codeLookup(bibliography007VIRConfigurationOfPlaybackChannels, s, 8, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 007 -- UNSPECIFIED
var bibliography007UNSSpecificMaterialDesignation = map[string]string{
	"m": "Multiple physical forms",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography007UNS parses the 007 control field data for
// Bibliography records UNSPECIFIED (UNS) data
func parseBibliography007UNS(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(bibliography007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(bibliography007UNSSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- ALL MATERIALS
var bibliography008TypeOfDatePublicationStatus = map[string]string{
	"b": "No dates given; B.C. date involved",
	"c": "Continuing resource currently published",
	"d": "Continuing resource ceased publication",
	"e": "Detailed date",
	"i": "Inclusive dates of collection",
	"k": "Range of years of bulk of collection",
	"m": "Multiple dates",
	"n": "Dates unknown",
	"p": "Date of distribution/release/issue and production/recording session when different",
	"q": "Questionable date",
	"r": "Reprint/reissue date and original date",
	"s": "Single known date/probable date",
	"t": "Publication date and copyright date",
	"u": "Continuing resource status unknown",
	"|": "No attempt to code",
}
var bibliography008Date1 = map[string]string{
	"1-9": "Date digit",
	" ":   "Date element is not applicable",
	"u":   "Date element is totally or partially unknown",
	"|":   "No attempt to code",
}
var bibliography008Date2 = map[string]string{
	"1-9": "Date digit",
	" ":   "Date element is not applicable",
	"u":   "Date element is totally or partially unknown",
	"|":   "No attempt to code",
}
var bibliography008ModifiedRecord = map[string]string{
	" ": "Not modified",
	"d": "Dashed-on information omitted",
	"o": "Completely romanized/printed cards romanized",
	"r": "Completely romanized/printed cards in script",
	"s": "Shortened",
	"x": "Missing characters",
	"|": "No attempt to code",
}
var bibliography008CatalogingSource = map[string]string{
	" ": "National bibliographic agency",
	"c": "Cooperative cataloging program",
	"d": "Other",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseBibliography008 parses the 008 control field data for
// Bibliography records ALL MATERIALS () data
func parseBibliography008(d *Cf008Desc, s string) {

	d.append("(00/06) Date entered on file", CodeValue{Code: pluckBytes(s, 0, 6), Label: ""})
	d.append("(06/01) Type of date/Publication status", codeLookup(bibliography008TypeOfDatePublicationStatus, s, 6, 1))

	rt07 := codeLookup(bibliography008Date1, s, 7, 1)
	if rt07.Label == "" {
		rt07 = CodeValue{Code: pluckBytes(s, 7, 4), Label: "Date"}
	}
	d.append("(07/04) Date 1", rt07)

	rt11 := codeLookup(bibliography008Date2, s, 11, 1)
	if rt11.Label == "" {
		rt11 = CodeValue{Code: pluckBytes(s, 11, 4), Label: "Date"}
	}
	d.append("(11/04) Date 2", rt11)

	d.append("(15/03) Place of publication, production, or execution", CodeValue{Code: pluckBytes(s, 15, 3), Label: ""})
	d.append("(35/03) Language", CodeValue{Code: pluckBytes(s, 35, 3), Label: ""})
	d.append("(38/01) Modified record", codeLookup(bibliography008ModifiedRecord, s, 38, 1))
	d.append("(39/01) Cataloging source", codeLookup(bibliography008CatalogingSource, s, 39, 1))
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- BOOKS
var bibliography008BKIllustrations = map[string]string{
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
var bibliography008BKTargetAudience = map[string]string{
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
var bibliography008BKFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008BKNatureOfContents = map[string]string{
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
var bibliography008BKGovernmentPublication = map[string]string{
	" ": "Not a government publication",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government publication--level undetermined",
	"s": "State, provincial, territorial, dependent, etc.",
	"u": "Unknown if item is government publication",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography008BKConferencePublication = map[string]string{
	"0": "Not a conference publication",
	"1": "Conference publication",
	"|": "No attempt to code",
}
var bibliography008BKFestschrift = map[string]string{
	"0": "Not a festschrift",
	"1": "Festschrift",
	"|": "No attempt to code",
}
var bibliography008BKIndex = map[string]string{
	"0": "No index",
	"1": "Index present",
	"|": "No attempt to code",
}
var bibliography008BKLiteraryForm = map[string]string{
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
var bibliography008BKBiography = map[string]string{
	" ": "No biographical material",
	"a": "Autobiography",
	"b": "Individual biography",
	"c": "Collective biography",
	"d": "Contains biographical information",
	"|": "No attempt to code",
}

// parseBibliography008BK parses the 008 control field data for
// Bibliography records BOOKS (BK) data
func parseBibliography008BK(d *Cf008Desc, s string) {

	for i := 0; i < 4; i++ {
		d.append("(18/04) Illustrations", codeLookup(bibliography008BKIllustrations, s, i, 1))
	}

	d.append("(22/01) Target audience", codeLookup(bibliography008BKTargetAudience, s, 4, 1))
	d.append("(23/01) Form of item", codeLookup(bibliography008BKFormOfItem, s, 5, 1))

	for i := 6; i < 10; i++ {
		d.append("(24/04) Nature of contents", codeLookup(bibliography008BKNatureOfContents, s, i, 1))
	}

	d.append("(28/01) Government publication", codeLookup(bibliography008BKGovernmentPublication, s, 10, 1))
	d.append("(29/01) Conference publication", codeLookup(bibliography008BKConferencePublication, s, 11, 1))
	d.append("(30/01) Festschrift", codeLookup(bibliography008BKFestschrift, s, 12, 1))
	d.append("(31/01) Index", codeLookup(bibliography008BKIndex, s, 13, 1))
	d.append("(32/01) Undefined", CodeValue{Code: pluckBytes(s, 14, 1), Label: ""})
	d.append("(33/01) Literary form", codeLookup(bibliography008BKLiteraryForm, s, 15, 1))
	d.append("(34/01) Biography", codeLookup(bibliography008BKBiography, s, 16, 1))
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- COMPUTER FILES
var bibliography008CFTargetAudience = map[string]string{
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
var bibliography008CFFormOfItem = map[string]string{
	"o": "Online",
	"q": "Direct electronic",
}
var bibliography008CFTypeOfComputerFile = map[string]string{
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
var bibliography008CFGovernmentPublication = map[string]string{
	" ": "Not a government publication",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government publication--level undetermined",
	"s": "State, provincial, territorial, dependent, etc.",
	"u": "Unknown if item is government publication",
	"z": "Other",
	"|": "No attempt to code",
}

// parseBibliography008CF parses the 008 control field data for
// Bibliography records COMPUTER FILES (CF) data
func parseBibliography008CF(d *Cf008Desc, s string) {

	d.append("(18/04) Undefined", CodeValue{Code: pluckBytes(s, 0, 4), Label: ""})
	d.append("(22/01) Target audience", codeLookup(bibliography008CFTargetAudience, s, 4, 1))
	d.append("(23/01) Form of item", codeLookup(bibliography008CFFormOfItem, s, 5, 1))
	d.append("(24/02) Undefined", CodeValue{Code: pluckBytes(s, 6, 2), Label: ""})
	d.append("(26/01) Type of computer file", codeLookup(bibliography008CFTypeOfComputerFile, s, 8, 1))
	d.append("(27/01) Undefined", CodeValue{Code: pluckBytes(s, 9, 1), Label: ""})
	d.append("(28/01) Government publication", codeLookup(bibliography008CFGovernmentPublication, s, 10, 1))
	d.append("(29/06) Undefined", CodeValue{Code: pluckBytes(s, 11, 6), Label: ""})
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- MAPS
var bibliography008MPRelief = map[string]string{
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
	"z": "Other relief type",
	"|": "No attempt to code",
}
var bibliography008MPProjection = map[string]string{
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
	"ca": "Alber's equal area",
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
	"|":  "No attempt to code",
}
var bibliography008MPTypeOfCartographicMaterial = map[string]string{
	"a": "Single map",
	"b": "Map series",
	"c": "Map serial",
	"d": "Globe",
	"e": "Atlas",
	"f": "Separate map supplement to another work",
	"g": "Map bound as part of another work",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography008MPGovernmentPublication = map[string]string{
	" ": "Not a government publication",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government publication--level undetermined",
	"s": "State, provincial, territorial, dependent, etc.",
	"u": "Unknown if item is government publication",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography008MPFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008MPIndex = map[string]string{
	"0": "No index",
	"1": "Index present",
	"|": "No attempt to code",
}
var bibliography008MPSpecialFormatCharacteristics = map[string]string{
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

// parseBibliography008MP parses the 008 control field data for
// Bibliography records MAPS (MP) data
func parseBibliography008MP(d *Cf008Desc, s string) {

	for i := 0; i < 4; i++ {
		d.append("(18/04) Relief", codeLookup(bibliography008MPRelief, s, i, 1))
	}

	d.append("(22/02) Projection", codeLookup(bibliography008MPProjection, s, 4, 2))
	d.append("(24/01) Undefined", CodeValue{Code: pluckBytes(s, 6, 1), Label: ""})
	d.append("(25/01) Type of cartographic material", codeLookup(bibliography008MPTypeOfCartographicMaterial, s, 7, 1))
	d.append("(26/02) Undefined", CodeValue{Code: pluckBytes(s, 8, 2), Label: ""})
	d.append("(28/01) Government publication", codeLookup(bibliography008MPGovernmentPublication, s, 10, 1))
	d.append("(29/01) Form of item", codeLookup(bibliography008MPFormOfItem, s, 11, 1))
	d.append("(30/01) Undefined", CodeValue{Code: pluckBytes(s, 12, 1), Label: ""})
	d.append("(31/01) Index", codeLookup(bibliography008MPIndex, s, 13, 1))
	d.append("(32/01) Undefined", CodeValue{Code: pluckBytes(s, 14, 1), Label: ""})

	for i := 15; i < 17; i++ {
		d.append("(33/02) Special format characteristics", codeLookup(bibliography008MPSpecialFormatCharacteristics, s, i, 1))
	}

}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- MUSIC
var bibliography008MUFormOfComposition = map[string]string{
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
	"dv": "Divertimentos, serenades, cassations, divertissements, notturni",
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
var bibliography008MUFormatOfMusic = map[string]string{
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
var bibliography008MUMusicParts = map[string]string{
	" ": "No parts in hand or not specified",
	"d": "Instrumental and vocal parts",
	"e": "Instrumental parts",
	"f": "Vocal parts",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var bibliography008MUTargetAudience = map[string]string{
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
var bibliography008MUFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"g": "Punched paper tape [   ]",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008MUAccompanyingMatter = map[string]string{
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
var bibliography008MULiteraryTextForSoundRecordings = map[string]string{
	" ": "Item is a musical sound recording",
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
var bibliography008MUTranspositionAndArrangement = map[string]string{
	" ": "Not arrangement or transposition or not specified",
	"a": "Transposition",
	"b": "Arrangement",
	"c": "Both transposed and arranged",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseBibliography008MU parses the 008 control field data for
// Bibliography records MUSIC (MU) data
func parseBibliography008MU(d *Cf008Desc, s string) {

	d.append("(18/02) Form of composition", codeLookup(bibliography008MUFormOfComposition, s, 0, 2))
	d.append("(20/01) Format of music", codeLookup(bibliography008MUFormatOfMusic, s, 2, 1))
	d.append("(21/01) Music parts", codeLookup(bibliography008MUMusicParts, s, 3, 1))
	d.append("(22/01) Target audience", codeLookup(bibliography008MUTargetAudience, s, 4, 1))
	d.append("(23/01) Form of item", codeLookup(bibliography008MUFormOfItem, s, 5, 1))

	for i := 6; i < 12; i++ {
		d.append("(24/06) Accompanying matter", codeLookup(bibliography008MUAccompanyingMatter, s, i, 1))
	}

	for i := 12; i < 14; i++ {
		d.append("(30/02) Literary text for sound recordings", codeLookup(bibliography008MULiteraryTextForSoundRecordings, s, i, 1))
	}

	d.append("(32/01) Undefined", CodeValue{Code: pluckBytes(s, 14, 1), Label: ""})
	d.append("(33/01) Transposition and arrangement", codeLookup(bibliography008MUTranspositionAndArrangement, s, 15, 1))
	d.append("(34/01) Undefined", CodeValue{Code: pluckBytes(s, 16, 1), Label: ""})
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- CONTINUING RESOURCES
var bibliography008CRFrequency = map[string]string{
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
var bibliography008CRRegularity = map[string]string{
	"n": "Normalized irregular",
	"r": "Regular",
	"u": "Unknown",
	"x": "Completely irregular",
	"|": "No attempt to code",
}
var bibliography008CRTypeOfContinuingResource = map[string]string{
	" ": "None of the following",
	"d": "Updating database",
	"l": "Updating loose-leaf",
	"m": "Monographic series",
	"n": "Newspaper",
	"p": "Periodical",
	"w": "Updating Web site",
	"|": "No attempt to code",
}
var bibliography008CRFormOfOriginalItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"e": "Newspaper format",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008CRFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008CRNatureOfEntireWork = map[string]string{
	" ": "No specified nature of entire work",
	"a": "Abstracts/summaries",
	"b": "Bibliographies",
	"c": "Catalogs",
	"d": "Dictionaries",
	"e": "Encyclopedias",
	"f": "Handbooks",
	"g": "Legal articles",
	"h": "Biography",
	"i": "Indexes",
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
	"5": "Calendars",
	"6": "Comics/graphic novels",
	"|": "No attempt to code",
}
var bibliography008CRNatureOfContents = map[string]string{
	" ": "No specified nature of contents",
	"a": "Abstracts/summaries",
	"b": "Bibliographies",
	"c": "Catalogs",
	"d": "Dictionaries",
	"e": "Encyclopedias",
	"f": "Handbooks",
	"g": "Legal articles",
	"h": "Biography",
	"i": "Indexes",
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
	"5": "Calendars",
	"6": "Comics/graphic novels",
	"|": "No attempt to code",
}
var bibliography008CRGovernmentPublication = map[string]string{
	" ": "Not a government publication",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government publication--level undetermined",
	"s": "State, provincial, territorial, dependent,etc.",
	"u": "Unknown if item is government publication",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography008CRConferencePublication = map[string]string{
	"0": "Not a conference publication",
	"1": "Conference publication",
	"|": "No attempt to code",
}
var bibliography008CROriginalAlphabetOrScriptOfTitle = map[string]string{
	" ": "No alphabet or script given/no key title",
	"a": "Basic roman",
	"b": "Extended roman",
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
var bibliography008CREntryConvention = map[string]string{
	"0": "Successive entry",
	"1": "Latest entry",
	"2": "Integrated entry",
	"|": "No attempt to code",
}

// parseBibliography008CR parses the 008 control field data for
// Bibliography records CONTINUING RESOURCES (CR) data
func parseBibliography008CR(d *Cf008Desc, s string) {

	d.append("(18/01) Frequency", codeLookup(bibliography008CRFrequency, s, 0, 1))
	d.append("(19/01) Regularity", codeLookup(bibliography008CRRegularity, s, 1, 1))
	d.append("(21/01) Type of continuing resource", codeLookup(bibliography008CRTypeOfContinuingResource, s, 3, 1))
	d.append("(22/01) Form of original item", codeLookup(bibliography008CRFormOfOriginalItem, s, 4, 1))
	d.append("(23/01) Form of item", codeLookup(bibliography008CRFormOfItem, s, 5, 1))
	d.append("(24/01) Nature of entire work", codeLookup(bibliography008CRNatureOfEntireWork, s, 6, 1))

	for i := 7; i < 10; i++ {
		d.append("(25/03) Nature of contents", codeLookup(bibliography008CRNatureOfContents, s, i, 1))
	}

	d.append("(28/01) Government publication", codeLookup(bibliography008CRGovernmentPublication, s, 10, 1))
	d.append("(29/01) Conference publication", codeLookup(bibliography008CRConferencePublication, s, 11, 1))
	d.append("(30/03) Undefined", CodeValue{Code: pluckBytes(s, 12, 3), Label: ""})
	d.append("(33/01) Original alphabet or script of title", codeLookup(bibliography008CROriginalAlphabetOrScriptOfTitle, s, 15, 1))
	d.append("(34/01) Entry convention", codeLookup(bibliography008CREntryConvention, s, 16, 1))
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- VISUAL MATERIALS
var bibliography008VMRunningTimeForMotionPicturesAndVideorecordings = map[string]string{
	"000":     "Running time exceeds three characters",
	"001-999": "Running time",
	"---":     "Running time unknown",
	"nnn":     "Not applicable",
	"|||":     "No attempt to code",
}
var bibliography008VMTargetAudience = map[string]string{
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
var bibliography008VMGovernmentPublication = map[string]string{
	" ": "Not a government publication",
	"a": "Autonomous or semi-autonomous component",
	"c": "Multilocal",
	"f": "Federal/national",
	"i": "International intergovernmental",
	"l": "Local",
	"m": "Multistate",
	"o": "Government publication--level undetermined",
	"s": "State, provincial, territorial, dependent, etc.",
	"u": "Unknown if item is government publication",
	"z": "Other",
	"|": "No attempt to code",
}
var bibliography008VMFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}
var bibliography008VMTypeOfVisualMaterial = map[string]string{
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
var bibliography008VMTechnique = map[string]string{
	"a": "Animation",
	"c": "Animation and live action",
	"l": "Live action",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other technique",
	"|": "No attempt to code",
}

// parseBibliography008VM parses the 008 control field data for
// Bibliography records VISUAL MATERIALS (VM) data
func parseBibliography008VM(d *Cf008Desc, s string) {

	rt18 := codeLookup(bibliography008VMRunningTimeForMotionPicturesAndVideorecordings, s, 0, 3)
	if rt18.Code != "" && rt18.Label == "" {
		rt18.Label = "Running time"
	}
	d.append("(18/03) Running time for motion pictures and videorecordings", rt18)

	d.append("(21/01) Undefined", CodeValue{Code: pluckBytes(s, 3, 1), Label: ""})
	d.append("(22/01) Target audience", codeLookup(bibliography008VMTargetAudience, s, 4, 1))
	d.append("(23/05) Undefined", CodeValue{Code: pluckBytes(s, 5, 5), Label: ""})
	d.append("(28/01) Government publication", codeLookup(bibliography008VMGovernmentPublication, s, 10, 1))
	d.append("(29/01) Form of item", codeLookup(bibliography008VMFormOfItem, s, 11, 1))
	d.append("(30/03) Undefined", CodeValue{Code: pluckBytes(s, 12, 3), Label: ""})
	d.append("(33/01) Type of visual material", codeLookup(bibliography008VMTypeOfVisualMaterial, s, 15, 1))
	d.append("(34/01) Technique", codeLookup(bibliography008VMTechnique, s, 16, 1))
}

////////////////////////////////////////////////////////////////////////
// Bibliography -- 008 -- MIXED MATERIALS
var bibliography008MXFormOfItem = map[string]string{
	" ": "None of the following",
	"a": "Microfilm",
	"b": "Microfiche",
	"c": "Microopaque",
	"d": "Large print",
	"f": "Braille",
	"o": "Online",
	"q": "Direct electronic",
	"r": "Regular print reproduction",
	"s": "Electronic",
	"|": "No attempt to code",
}

// parseBibliography008MX parses the 008 control field data for
// Bibliography records MIXED MATERIALS (MX) data
func parseBibliography008MX(d *Cf008Desc, s string) {

	d.append("(18/05) Undefined", CodeValue{Code: pluckBytes(s, 0, 5), Label: ""})
	d.append("(23/01) Form of item", codeLookup(bibliography008MXFormOfItem, s, 5, 1))
	d.append("(24/11) Undefined", CodeValue{Code: pluckBytes(s, 6, 11), Label: ""})
}

////////////////////////////////////////////////////////////////////////
// Classification
////////////////////////////////////////////////////////////////////////
// Classification -- 008
var classification008KindOfRecord = map[string]string{
	"a": "Schedule record",
	"b": "Table record",
	"c": "Index term record",
}
var classification008TypeOfNumber = map[string]string{
	"a": "Single number",
	"b": "Defined number span",
	"c": "Summary number span",
	"n": "Not applicable",
}
var classification008ClassificationValidity = map[string]string{
	"a": "Valid",
	"b": "First number of span invalid",
	"c": "Last number of span invalid",
	"d": "Completely invalid",
	"e": "Obsolete",
	"n": "Not applicable",
}
var classification008StandardOrOptionalDesignation = map[string]string{
	"a": "Standard",
	"b": "Optional",
	"n": "Not applicable",
}
var classification008RecordUpdateInProcess = map[string]string{
	"a": "Record can be used",
	"b": "Record is being updated",
}
var classification008LevelOfEstablishment = map[string]string{
	"a": "Fully established",
	"c": "Provisional",
}
var classification008SynthesizedNumberIndication = map[string]string{
	"a": "Not synthesized",
	"b": "Synthesized",
	"n": "Not applicable",
}
var classification008DisplayController = map[string]string{
	"a": "Displayed in standard schedules or tables",
	"b": "Extended display",
}

// parseClassification008 parses the 008 control field data for
// Classification records data
func parseClassification008(d *Cf008Desc, s string) {

	d.append("(00/06) Date entered on file", CodeValue{Code: pluckBytes(s, 0, 6), Label: ""})
	d.append("(06/01) Kind of record", codeLookup(classification008KindOfRecord, s, 6, 1))
	d.append("(07/01) Type of number", codeLookup(classification008TypeOfNumber, s, 7, 1))
	d.append("(08/01) Classification validity", codeLookup(classification008ClassificationValidity, s, 8, 1))
	d.append("(09/01) Standard or optional designation", codeLookup(classification008StandardOrOptionalDesignation, s, 9, 1))
	d.append("(10/01) Record update in process", codeLookup(classification008RecordUpdateInProcess, s, 10, 1))
	d.append("(11/01) Level of establishment", codeLookup(classification008LevelOfEstablishment, s, 11, 1))
	d.append("(12/01) Synthesized number indication", codeLookup(classification008SynthesizedNumberIndication, s, 12, 1))
	d.append("(13/01) Display controller", codeLookup(classification008DisplayController, s, 13, 1))
}

////////////////////////////////////////////////////////////////////////
// Community
////////////////////////////////////////////////////////////////////////
// Community -- 007
var community007Category = map[string]string{
	"e": "Disabled",
}
var community007StairwayRamps = map[string]string{
	"a": "No ramps",
	"b": "Entrance and internal ramps",
	"c": "Entrance ramp only--multiple floors",
	"d": "Entrance ramp only--single floor",
	"e": "Internal ramps only",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007Doors = map[string]string{
	"a": "No wide or offset-hinge doors",
	"b": "Wide or offset-hinge doors",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007FurnitureEquipmentDisplayRacks = map[string]string{
	"a": "Unaccessible furniture, equipment, display racks",
	"b": "Accessible furniture, equipment, display racks",
	"c": "No furniture, equipment, display racks",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007Restrooms = map[string]string{
	"a": "No special restroom accommodations or grab bars",
	"b": "Special restroom accommodations and grab bars",
	"c": "Special restroom accommodations only",
	"d": "Grab bars only",
	"e": "No restrooms",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007Elevators = map[string]string{
	"a": "No special elevators or control buttons",
	"b": "Special elevators and control buttons",
	"c": "Special elevators only",
	"d": "Special control buttons only",
	"e": "No elevators",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007Telephones = map[string]string{
	"a": "No lowered telephones or handset amplifiers",
	"b": "Lowered telephones and handset amplifiers",
	"c": "Lowered telephones only",
	"d": "Handset amplifiers only",
	"e": "No telephones",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007FlashingEmergencyLights = map[string]string{
	"a": "No flashing emergency lights",
	"b": "Flashing emergency lights",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007SignLanguage = map[string]string{
	"a": "No sign language",
	"b": "Sign language",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007SubtitlesAndOrSupertitles = map[string]string{
	"a": "No subtitles or supertitles",
	"b": "Subtitles and supertitles",
	"c": "Subtitles only",
	"d": "Supertitles only",
	"n": "Not applicable",
	"u": "Unknown",
}
var community007Parking = map[string]string{
	"a": "No handicapped accessible parking available",
	"b": "Handicapped accessible parking available with high clearance for special vehicles",
	"c": "Handicapped accessible parking available with low clearance only",
	"n": "Not applicable",
	"u": "Unknown",
}

// parseCommunity007 parses the 007 control field data for
// Community records data
func parseCommunity007(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category"] = codeLookup(community007Category, s, 0, 1)
	pd["(01/01) Stairway ramps"] = codeLookup(community007StairwayRamps, s, 1, 1)
	pd["(02/01) Doors"] = codeLookup(community007Doors, s, 2, 1)
	pd["(03/01) Furniture, equipment, display racks"] = codeLookup(community007FurnitureEquipmentDisplayRacks, s, 3, 1)
	pd["(04/01) Restrooms"] = codeLookup(community007Restrooms, s, 4, 1)
	pd["(05/01) Elevators"] = codeLookup(community007Elevators, s, 5, 1)
	pd["(06/01) Telephones"] = codeLookup(community007Telephones, s, 6, 1)
	pd["(07/01) Flashing emergency lights"] = codeLookup(community007FlashingEmergencyLights, s, 7, 1)
	pd["(08/01) Sign language"] = codeLookup(community007SignLanguage, s, 8, 1)
	pd["(09/01) Subtitles and/or supertitles"] = codeLookup(community007SubtitlesAndOrSupertitles, s, 9, 1)
	pd["(10/01) Parking"] = codeLookup(community007Parking, s, 10, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Community -- 008
var community008VolunteerOpportunities = map[string]string{
	"a": "No volunteer opportunities",
	"b": "Volunteer opportunities",
	"n": "Not applicable",
	"u": "Unknown",
}
var community008VolunteersProvided = map[string]string{
	"a": "No volunteers provided",
	"b": "Volunteers provided",
	"n": "Not applicable",
	"u": "Unknown",
}
var community008ChildCareArrangements = map[string]string{
	"a": "No child care arrangements",
	"b": "Child care arrangements",
	"n": "Not applicable",
	"u": "Unknown",
}
var community008SpeakersBureau = map[string]string{
	"a": "No speakers bureau",
	"b": "Speakers bureau",
	"n": "Not applicable",
	"u": "Unknown",
}
var community008MutualSupportGroups = map[string]string{
	"a": "No mutual support groups",
	"b": "Mutual support groups",
	"n": "Not applicable",
	"u": "Unknown",
}
var community008MeetingRoomsAndFacilitiesAvailable = map[string]string{
	"a": "No meeting rooms and facilities",
	"b": "Meeting rooms and facilities",
	"n": "Not applicable",
	"u": "Unknown",
}

// parseCommunity008 parses the 008 control field data for
// Community records data
func parseCommunity008(d *Cf008Desc, s string) {

	d.append("(00/06) Date entered on file", CodeValue{Code: pluckBytes(s, 0, 6), Label: ""})
	d.append("(06/01) Volunteer opportunities", codeLookup(community008VolunteerOpportunities, s, 6, 1))
	d.append("(07/01) Volunteers provided", codeLookup(community008VolunteersProvided, s, 7, 1))
	d.append("(08/01) Child care arrangements", codeLookup(community008ChildCareArrangements, s, 8, 1))
	d.append("(09/01) Speakers bureau", codeLookup(community008SpeakersBureau, s, 9, 1))
	d.append("(10/01) Mutual support groups", codeLookup(community008MutualSupportGroups, s, 10, 1))
	d.append("(11/01) Meeting rooms and facilities available", codeLookup(community008MeetingRoomsAndFacilitiesAvailable, s, 11, 1))
	d.append("(12/03) Language", CodeValue{Code: pluckBytes(s, 12, 3), Label: ""})
}

////////////////////////////////////////////////////////////////////////
// Holdings
////////////////////////////////////////////////////////////////////////
// Holdings -- 007
var holdings007CategoryOfMaterial = map[string]string{
	"a": "Map",
	"c": "Computer file",
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

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- MAP
var holdings007MAPSpecificMaterialDesignation = map[string]string{
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
var holdings007MAPColor = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}
var holdings007MAPPhysicalMedium = map[string]string{
	"a": "Paper",
	"b": "Wood",
	"c": "Stone",
	"d": "Metal",
	"e": "Synthetic",
	"f": "Skin",
	"g": "Textile",
	"i": "Plastic",
	"j": "Glass",
	"l": "Vinyl",
	"n": "Vellum",
	"p": "Plaster",
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
var holdings007MAPTypeOfReproduction = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MAPProductionReproductionDetails = map[string]string{
	"a": "Photocopy, blueline print",
	"b": "Photocopy",
	"c": "Pre-production",
	"d": "Film",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MAPPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"n": "Not applicable",
	"|": "No attempt to code",
}

// parseHoldings007MAP parses the 007 control field data for
// Holdings records MAP (MAP) data
func parseHoldings007MAP(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007MAPSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007MAPColor, s, 3, 1)
	pd["(04/01) Physical medium"] = codeLookup(holdings007MAPPhysicalMedium, s, 4, 1)
	pd["(05/01) Type of reproduction"] = codeLookup(holdings007MAPTypeOfReproduction, s, 5, 1)
	pd["(06/01) Production/reproduction details"] = codeLookup(holdings007MAPProductionReproductionDetails, s, 6, 1)
	pd["(07/01) Positive/negative aspect"] = codeLookup(holdings007MAPPositiveNegativeAspect, s, 7, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- ELECTRONIC RESOURCE
var holdings007ELRSpecificMaterialDesignation = map[string]string{
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
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007ELRColor = map[string]string{
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
var holdings007ELRDimensions = map[string]string{
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
var holdings007ELRSound = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007ELRImageBitDepth = map[string]string{
	"001-999": "Exact bit depth",
	"mmm":     "Multiple",
	"nnn":     "Not applicable",
	"---":     "Unknown",
	"|||":     "No attempt to code",
}
var holdings007ELRFileFormats = map[string]string{
	"a": "One",
	"m": "Multiple",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007ELRQualityAssuranceTargetS = map[string]string{
	"a": "Absent",
	"n": "Not applicable",
	"p": "Present",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007ELRAntecedentSource = map[string]string{
	"a": "File reproduced from original",
	"b": "File reproduced from microform",
	"c": "File reproduced from an electronic resource",
	"d": "File reproduced from an intermediate (not microform)",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007ELRLevelOfCompression = map[string]string{
	"a": "Uncompressed",
	"b": "Lossless",
	"d": "Lossy",
	"m": "Mixed",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007ELRReformattingQuality = map[string]string{
	"a": "Access",
	"n": "Not applicable",
	"p": "Preservation",
	"r": "Replacement",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseHoldings007ELR parses the 007 control field data for
// Holdings records ELECTRONIC RESOURCE (ELR) data
func parseHoldings007ELR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007ELRSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007ELRColor, s, 3, 1)
	pd["(04/01) Dimensions"] = codeLookup(holdings007ELRDimensions, s, 4, 1)
	pd["(05/01) Sound"] = codeLookup(holdings007ELRSound, s, 5, 1)

	rt06 := codeLookup(holdings007ELRImageBitDepth, s, 6, 3)
	if rt06.Code != "" && rt06.Label == "" {
		rt06.Label = "Exact bit depth"
	}
	pd["(06/03) Image bit depth"] = rt06

	pd["(09/01) File formats"] = codeLookup(holdings007ELRFileFormats, s, 9, 1)
	pd["(10/01) Quality assurance target(s)"] = codeLookup(holdings007ELRQualityAssuranceTargetS, s, 10, 1)
	pd["(11/01) Antecedent/source"] = codeLookup(holdings007ELRAntecedentSource, s, 11, 1)
	pd["(12/01) Level of compression"] = codeLookup(holdings007ELRLevelOfCompression, s, 12, 1)
	pd["(13/01) Reformatting quality"] = codeLookup(holdings007ELRReformattingQuality, s, 13, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- GLOBE
var holdings007GLBSpecificMaterialDesignation = map[string]string{
	"a": "Celestial globe",
	"b": "Planetary or lunar globe",
	"c": "Terrestrial globe",
	"e": "Earth moon globe",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007GLBColor = map[string]string{
	"a": "One color",
	"c": "Multicolored",
	"|": "No attempt to code",
}
var holdings007GLBPhysicalMedium = map[string]string{
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
var holdings007GLBTypeOfReproduction = map[string]string{
	"f": "Facsimile",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007GLB parses the 007 control field data for
// Holdings records GLOBE (GLB) data
func parseHoldings007GLB(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007GLBSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007GLBColor, s, 3, 1)
	pd["(04/01) Physical medium"] = codeLookup(holdings007GLBPhysicalMedium, s, 4, 1)
	pd["(05/01) Type of reproduction"] = codeLookup(holdings007GLBTypeOfReproduction, s, 5, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- TACTILE MATERIAL
var holdings007TAMSpecificMaterialDesignation = map[string]string{
	"a": "Moon",
	"b": "Braille",
	"c": "Combination",
	"d": "Tactile, with no writing system",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007TAMClassOfBrailleWriting = map[string]string{
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
var holdings007TAMLevelOfContraction = map[string]string{
	"a": "Uncontracted",
	"b": "Contracted",
	"m": "Combination",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007TAMBrailleMusicFormat = map[string]string{
	" ": "No specified braille music format",
	"a": "Bar over bar",
	"b": "Bar by bar",
	"c": "Line over line",
	"d": "Paragraph",
	"e": "Single line",
	"f": "Section by section",
	"g": "Line by line",
	"h": "Open score",
	"i": "Spanner short form scoring",
	"j": "Short form scoring",
	"k": "Outline",
	"l": "Vertical score",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007TAMSpecificPhysicalCharacteristics = map[string]string{
	"a": "Print/braille",
	"b": "Jumbo or enlarged braille",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007TAM parses the 007 control field data for
// Holdings records TACTILE MATERIAL (TAM) data
func parseHoldings007TAM(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007TAMSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}

	pd["(03/02) Class of braille writing - 1"] = codeLookup(holdings007TAMClassOfBrailleWriting, s, 3, 1)
	pd["(03/02) Class of braille writing - 2"] = codeLookup(holdings007TAMClassOfBrailleWriting, s, 4, 1)

	pd["(05/01) Level of contraction"] = codeLookup(holdings007TAMLevelOfContraction, s, 5, 1)

	pd["(06/03) Braille music format - 1"] = codeLookup(holdings007TAMBrailleMusicFormat, s, 6, 1)
	pd["(06/03) Braille music format - 2"] = codeLookup(holdings007TAMBrailleMusicFormat, s, 7, 1)
	pd["(06/03) Braille music format - 3"] = codeLookup(holdings007TAMBrailleMusicFormat, s, 8, 1)

	pd["(09/01) Specific physical characteristics"] = codeLookup(holdings007TAMSpecificPhysicalCharacteristics, s, 9, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- PROJECTED GRAPHIC
var holdings007PRGSpecificMaterialDesignation = map[string]string{
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
var holdings007PRGColor = map[string]string{
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
var holdings007PRGBaseOfEmulsion = map[string]string{
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
var holdings007PRGSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007PRGMediumForSound = map[string]string{
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
var holdings007PRGDimensions = map[string]string{
	"a": "Standard 8 mm.",
	"b": "Super 8 mm./single 8 mm.",
	"c": "9.5 mm.",
	"d": "16 mm.",
	"e": "28 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"j": "2x2 in. or 5x5 cm.",
	"k": "2 1/4 x 2 1/4 in. or 6x6 cm.",
	"s": "4x5 in. or 10x13 cm.",
	"t": "5x7 in. or 13x18 cm.",
	"u": "Unknown",
	"v": "8x10 in. or 21x26 cm.",
	"w": "9x9 in. or 23x23 cm.",
	"x": "10x10 in. or 26x26 cm.",
	"y": "7x7 in. or 18x18 cm.",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007PRGSecondarySupportMaterial = map[string]string{
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

// parseHoldings007PRG parses the 007 control field data for
// Holdings records PROJECTED GRAPHIC (PRG) data
func parseHoldings007PRG(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007PRGSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007PRGColor, s, 3, 1)
	pd["(04/01) Base of emulsion"] = codeLookup(holdings007PRGBaseOfEmulsion, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(holdings007PRGSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(holdings007PRGMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(holdings007PRGDimensions, s, 7, 1)
	pd["(08/01) Secondary support material"] = codeLookup(holdings007PRGSecondarySupportMaterial, s, 8, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- MICROFORM
var holdings007MICSpecificMaterialDesignation = map[string]string{
	"a": "Aperture card",
	"b": "Microfilm cartridge",
	"c": "Microfilm cassette",
	"d": "Microfilm reel",
	"e": "Microfiche",
	"f": "Microfiche cassette",
	"g": "Microopaque",
	"h": "Microfiche slip",
	"j": "Microfilm roll",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MICPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"m": "Mixed polarity",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007MICDimensions = map[string]string{
	"a": "8 mm.",
	"d": "16 mm.",
	"f": "35 mm.",
	"g": "70 mm.",
	"h": "105 mm.",
	"l": "3x5 in. or 8x13 cm.",
	"m": "4x6 in. or 11x15 cm.",
	"o": "6x9 in. or 16x23 cm.",
	"p": "3 1/4 x 7 3/8 in. or 9x19 cm.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MICReductionRatioRange = map[string]string{
	"a": "Low reduction",
	"b": "Normal reduction",
	"c": "High reduction",
	"d": "Very high reduction",
	"e": "Ultra high reduction",
	"u": "Unknown",
	"v": "Reduction rate varies",
	"|": "No attempt to code",
}
var holdings007MICColor = map[string]string{
	"b": "Black-and-white (or monochrome)",
	"c": "Multicolored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MICEmulsionOnFilm = map[string]string{
	"a": "Silver halide",
	"b": "Diazo",
	"c": "Vesicular",
	"m": "Mixed emulsion",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MICGeneration = map[string]string{
	"a": "First generation (master)",
	"b": "Printing master",
	"c": "Service copy",
	"m": "Mixed generation",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007MICBaseOfFilm = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007MIC parses the 007 control field data for
// Holdings records MICROFORM (MIC) data
func parseHoldings007MIC(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007MICSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Positive/negative aspect"] = codeLookup(holdings007MICPositiveNegativeAspect, s, 3, 1)
	pd["(04/01) Dimensions"] = codeLookup(holdings007MICDimensions, s, 4, 1)
	pd["(05/01) Reduction ratio range"] = codeLookup(holdings007MICReductionRatioRange, s, 5, 1)
	pd["(06/03) Reduction ratio"] = CodeValue{Code: pluckBytes(s, 6, 3), Label: ""}
	pd["(09/01) Color"] = codeLookup(holdings007MICColor, s, 9, 1)
	pd["(10/01) Emulsion on film"] = codeLookup(holdings007MICEmulsionOnFilm, s, 10, 1)
	pd["(11/01) Generation"] = codeLookup(holdings007MICGeneration, s, 11, 1)
	pd["(12/01) Base of film"] = codeLookup(holdings007MICBaseOfFilm, s, 12, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- NONPROJECTED GRAPHIC
var holdings007NPGSpecificMaterialDesignation = map[string]string{
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
var holdings007NPGColor = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007NPGPrimarySupportMaterial = map[string]string{
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
var holdings007NPGSecondarySupportMaterial = map[string]string{
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

// parseHoldings007NPG parses the 007 control field data for
// Holdings records NONPROJECTED GRAPHIC (NPG) data
func parseHoldings007NPG(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007NPGSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007NPGColor, s, 3, 1)
	pd["(04/01) Primary support material"] = codeLookup(holdings007NPGPrimarySupportMaterial, s, 4, 1)
	pd["(05/01) Secondary support material"] = codeLookup(holdings007NPGSecondarySupportMaterial, s, 5, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- MOTION PICTURE
var holdings007MOPSpecificMaterialDesignation = map[string]string{
	"c": "Film cartridge",
	"f": "Film cassette",
	"o": "Film roll",
	"r": "Film reel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPColor = map[string]string{
	"b": "Black-and-white",
	"c": "Multicolored",
	"h": "Hand colored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPMotionPicturePresentationFormat = map[string]string{
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
var holdings007MOPSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007MOPMediumForSound = map[string]string{
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
var holdings007MOPDimensions = map[string]string{
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
var holdings007MOPConfigurationOfPlaybackChannels = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPProductionElements = map[string]string{
	"a": "Workprint",
	"b": "Trims",
	"c": "Outtakes",
	"d": "Rushes",
	"e": "Mixing tracks",
	"f": "Title bands/intertitle rolls",
	"g": "Production rolls",
	"n": "Not applicable",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPPositiveNegativeAspect = map[string]string{
	"a": "Positive",
	"b": "Negative",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPGeneration = map[string]string{
	"d": "Duplicate",
	"e": "Master",
	"o": "Original",
	"r": "Reference print/viewing copy",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPBaseOfFilm = map[string]string{
	"a": "Safety base, undetermined",
	"c": "Safety base, acetate undetermined",
	"d": "Safety base, diacetate",
	"p": "Safety base, polyester",
	"r": "Safety base, mixed",
	"t": "Safety base, triacetate",
	"i": "Nitrate base",
	"m": "Mixed base (nitrate and safety)",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPRefinedCategoriesOfColor = map[string]string{
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
var holdings007MOPKindOfColorStockOrPrint = map[string]string{
	"a": "Imbibition dye transfer prints",
	"b": "Three layer stock",
	"c": "Three layer stock, low fade",
	"d": "Duplitized stock",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007MOPDeteriorationStage = map[string]string{
	"a": "None apparent",
	"b": "Nitrate: suspicious odor",
	"c": "Nitrate: pungent odor",
	"d": "Nitrate: brownish, discoloration, fading, dusty",
	"e": "Nitrate: sticky",
	"f": "Nitrate: frothy, bubbles, blisters",
	"g": "Nitrate: congealed",
	"h": "Nitrate: powder",
	"k": "Non-nitrate: detectable deterioration (diacetate odor)",
	"l": "Non-nitrate: advanced deterioration",
	"m": "Non-nitrate: disaster",
	"|": "No attempt to code",
}
var holdings007MOPCompleteness = map[string]string{
	"c": "Complete",
	"i": "Incomplete",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}

// parseHoldings007MOP parses the 007 control field data for
// Holdings records MOTION PICTURE (MOP) data
func parseHoldings007MOP(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007MOPSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007MOPColor, s, 3, 1)
	pd["(04/01) Motion picture presentation format"] = codeLookup(holdings007MOPMotionPicturePresentationFormat, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(holdings007MOPSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(holdings007MOPMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(holdings007MOPDimensions, s, 7, 1)
	pd["(08/01) Configuration of playback channels"] = codeLookup(holdings007MOPConfigurationOfPlaybackChannels, s, 8, 1)
	pd["(09/01) Production elements"] = codeLookup(holdings007MOPProductionElements, s, 9, 1)
	pd["(10/01) Positive/negative aspect"] = codeLookup(holdings007MOPPositiveNegativeAspect, s, 10, 1)
	pd["(11/01) Generation"] = codeLookup(holdings007MOPGeneration, s, 11, 1)
	pd["(12/01) Base of film"] = codeLookup(holdings007MOPBaseOfFilm, s, 12, 1)
	pd["(13/01) Refined categories of color"] = codeLookup(holdings007MOPRefinedCategoriesOfColor, s, 13, 1)
	pd["(14/01) Kind of color stock or print"] = codeLookup(holdings007MOPKindOfColorStockOrPrint, s, 14, 1)
	pd["(15/01) Deterioration stage"] = codeLookup(holdings007MOPDeteriorationStage, s, 15, 1)
	pd["(16/01) Completeness"] = codeLookup(holdings007MOPCompleteness, s, 16, 1)
	pd["(17/06) Film inspection date"] = CodeValue{Code: pluckBytes(s, 17, 6), Label: ""}

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- KIT
var holdings007KITSpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// parseHoldings007KIT parses the 007 control field data for
// Holdings records KIT (KIT) data
func parseHoldings007KIT(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007KITSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- NOTATED MUSIC
var holdings007NMUSpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}

// parseHoldings007NMU parses the 007 control field data for
// Holdings records NOTATED MUSIC (NMU) data
func parseHoldings007NMU(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007NMUSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- REMOTE-SENSING IMAGE
var holdings007RSISpecificMaterialDesignation = map[string]string{
	"u": "Unspecified",
	"|": "No attempt to code",
}
var holdings007RSIAltitudeOfSensor = map[string]string{
	"a": "Surface",
	"b": "Airborne",
	"c": "Spaceborne",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007RSIAttitudeOfSensor = map[string]string{
	"a": "Low oblique",
	"b": "High oblique",
	"c": "Vertical",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007RSICloudCover = map[string]string{
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
var holdings007RSIPlatformConstructionType = map[string]string{
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
var holdings007RSIPlatformUseCategory = map[string]string{
	"a": "Meteorological",
	"b": "Surface observing",
	"c": "Space observing",
	"m": "Mixed uses",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007RSISensorType = map[string]string{
	"a": "Active",
	"b": "Passive",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007RSIDataType = map[string]string{
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
	"nn": "Not applicable",
	"pa": "Sonar--water depth",
	"pb": "Sonar--bottom topography images, sidescan",
	"pc": "Sonar--bottom topography, near surface",
	"pd": "Sonar--bottom topography, near bottom",
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

// parseHoldings007RSI parses the 007 control field data for
// Holdings records REMOTE-SENSING IMAGE (RSI) data
func parseHoldings007RSI(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007RSISpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Altitude of sensor"] = codeLookup(holdings007RSIAltitudeOfSensor, s, 3, 1)
	pd["(04/01) Attitude of sensor"] = codeLookup(holdings007RSIAttitudeOfSensor, s, 4, 1)
	pd["(05/01) Cloud cover"] = codeLookup(holdings007RSICloudCover, s, 5, 1)
	pd["(06/01) Platform construction type"] = codeLookup(holdings007RSIPlatformConstructionType, s, 6, 1)
	pd["(07/01) Platform use category"] = codeLookup(holdings007RSIPlatformUseCategory, s, 7, 1)
	pd["(08/01) Sensor type"] = codeLookup(holdings007RSISensorType, s, 8, 1)
	pd["(09/02) Data type"] = codeLookup(holdings007RSIDataType, s, 9, 2)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- SOUND RECORDING
var holdings007SORSpecificMaterialDesignation = map[string]string{
	"d": "Sound disc",
	"e": "Cylinder",
	"g": "Sound cartridge",
	"i": "Sound-track film",
	"q": "Roll",
	"s": "Sound cassette",
	"t": "Sound-tape reel",
	"u": "Unspecified",
	"w": "Wire recording",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORSpeed = map[string]string{
	"a": "16 rpm",
	"b": "33 1/3 rpm",
	"c": "45 rpm",
	"d": "78 rpm",
	"e": "8 rpm",
	"f": "1.4 m. per sec.",
	"h": "120 rpm",
	"i": "160 rpm",
	"k": "15/16 ips",
	"l": "1 7/8 ips",
	"m": "3 3/4 ips",
	"o": "7 1/2 ips",
	"p": "15 ips",
	"r": "30 ips",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORConfigurationOfPlaybackChannels = map[string]string{
	"m": "Monaural",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORGrooveWidthGroovePitch = map[string]string{
	"m": "Microgroove/fine",
	"n": "Not applicable",
	"s": "Coarse/standard",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORDimensions = map[string]string{
	"a": "3 in.",
	"b": "5 in.",
	"c": "7 in.",
	"d": "10 in.",
	"e": "12 in.",
	"f": "16 in.",
	"g": "4 3/4 in. or 12 cm.",
	"j": "3 7/8 x 2 1/2 in.",
	"o": "5 1/4 x 3 7/8 in.",
	"n": "Not applicable",
	"s": "2 3/4 x 4 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORTapeWidth = map[string]string{
	"l": "1/8 in.",
	"m": "1/4 in.",
	"n": "Not applicable",
	"o": "1/2 in.",
	"p": "1 in.",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORTapeConfiguration = map[string]string{
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
var holdings007SORKindOfDiscCylinderOrTape = map[string]string{
	"a": "Master tape",
	"b": "Tape duplication master",
	"d": "Disc master (negative)",
	"i": "Instantaneous (recorded on the spot)",
	"m": "Mass produced",
	"n": "Not applicable",
	"r": "Mother (positive)",
	"s": "Stamper (negative)",
	"t": "Test pressing",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORKindOfMaterial = map[string]string{
	"a": "Lacquer coating",
	"b": "Cellulose nitrate",
	"c": "Acetate tape with ferrous oxide",
	"g": "Glass with lacquer",
	"i": "Aluminum with lacquer",
	"r": "Paper with lacquer or ferrous oxide",
	"l": "Metal",
	"m": "Plastic with metal",
	"p": "Plastic",
	"s": "Shellac",
	"u": "Unknown",
	"w": "Wax",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007SORKindOfCutting = map[string]string{
	"h": "Hill-and-dale cutting",
	"l": "Lateral or combined cutting",
	"n": "Not applicable",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007SORSpecialPlaybackCharacteristics = map[string]string{
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
var holdings007SORCaptureAndStorageTechnique = map[string]string{
	"a": "Acoustical capture, direct storage",
	"b": "Direct storage, not acoustical",
	"d": "Digital storage",
	"e": "Analog electrical storage",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007SOR parses the 007 control field data for
// Holdings records SOUND RECORDING (SOR) data
func parseHoldings007SOR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007SORSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Speed"] = codeLookup(holdings007SORSpeed, s, 3, 1)
	pd["(04/01) Configuration of playback channels"] = codeLookup(holdings007SORConfigurationOfPlaybackChannels, s, 4, 1)
	pd["(05/01) Groove width/groove pitch"] = codeLookup(holdings007SORGrooveWidthGroovePitch, s, 5, 1)
	pd["(06/01) Dimensions"] = codeLookup(holdings007SORDimensions, s, 6, 1)
	pd["(07/01) Tape width"] = codeLookup(holdings007SORTapeWidth, s, 7, 1)
	pd["(08/01) Tape configuration"] = codeLookup(holdings007SORTapeConfiguration, s, 8, 1)
	pd["(09/01) Kind of disc, cylinder or tape"] = codeLookup(holdings007SORKindOfDiscCylinderOrTape, s, 9, 1)
	pd["(10/01) Kind of material"] = codeLookup(holdings007SORKindOfMaterial, s, 10, 1)
	pd["(11/01) Kind of cutting"] = codeLookup(holdings007SORKindOfCutting, s, 11, 1)
	pd["(12/01) Special playback characteristics"] = codeLookup(holdings007SORSpecialPlaybackCharacteristics, s, 12, 1)
	pd["(13/01) Capture and storage technique"] = codeLookup(holdings007SORCaptureAndStorageTechnique, s, 13, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- TEXT
var holdings007TXTSpecificMaterialDesignation = map[string]string{
	"a": "Regular print",
	"b": "Large print",
	"c": "Braille",
	"d": "Text in looseleaf binder",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007TXT parses the 007 control field data for
// Holdings records TEXT (TXT) data
func parseHoldings007TXT(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007TXTSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- VIDEORECORDING
var holdings007VIRSpecificMaterialDesignation = map[string]string{
	"c": "Videocartridge",
	"d": "Videodisc",
	"f": "Videocassette",
	"r": "Videoreel",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007VIRColor = map[string]string{
	"a": "One color",
	"b": "Black-and-white",
	"c": "Multicolored",
	"m": "Mixed",
	"n": "Not applicable",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}
var holdings007VIRVideorecordingFormat = map[string]string{
	"a": "Beta (1/2 in., videocassette)",
	"b": "VHS (1/2 in., videocassette)",
	"c": "U-matic (3/4 in., videocassette)",
	"d": "EIAJ (1/2 in. reel)",
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
var holdings007VIRSoundOnMediumOrSeparate = map[string]string{
	" ": "No sound (silent)",
	"a": "Sound on medium",
	"b": "Sound separate from medium",
	"u": "Unknown",
	"|": "No attempt to code",
}
var holdings007VIRMediumForSound = map[string]string{
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
var holdings007VIRDimensions = map[string]string{
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
var holdings007VIRConfigurationOfPlaybackChannels = map[string]string{
	"k": "Mixed",
	"m": "Monaural",
	"n": "Not applicable",
	"q": "Quadraphonic, multichannel, or surround",
	"s": "Stereophonic",
	"u": "Unknown",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007VIR parses the 007 control field data for
// Holdings records VIDEORECORDING (VIR) data
func parseHoldings007VIR(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007VIRSpecificMaterialDesignation, s, 1, 1)
	pd["(02/01) Undefined"] = CodeValue{Code: pluckBytes(s, 2, 1), Label: ""}
	pd["(03/01) Color"] = codeLookup(holdings007VIRColor, s, 3, 1)
	pd["(04/01) Videorecording format"] = codeLookup(holdings007VIRVideorecordingFormat, s, 4, 1)
	pd["(05/01) Sound on medium or separate"] = codeLookup(holdings007VIRSoundOnMediumOrSeparate, s, 5, 1)
	pd["(06/01) Medium for sound"] = codeLookup(holdings007VIRMediumForSound, s, 6, 1)
	pd["(07/01) Dimensions"] = codeLookup(holdings007VIRDimensions, s, 7, 1)
	pd["(08/01) Configuration of playback channels"] = codeLookup(holdings007VIRConfigurationOfPlaybackChannels, s, 8, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 007 -- UNSPECIFIED
var holdings007UNSSpecificMaterialDesignation = map[string]string{
	"m": "Multiple physical forms",
	"u": "Unspecified",
	"z": "Other",
	"|": "No attempt to code",
}

// parseHoldings007UNS parses the 007 control field data for
// Holdings records UNSPECIFIED (UNS) data
func parseHoldings007UNS(s string) (pd Cf007Desc) {

	pd = make(Cf007Desc)

	pd["(00/01) Category of material"] = codeLookup(holdings007CategoryOfMaterial, s, 0, 1)
	pd["(01/01) Specific material designation"] = codeLookup(holdings007UNSSpecificMaterialDesignation, s, 1, 1)

	return pd
}

////////////////////////////////////////////////////////////////////////
// Holdings -- 008
var holdings008ReceiptOrAcquisitionStatus = map[string]string{
	"0": "Unknown",
	"1": "Other receipt or acquisition status",
	"2": "Received and complete or ceased",
	"3": "On order",
	"4": "Currently received",
	"5": "Not currently received",
}
var holdings008MethodOfAcquisition = map[string]string{
	"c": "Cooperative or consortial purchase",
	"d": "Deposit",
	"e": "Exchange",
	"f": "Free",
	"g": "Gift",
	"l": "Legal deposit",
	"m": "Membership",
	"n": "Non-library purchase",
	"p": "Purchase",
	"q": "Lease",
	"u": "Unknown",
	"z": "Other method of acquisition",
}
var holdings008ExpectedAcquisitionEndDate = map[string]string{
	"[yymm]": "Date of cancellation or last expected part",
	"uuuu":   "Intent to cancel; effective date not known",
	"####":   "No intent to cancel or not applicable",
}
var holdings008GeneralRetentionPolicy = map[string]string{
	"0": "Unknown",
	"1": "Other general retention policy",
	"2": "Retained except as replaced by updates",
	"3": "Sample issue retained",
	"4": "Retained until replaced by microform",
	"5": "Retained until replaced by cumulation, replacement volume, or revision",
	"6": "Retained for a limited period",
	"7": "Not retained",
	"8": "Permanently retained",
}
var holdings008PolicyType = map[string]string{
	"l": "Latest",
	"p": "Previous",
}
var holdings008Completeness = map[string]string{
	"0": "Other",
	"1": "Complete",
	"2": "Incomplete",
	"3": "Scattered",
	"4": "Not applicable",
}
var holdings008LendingPolicy = map[string]string{
	"a": "Will lend",
	"b": "Will not lend",
	"c": "Will lend hard copy only",
	"l": "Limited lending policy",
	"u": "Unknown",
}
var holdings008ReproductionPolicy = map[string]string{
	"a": "Will reproduce",
	"b": "Will not reproduce",
	"u": "Unknown",
}
var holdings008Language = map[string]string{
	"   ": "Blanks",
	"und": "Undetermined",
}
var holdings008SeparateOrCompositeCopyReport = map[string]string{
	"0": "Separate copy report",
	"1": "Composite copy report",
}

// parseHoldings008 parses the 008 control field data for
// Holdings records data
func parseHoldings008(d *Cf008Desc, s string) {

	d.append("(00/06) Date entered on file", CodeValue{Code: pluckBytes(s, 0, 6), Label: ""})
	d.append("(06/01) Receipt or acquisition status", codeLookup(holdings008ReceiptOrAcquisitionStatus, s, 6, 1))
	d.append("(07/01) Method of acquisition", codeLookup(holdings008MethodOfAcquisition, s, 7, 1))

	rt08 := codeLookup(holdings008ExpectedAcquisitionEndDate, s, 8, 6)
	if rt08.Code != "" && rt08.Label == "" {
		rt08.Label = "Date of cancellation or last expected part"
	}
	d.append("(08/04) Expected acquisition end date", rt08)

	d.append("(12/01) General retention policy", codeLookup(holdings008GeneralRetentionPolicy, s, 12, 1))
	d.append("(13/01) Policy Type", codeLookup(holdings008PolicyType, s, 13, 1))
	d.append("(14/01) Number of units", CodeValue{Code: pluckBytes(s, 14, 1), Label: ""})
	d.append("(15/01) Unit type", CodeValue{Code: pluckBytes(s, 15, 1), Label: ""})
	d.append("(16/01) Completeness", codeLookup(holdings008Completeness, s, 16, 1))
	d.append("(17/03) Number of copies reported", CodeValue{Code: pluckBytes(s, 17, 3), Label: ""})
	d.append("(20/01) Lending policy", codeLookup(holdings008LendingPolicy, s, 20, 1))
	d.append("(21/01) Reproduction policy", codeLookup(holdings008ReproductionPolicy, s, 21, 1))
	d.append("(22/03) Language", codeLookup(holdings008Language, s, 22, 3))
	d.append("(25/01) Separate or composite copy report", codeLookup(holdings008SeparateOrCompositeCopyReport, s, 25, 1))
	d.append("(26/06) Date of report", CodeValue{Code: pluckBytes(s, 26, 6), Label: ""})
}
