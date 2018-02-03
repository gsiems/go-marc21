package main

import (
	"fmt"
	//
	codegen "github.com/gsiems/go-marc21/codegen/pkg"
)

func main() {

	htmlfiles := []string{
		"input/ecadlist.html",
		"input/ecbdlist.html",
		"input/eccdlist.html",
		"input/eccilist.html",
		"input/echdlist.html",
	}

	for _, file := range htmlfiles {
		fmt.Printf("\n\n////////// %q\n", file)
		cfeList := codegen.ExtractCfStruct(file)

		for _, v := range cfeList.Tags {
			fmt.Print(v)
		}
	}
}
