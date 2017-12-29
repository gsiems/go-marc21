// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package marc21

import (
	"encoding/xml"
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

var CollectionXMLHeader string = `<?xml version="1.0" encoding="UTF-8"?>
<collection xmlns="http://www.loc.gov/MARC21/slim>"
`
var CollectionXMLFooter string = `</collection>
`

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

	for _, r := range doc.Records {
		r.Leader, err = parseLeader([]byte(r.LeaderRaw.Text))
		if err != nil {
			return doc, err
		}
	}

	return doc, nil
}

// https://www.loc.gov/standards/marcxml/
// <collection xsi:schemaLocation="http://www.loc.gov/MARC21/slim http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd">
// looks like various samples do not mess with the <marc:TAG> and simply use <TAG>

func CollectionAsXML(c Collection) (ret string, err error) {

	ret = CollectionXMLHeader

	for _, r := range c.Records {
		rx, err := RecordAsXML(r)
		if err != nil {
			return ret, err
		}
		ret += rx
	}

	ret += CollectionXMLFooter
	return ret, nil
}

func RecordAsXML(r *Record) (ret string, err error) {

	b, err := xml.MarshalIndent(r, "    ", "    ")
	if err != nil {
		return ret, err
	}
	ret = string(b)
	return ret, nil
}
