package main

import (
	"fmt"
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

	c, err := marc21.LoadXML(marcfile)
	if err != nil {
		log.Fatal(err)
	}

	for _, rec := range c.Records {
		marc, err := rec.RecordAsMARC()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(marc))
	}
}

func showHelp() {
	fmt.Println(os.Args[0])
	fmt.Println("   Converts a MARCXML file to MARC.")
	fmt.Printf("    Usage: %s <MARCXML file to convert>\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}
