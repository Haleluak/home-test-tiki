package util

import (
	"errors"
	"github.com/Haleluak/home-test-tiki/core"
)

type Input struct {
	InputFile       string // File name containing phone numbers details
	OutputToFile    bool   // If false, write on stdout
	OutputFile      string // (-f) File name with final results
	OverwriteOutput bool   // (-o) Should we overwrite output file or not
	PageSize        uint   // (-x) Maximum size of TMP files used to parse data
	DisplayClock    bool   // (-c) Display elapsed time to check performances
}

// ParseArgsToInput will parse args and return its value into Parameters
func ParseArgsToInput(args []string) (*Input, error) {
	if len(args) < 2 || len(args) > 8 {
		return nil, errors.New("Invalid number of arguments")
	}

	displayClock := core.IndexOf(args, "-c") >= 0
	allowOutputOverwrite := core.IndexOf(args, "-o") >= 0

	inputFile, err := getInputfile(args)
	if err != nil {
		return nil, err
	}

	outputToFile, outputFile, err := getOutputDetails(args)
	if err != nil {
		return nil, err
	}

	pageSize, err := getPageSize(args)
	if err != nil {
		return nil, err
	}

	return &Input{inputFile, outputToFile, outputFile, allowOutputOverwrite, pageSize, displayClock}, nil
}
