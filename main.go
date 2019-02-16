package main

import (
	"fmt"
	"github.com/Haleluak/home-test-tiki/core"
	"github.com/Haleluak/home-test-tiki/parser"
	"github.com/Haleluak/home-test-tiki/util"
	"os"
)

func main() {
	filename := "test.csv"
	resFilename := "result.csv"
	os.Args = []string{"", filename, "-f", resFilename, "-o"}
	input, err := util.ParseArgsToInput(os.Args)

	// Init our writer used to write our TMP and final files
	w, err := util.New(input.OutputToFile, input.OutputFile, input.PageSize, input.OverwriteOutput)
	if err != nil {
		core.PrintAndExit(err)
	}

	nb := parser.Parse(input, w)
	fmt.Printf("Number of records = %d\n", nb)
}