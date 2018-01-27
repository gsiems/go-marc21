package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	//
	"github.com/gsiems/go-marc21/pkg/marc21"
)

func main() {

	var recsPerFile int
	var marcFile string
	var dir string

	flag.IntVar(&recsPerFile, "c", 1000, "The number of MARC records per output file (defaults to 1000).")
	flag.StringVar(&marcFile, "m", "", "The file that contains the MARC records.")
	flag.StringVar(&dir, "d", "mark_split", "The directory to write the output files to (defaults to mark_split).")
	flag.Parse()

	fi, err := os.Open(marcFile)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	defer func() {
		if cerr := fi.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	recCount := 0
	fOut, fileCount := nextFile(dir, 0)

	for {
		rawRec, err := marc21.NextRecord(fi)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if _, err := fOut.Write(rawRec); err != nil {
			log.Fatal(err)
		}

		recCount++
		if recCount >= recsPerFile {
			closeFile(fOut)
			fOut, fileCount = nextFile(dir, fileCount)
			recCount = 0
		}
	}

	closeFile(fOut)
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func nextFile(dir string, i int) (*os.File, int) {
	i++
	fileName := fmt.Sprintf("%s/%06d.mrc", dir, i)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0640)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	return f, i
}
