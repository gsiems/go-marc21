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

		// TODO: at some point this needs to move to specific test case
		// MARC records with specific expected outcomes. "Good" records
		// should have valid values while intentionally "bad" records
		// should have invalid or no values and should be evaluated
		// accordingly.
		rf := rec.RecordFormat()
		if rf != FmtUnknown {
			code, _ := rec.RecordStatus()
			if code == "" {
				t.Errorf("RecordStatus() failed")
			}
			code, _ = rec.RecordType()
			if code == "" {
				t.Errorf("RecordType() failed")
			}
			code, label := rec.CharacterCodingScheme()
			if code == "" {
				t.Errorf("CharacterCodingScheme() failed")
			}

			nc, nl := rec.LookupLeaderField("CharacterCodingScheme")
			if code != nc || label != nl {
				t.Errorf("LookupLeaderField(CharacterCodingScheme) failed")
			}

			s := rec.ValidLeaderFields()
			if len(s) == 0 {
				t.Errorf("ValidLeaderFields() failed")
			}

		}
		/*
			if rf != Bibliography {
				code, _ := rec.BibliographicLevel()
				if code == "" {
					t.Errorf("BibliographicLevel() failed")
				}
				code, _ = rec.ControlType()
				if code == "" {
					t.Errorf("ControlType() failed")
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
			}
		*/
		//cfd := rec.ParseControlfields()
		//if false {
		//	fmt.Printf("%q\n", cfd)
		//	//t.Errorf("Record.parseControlfields() failed")
		//}

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

func TestLeaderPrint(t *testing.T) {

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
		out := fmt.Sprint(rec.Leader)
		// TODO: this really needs to test against a specific MARC record
		// and compare the output to the specific expected value.
		if out == "" {
			t.Errorf("Sprint(rec.Leader) failed")
		}
	}
}
