#!/bin/sh

echo "Building marcdump"
go build marcdump.go

echo "Building xml2marc"
go build xml2marc.go

echo "Building marc2xml"
go build marc2xml.go

echo ""
echo "Testing marcdump"
time ./marcdump ../git_ignore/malc-20171213.mrc > marcdump.out

echo ""
echo "Testing xml2marc"
time ./xml2marc ../git_ignore/collection.xml > xml2marc.out

echo ""
echo "Testing marc2xml"
time ./marc2xml ../git_ignore/malc-20171213.mrc > marc2xml.out

# marksplit

