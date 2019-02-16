package parser

import (
	"bufio"
	"errors"
	"fmt"
	"home-test-tiki/core"
	"home-test-tiki/util"
	"os"
)

func Parse(params *util.Input, w *util.Writer) uint {
	fd, err := os.Open(params.InputFile)

	if err != nil {
		core.PrintAndExit(err)
	}

	defer fd.Close()

	// Make sure we skip file header
	fd.Seek(int64(len(core.CsvHeader)), 0)

	// Use scanner to read line by line
	scanner := bufio.NewScanner(fd)
	var line string
	var phoneNumber string
	var i uint
	i = 1
	for scanner.Scan() {
		// For each line
		line = scanner.Text()
		l := len(line)
		if l == 0 {
			// skip
			i++
			continue
		}
		if l != 22 && l != 32 {
			core.PrintAndExit(errors.New(fmt.Sprintf("Invalid line %d\n", i+1)))
		}
		i++
		phoneNumber = getNumber(line)
		// updateIfDuplicate will handle situation where @phoneNumber is already in our TMP files
		isDuplicate, err := updateIfDuplicateAsync(w, phoneNumber, line)
		if err != nil {
			core.PrintAndExit(err)
		}

		if isDuplicate {
			// We already manage this line
			continue
		}

		// Here, it is a new phone number. We simply insert the new line in our TMP files
		err = w.PushLine(formatLineForTmpFile(line))
		if err != nil {
			core.PrintAndExit(err)
		}

	}
	return i - 1
}
