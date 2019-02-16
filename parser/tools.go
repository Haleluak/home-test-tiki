package parser

import ("strings"
		"github.com/Haleluak/home-test-tiki/core")

func getNumber(line string) string {
	return line[0:core.PhoneNumberLength]
}

// We format the line to make sure we can use it properly in our TMP files
func formatLineForTmpFile(line string) string {
	if len(line) < core.LineFullSize {
		line = strings.Join([]string{line, "....-..-.."}, "")
	}
	return line
}
