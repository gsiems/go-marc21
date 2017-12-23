// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"encoding/xml"
	"fmt"
	"os"
)

/*
http://www.loc.gov/standards/marcxml/
http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd
http://www.loc.gov/standards/marcxml/xml/collection.xml
*/

/*
$ zek -p < collection.xml > xml.go

type Collection struct {
	XMLName        xml.Name `xml:"collection"`
	Text           string   `xml:",chardata"`
	Marc           string   `xml:"marc,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Record         []struct {
		Text   string `xml:",chardata"`
		Leader struct {
			Text string `xml:",chardata"`
		} `xml:"leader"`
		Controlfield []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"controlfield"`
		Datafield []struct {
			Text     string `xml:",chardata"`
			Tag      string `xml:"tag,attr"`
			Ind1     string `xml:"ind1,attr"`
			Ind2     string `xml:"ind2,attr"`
			Subfield []struct {
				Text string `xml:",chardata"`
				Code string `xml:"code,attr"`
			} `xml:"subfield"`
		} `xml:"datafield"`
	} `xml:"record"`
}
*/

// TODO: indent or not to indent marshalled XML? make it optional? user specified indent chars?
// TODO: fatten-up the marshalled marc:collection tag?
// TODO: escape XML (a-la xml func EscapeText(w io.Writer, s []byte) error)

// LoadXML reads a MARCXML document
func LoadXML(filename string) (Collection, error) {

	var doc Collection

	f, err := os.Open(filename)
	if err != nil {
		return doc, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	if err := dec.Decode(&doc); err != nil {
		return doc, err
	}

	return doc, nil
}

func CollectionAsXML(c Collection) (ret string) {

	ret = fmt.Sprintln("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	//ret += fmt.Sprintf("<marc:collection xmlns:marc=%q xmlns:xsi=%q xsi:schemaLocation=\"%s\">\n", c.Marc, c.Xsi, c.SchemaLocation)
	ret += "<marc:collection>\n"

	for _, r := range c.Records {
		ret += RecordAsXML(r)
	}
	ret += "</marc:collection>\n"
	return ret
}

func RecordAsXML(r Record) (ret string) {

	ret = "\t<marc:record>\n"
	ret += fmt.Sprintf("\t\t<marc:leader>%s</marc:leader>\n", r.LeaderRaw.Text)
	for _, cf := range r.Controlfields {
		ret += fmt.Sprintf("\t\t<marc:controlfield tag=%q>%s</marc:controlfield>\n", cf.Tag, cf.Text)
	}
	for _, df := range r.Datafields {
		ret += DatafieldAsXML(df)
	}
	ret += "\t</marc:record>\n"
	return ret
}

func DatafieldAsXML(df Datafield) (ret string) {

	ret = fmt.Sprintf("\t\t<marc:datafield tag=%q ind1=%q ind2=%q>\n", df.Tag, df.Ind1, df.Ind2)
	for _, sf := range df.Subfields {
		ret += SubfieldAsXML(sf)
	}
	ret += fmt.Sprintf("\t\t</marc:datafield>\n")
	return ret
}

func SubfieldAsXML(sf Subfield) (ret string) {
	ret = fmt.Sprintf("\t\t\t<marc:subfield code=%q>%s</marc:subfield>\n", sf.Code, sf.Text)
	return ret
}
