package marc21

import (
	"fmt"
	"os"
	"testing"
)

func TestXML(t *testing.T) {

	xml_file := os.Getenv("TEST_MARCXML_FILE")

	doc, err := LoadXML(xml_file)

	if err != nil {
		t.Errorf("LoadXML() failed: %q", err)
	}

	out := CollectionAsXML(doc)
	if out == "" {
		t.Errorf("CollectionAsXML() failed")
	}

}

func TestMARC(t *testing.T) {

	marc_file := os.Getenv("TEST_MARC_FILE")

	fi, err := os.Open(marc_file)
	if err != nil {
		t.Errorf("os.Open() failed: %q", err)
	}
	defer fi.Close()

	rec, err := ParseNextRecord(fi)
	if err != nil {
		t.Errorf("ParseNextRecord() failed: %q", err)
	} else {
		out := fmt.Sprint(rec)
		if out == "" {
			t.Errorf("Print() failed")
		}
	}
}

func TestMARC2XML(t *testing.T) {

	marc_file := os.Getenv("TEST_MARC_FILE")

	fi, err := os.Open(marc_file)
	if err != nil {
		t.Errorf("os.Open() failed: %q", err)
	}
	defer fi.Close()

	rec, err := ParseNextRecord(fi)
	if err != nil {
		t.Errorf("ParseNextRecord() failed: %q", err)
	}

	var doc Collection
	doc.Records = append(doc.Records, rec)

	out := CollectionAsXML(doc)
	if out == "" {
		t.Errorf("CollectionAsXML() failed")
	}
}

func TestXML2PrettyPrint(t *testing.T) {

	xml_file := os.Getenv("TEST_MARCXML_FILE")

	doc, err := LoadXML(xml_file)

	if err != nil {
		t.Errorf("LoadXML() failed: %q", err)
	}

	for _, rec := range doc.Records {
		out := fmt.Sprint(rec)
		if out == "" {
			t.Errorf("Print() failed")
		}
	}
}
