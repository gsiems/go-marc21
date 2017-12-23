#!/bin/sh

export TEST_MARCXML_FILE=../../git_ignore/collection.xml
export TEST_MARC_FILE=../../git_ignore/malc-20171213.mrc

[ -f coverage.out ] && rm coverage.out

echo ""
echo "### gocyclo:"
gocyclo *.go

echo ""
echo "### test:"
go test -coverprofile=coverage.out

echo ""
echo "### coverage:"
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o=coverage.html
