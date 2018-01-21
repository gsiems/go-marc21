package marc21

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestXML(t *testing.T) {

	xmlFile := os.Getenv("TEST_MARCXML_FILE")

	doc, err := LoadXML(xmlFile)

	if err != nil {
		t.Errorf("LoadXML() failed: %q", err)
	}

	out, err := doc.AsXML()
	if err != nil {
		t.Errorf("Collection.AsXML() failed: %q", err)
	} else if out == "" {
		t.Errorf("Collection.AsXML() failed")
	}

}

func TestMARC(t *testing.T) {

	marcFile := os.Getenv("TEST_MARC_FILE")

	fi, err := os.Open(marcFile)
	if err != nil {
		t.Errorf("os.Open() failed: %q", err)
	}
	defer fi.Close()

	for {
		rec, err := ParseNextRecord(fi)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Errorf("ParseNextRecord() failed: %q", err)
			continue
		}

		code, _ := rec.RecordStatus()
		if code == "" {
			t.Errorf("RecordStatus() failed")
		}
		code, _ = rec.RecordType()
		if code == "" {
			t.Errorf("RecordType() failed")
		}
		code, _ = rec.BibliographicLevel()
		if code == "" {
			t.Errorf("BibliographicLevel() failed")
		}
		code, _ = rec.ControlType()
		if code == "" {
			t.Errorf("ControlType() failed")
		}
		code, _ = rec.CharacterCodingScheme()
		if code == "" {
			t.Errorf("CharacterCodingScheme() failed")
		}
		code, _ = rec.EncodingLevel()
		if code == "" {
			t.Errorf("EncodingLevel() failed")
		}
		code, _ = rec.CatalogingForm()
		if code == "" {
			t.Errorf("CatalogingForm() failed")
		}
		code, _ = rec.MultipartResourceRecordLevel()
		if code == "" {
			t.Errorf("MultipartResourceRecordLevel() failed")
		}

		cfd := rec.ParseControlfields()
		if false {
			fmt.Printf("%q\n", cfd)
			//t.Errorf("Record.parseControlfields() failed")
		}

		out := fmt.Sprint(rec)
		if out == "" {
			t.Errorf("Record.Print() failed")
		}

		out = fmt.Sprint(rec.Leader)
		if out == "" {
			t.Errorf("Leader.Print() failed")
		}
	}
}

func TestMARC2XML(t *testing.T) {

	marcFile := os.Getenv("TEST_MARC_FILE")

	fi, err := os.Open(marcFile)
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

	out, err := doc.AsXML()
	if err != nil {
		t.Errorf("Collection.AsXML() failed: %q", err)
	} else if out == "" {
		t.Errorf("Collection.AsXML() failed")
	}
}

func TestXML2PrettyPrint(t *testing.T) {

	xmlFile := os.Getenv("TEST_MARCXML_FILE")

	doc, err := LoadXML(xmlFile)
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

func TestRecordAsMARC(t *testing.T) {

	marcFile := os.Getenv("TEST_MARC_FILE")

	fi, err := os.Open(marcFile)
	if err != nil {
		t.Errorf("os.Open() failed: %q", err)
	}
	defer fi.Close()

	rec, err := ParseNextRecord(fi)
	if err != nil {
		t.Errorf("ParseNextRecord() failed: %q", err)
	} else {
		_, err := rec.RecordAsMARC()
		if err != nil {
			t.Errorf("RecordAsMARC() failed: %q", err)
		}
	}
}
