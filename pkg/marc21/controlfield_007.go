// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

// http://www.loc.gov/marc/bibliographic/bd007.html

// CfPhysDesc contains the values for a 007 controlfield entry
type CfPhysDesc map[string]CfValue

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

var pdImageBitDepthC = map[string]string{
	"mmm": "Multiple",
	"nnn": "Not applicable",
	"---": "Unknown",
	"|||": "No attempt to code",
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationA, b, 1)
	pd["Color"] = cfShortCode(pdColorA, b, 3)
	pd["PhysicalMedium"] = cfShortCode(pdPhysicalMediumA, b, 4)
	pd["ReproductionType"] = cfShortCode(pdReproductionTypeA, b, 5)
	pd["ProductionReproductionDetails"] = cfShortCode(pdProductionDetailsA, b, 6)
	pd["PositiveNegativeAspect"] = cfShortCode(pdPositiveNegativeAspectA, b, 7)
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationC, b, 1)
	pd["Color"] = cfShortCode(pdColorC, b, 3)
	pd["Dimensions"] = cfShortCode(pdDimensionsC, b, 4)
	pd["Sound"] = cfShortCode(pdSoundC, b, 5)

	ibd := cfWideCode(pdImageBitDepthC, b, 6, 3)
	if ibd.Code != "" && ibd.Label == "" {
		ibd.Label = "Exact bit depth"
	}
	pd["ImageBitDepth"] = ibd

	pd["FileFormats"] = cfShortCode(pdFileFormatsC, b, 9)
	pd["QualityAssuranceTarget"] = cfShortCode(pdQATargetsC, b, 10)
	pd["AntecedentSource"] = cfShortCode(pdSourceC, b, 11)
	pd["CompressionLevel"] = cfShortCode(pdCompressionC, b, 12)
	pd["ReformattingQuality"] = cfShortCode(pdReformattingQualityC, b, 13)
}

func (pd CfPhysDesc) parsePdD(b []byte) {
	// Globe (007/00=d)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Physical medium
	// 05 - Type of reproduction
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationD, b, 1)
	pd["Color"] = cfShortCode(pdColorD, b, 3)
	pd["PhysicalMedium"] = cfShortCode(pdPhysicalMediumD, b, 4)
	pd["ReproductionType"] = cfShortCode(pdReproductionTypeD, b, 5)
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

	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationF, b, 1)
	//	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationD, b, 1)
	pd["ContractionLevel"] = cfShortCode(pdContractionLevelF, b, 5)
	//	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationD, b, 1)
	pd["SpecialPhysicalCharacteristics"] = cfShortCode(pdSpecialCharacteristicsF, b, 9)
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationG, b, 1)
	pd["Color"] = cfShortCode(pdColorG, b, 3)
	pd["EmulsionBase"] = cfShortCode(pdEmulsionBaseG, b, 4)
	pd["SoundLocation"] = cfShortCode(pdSoundLocationG, b, 5)
	pd["SoundMedium"] = cfShortCode(pdSoundMediumG, b, 6)
	pd["Dimensions"] = cfShortCode(pdDimensionsG, b, 7)
	pd["SecondarySupportMaterial"] = cfShortCode(pdSecondarySupportMaterialG, b, 8)
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationH, b, 1)
	pd["PositiveNegativeAspect"] = cfShortCode(pdPositiveNegativeAspectH, b, 3)
	pd["Dimensions"] = cfShortCode(pdDimensionsH, b, 4)
	pd["ReductionRatioRange"] = cfShortCode(pdReductionRatioRangeH, b, 5)
	pd["ReductionRatio"] = CfValue{Code: string(b[6:9]), Label: "Reduction ratio"}
	pd["Color"] = cfShortCode(pdColorH, b, 9)
	pd["FilmEmulsion"] = cfShortCode(pdFilmEmulsionH, b, 10)
	pd["Generation"] = cfShortCode(pdGenerationH, b, 11)
	pd["FilmBase"] = cfShortCode(pdFilmBaseH, b, 12)
}

