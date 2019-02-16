package util

import (
	"errors"
	"github.com/Haleluak/home-test-tiki/core"
	"regexp"
	"strconv"
)

// getInputfile parse args to get Input Filename
func getInputfile(args []string) (string, error) {
	var filename = args[1]

	match, regexpErr := regexp.MatchString(core.RegexpFilename, filename)

	if regexpErr != nil {
		return "", regexpErr
	}

	if match == false {
		return "", errors.New("Not a CSV file")
	}

	return filename, nil
}

// getOutputDetails will parse args to check if we need to write result in a csv file
// If we do it will also return the name of this file
func getOutputDetails(args []string) (bool, string, error) {
	outputArgIndex := core.IndexOf(args, "-f")

	if outputArgIndex > 0 {
		if outputArgIndex+1 >= len(args) {
			return true, "", errors.New("Expected value after '-f'")
		}
		filename := args[outputArgIndex+1]
		return true, filename, nil
	}

	return false, "", nil
}

// getPageSize will parse args to get size of page defined by user or return default value
func getPageSize(args []string) (uint, error) {
	pageSizeArgIndex := core.IndexOf(args, "-x")
	if pageSizeArgIndex > 0 {
		if pageSizeArgIndex+1 >= len(args) {
			return 0, errors.New("Expected value after '-x'")
		}
		pageSizeStr := args[pageSizeArgIndex+1]

		pageSize, err := strconv.ParseUint(pageSizeStr, 10, 32)
		if err != nil {
			return 0, err
		}

		if pageSize < 50 {
			pageSize = 50
		}
		return uint(pageSize), nil
	}

	return core.DefaultPageSize, nil
}