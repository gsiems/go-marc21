# go-marc21

[![GoDoc](https://godoc.org/github.com/gsiems/go-marc21/pkg/marc21?status.svg)](https://godoc.org/github.com/gsiems/go-marc21/pkg/marc21)
[![Go Report Card](https://goreportcard.com/badge/github.com/gsiems/go-marc21)](https://goreportcard.com/report/github.com/gsiems/go-marc21)

## Currently does:

 * Read MARC21 and MARCXML data.

 * Convert between MARC21 and MARCXML data.

 * Provides lookup functions for extracting information from a record
    leader.

 * Write "Pretty-print" text (compatible with perl MARC::Record->as_formatted() output)

## Things that would be nice TODO:

 * Convert MARC-8 encoding to UTF-8

 * Functions for "human-friendly" extraction of titles, authors, etc.

 * Parse control fields

 * Perform error checking on MARC records

 * Read/write MARCMaker files https://www.loc.gov/marc/makrbrkr.html

