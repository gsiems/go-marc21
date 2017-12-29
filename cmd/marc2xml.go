package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gsiems/go-marc21/pkg/marc21"
)

func main() {

	var marcfile string
	if len(os.Args) > 1 {
		marcfile = os.Args[1]
	}

	if marcfile == "" {
		showHelp()
	}

	fi, err := os.Open(marcfile)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	defer fi.Close()

	fmt.Print(marc21.CollectionXMLHeader)

	for {
		rec, err := marc21.ParseNextRecord(fi)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		recxml, err := marc21.RecordAsXML(rec)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(recxml)
	}

	fmt.Print(marc21.CollectionXMLFooter)
}

func showHelp() {
	fmt.Println(os.Args[0])
	fmt.Println("   Converts a MARC file to MARCXML.")
	fmt.Printf("    Usage: %s <MARC file to convert>\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}
