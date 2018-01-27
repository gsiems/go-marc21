#!/bin/sh

echo "Building marcdump"
go build marcdump.go

echo "Building xml2marc"
go build xml2marc.go

echo "Building marc2xml"
go build marc2xml.go

echo "Building marcsplit"
go build marcsplit.go

echo ""
echo "Testing marcdump"
time ./marcdump ../git_ignore/malc-20180115.mrc > marcdump.out

echo ""
echo "Testing xml2marc"
time ./xml2marc ../git_ignore/collection.xml > xml2marc.out

echo ""
echo "Testing marc2xml"
time ./marc2xml ../git_ignore/malc-20180112.mrc > marc2xml.out

echo ""
echo "Testing marcsplit"
[ -d split_out ] && rm split_out/*.mrc
[ -d split_out ] || mkdir split_out
time ./marcsplit -m ../git_ignore/malc-20180112.mrc -d split_out