func (pd CfPhysDesc) parsePdK(b []byte) {
	// Nonprojected graphic (007/00=k)
	// 00 - Category of material
	// 01 - Specific material designation
	// 02 - Undefined
	// 03 - Color
	// 04 - Primary support material
	// 05 - Secondary support material
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationK, b, 1)
	pd["Color"] = cfShortCode(pdColorK, b, 3)
	pd["PrimarySupportMaterial"] = cfShortCode(pdPrimarySupportMaterialK, b, 4)
	pd["SecondarySupportMaterial"] = cfShortCode(pdSecondarySupportMaterialK, b, 5)
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationM, b, 1)
	pd["Color"] = cfShortCode(pdColorM, b, 3)
	pd["PresentationFormat"] = cfShortCode(pdPresentationFormatM, b, 4)
	pd["SoundLocation"] = cfShortCode(pdSoundLocationM, b, 5)
	pd["SoundMedium"] = cfShortCode(pdSoundMediumM, b, 6)
	pd["Dimensions"] = cfShortCode(pdDimensionsM, b, 7)
	pd["PlaybackChannelConfig"] = cfShortCode(pdPlaybackChannelConfigM, b, 8)
	pd["ProductionElements"] = cfShortCode(pdProductionElementsM, b, 9)
	pd["PositiveNegativeAspect"] = cfShortCode(pdPositiveNegativeAspectM, b, 10)
	pd["Generation"] = cfShortCode(pdGenerationM, b, 11)
	pd["FilmBase"] = cfShortCode(pdFilmBaseM, b, 12)
	pd["RefinedColorCategories"] = cfShortCode(pdColorCategoriesM, b, 13)
	pd["ColorType"] = cfShortCode(pdColorTypeM, b, 14)
	pd["DeteriorationStage"] = cfShortCode(pdDeteriorationStageM, b, 15)
	pd["Completeness"] = cfShortCode(pdCompletenessM, b, 16)

	pd["FilmInspectionDate"] = CfValue{Code: string(b[17:]), Label: "Film inspection date"}
}

func (pd CfPhysDesc) parsePdO(b []byte) {
	// Kit (007/00=o)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationO, b, 1)
}

func (pd CfPhysDesc) parsePdQ(b []byte) {
	// Notated music (007/00=q)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationQ, b, 1)
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

	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationR, b, 1)
	pd["SensorAltitude"] = cfShortCode(pdSensorAltitudeR, b, 3)
	pd["SensorAttitudeR"] = cfShortCode(pdSensorAttitudeR, b, 4)
	pd["CloudCover"] = cfShortCode(pdCloudCoverR, b, 5)
	pd["PlatformConstructionType"] = cfShortCode(pdPlatformConstructionTypeR, b, 6)
	pd["PlatformUseCategory"] = cfShortCode(pdPlatformUseCategoryR, b, 7)
	pd["SensorType"] = cfShortCode(pdSensorTypeR, b, 8)
	pd["DataType"] = cfWideCode(pdDataTypeR, b, 9, 2)

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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationS, b, 1)
	pd["Speed"] = cfShortCode(pdSpeedS, b, 3)
	pd["PlaybackChannelConfig"] = cfShortCode(pdPlaybackConfigS, b, 4)
	pd["GrooveWidthPitch"] = cfShortCode(pdGrooveS, b, 5)
	pd["Dimensions"] = cfShortCode(pdDimensionsS, b, 6)
	pd["TapeWidth"] = cfShortCode(pdTapeWidthS, b, 7)
	pd["TapeConfiguration"] = cfShortCode(pdTapeConfigS, b, 8)
	pd["DiscCylinderOrTapeType"] = cfShortCode(pdItemType, b, 9)
	pd["MaterialType"] = cfShortCode(pdMaterialTypeS, b, 10)
	pd["CuttingType"] = cfShortCode(pdCuttingTypeS, b, 11)
	pd["SpecialPlaybackCharacteristics"] = cfShortCode(pdSpecialPlaybackS, b, 12)
	pd["CaptureAndStorageTechnique"] = cfShortCode(pdCaptureStorageS, b, 13)
}

func (pd CfPhysDesc) parsePdT(b []byte) {
	// Text (007/00=t)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationT, b, 1)
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
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationV, b, 1)
	pd["Color"] = cfShortCode(pdColorV, b, 3)
	pd["VideorecordingFormat"] = cfShortCode(pdFormatV, b, 4)
	pd["SoundLocation"] = cfShortCode(pdSoundLocationV, b, 5)
	pd["SoundMedium"] = cfShortCode(pdSoundMediumV, b, 6)
	pd["Dimensions"] = cfShortCode(pdDimensionsV, b, 7)
	pd["PlaybackChannelConfig"] = cfShortCode(pdPlaybackChannelsV, b, 8)
}

func (pd CfPhysDesc) parsePdZ(b []byte) {
	// Unspecified (007/00=z)
	// 00 - Category of material
	// 01 - Specific material designation
	pd["SpecificMaterialDesignation"] = cfShortCode(pdSpecificDesignationZ, b, 1)
}
