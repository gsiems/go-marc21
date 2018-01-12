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
	defer func() {
		if cerr := fi.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for {
		rec, err := marc21.ParseNextRecord(fi)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rec)
	}
}

func showHelp() {
	fmt.Println(os.Args[0])
	fmt.Println("   Dumps a MARC file as \"pretty printed\" text.")
	fmt.Printf("    Usage: %s <MARC file to dump>\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}
