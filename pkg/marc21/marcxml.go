// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"encoding/xml"
	"fmt"
	"html"
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

// CollectionXMLHeader being the header (XML declaration and opening
// collection tag) of a MARCXML document
const CollectionXMLHeader = `<?xml version="1.0" encoding="UTF-8"?>
<marc:collection xmlns:marc="http://www.loc.gov/MARC21/slim"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://www.loc.gov/MARC21/slim http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd">
`

// CollectionXMLFooter being the footer (closing collection tag) of a
// MARCXML document
const CollectionXMLFooter = `</marc:collection>
`

// LoadXML reads a MARCXML document
func LoadXML(filename string) (Collection, error) {

	var doc Collection

	f, err := os.Open(filename)
	if err != nil {
		return doc, err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	dec := xml.NewDecoder(f)
	err = dec.Decode(&doc)
	return doc, err
}

// https://www.loc.gov/standards/marcxml/
// <collection xsi:schemaLocation="http://www.loc.gov/MARC21/slim http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd">
// looks like various samples do not mess with the <marc:TAG> and simply use <TAG>

// AsXML converts an entire collection to XML
func (c Collection) AsXML() (ret string, err error) {

	ret = CollectionXMLHeader

	for _, rec := range c.Records {
		rx, err := rec.AsXML()
		if err != nil {
			return "", err
		}
		ret += rx
	}
	ret += CollectionXMLFooter
	return ret, nil
}

// AsXML converts record to XML
func (rec Record) AsXML() (ret string, err error) {

	ret = "\t<marc:record>\n"
	ret += fmt.Sprintf("\t\t<marc:leader>%s</marc:leader>\n", html.EscapeString(rec.Leader.GetText()))
	for _, cf := range rec.Controlfields {
		ret += fmt.Sprintf("\t\t<marc:controlfield tag=%q>%s</marc:controlfield>\n", cf.GetTag(), html.EscapeString(cf.GetText()))
	}
	for _, df := range rec.Datafields {
		ret += fmt.Sprintf("\t\t<marc:datafield tag=%q ind1=%q ind2=%q>\n", df.GetTag(), df.GetInd1(), df.GetInd2())
		for _, sf := range df.Subfields {
			ret += fmt.Sprintf("\t\t\t<marc:subfield code=%q>%s</marc:subfield>\n", sf.GetCode(), html.EscapeString(sf.GetText()))
		}
		ret += fmt.Sprintf("\t\t</marc:datafield>\n")
	}
	ret += "\t</marc:record>\n"
	return ret, nil
}
