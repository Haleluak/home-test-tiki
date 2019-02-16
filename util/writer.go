package util

import (
	"errors"
	"home-test-tiki/core"
	"home-test-tiki/file"
	"os"
	"strings"
)

type Writer struct {
	outputToFile bool
	outputFile   string
	File     *file.FileHandle
}

// New will create a new Writer
func New(outputToFile bool, outputFile string, pageSize uint, allowOutputOverwrite bool) (*Writer, error) {
	if allowOutputOverwrite == false && core.IsFileExisting(outputFile) {
		return nil, errors.New("Output file already existing. You can use `-o` to allow overwrite")
	}

	filepool := file.New(outputFile, pageSize)

	return &Writer{outputToFile, outputFile, filepool}, nil
}

// PushLine insert @line to our TMP files
func (w *Writer) PushLine(line string) error {
	// Get appropriate TMP file
	fd, err := w.File.GetCurrent()
	if err != nil {
		core.PrintAndExit(err)
	}
	// Set cursor at the end of our file
	fd.Seek(0, 2)
	// Write line
	_, err = fd.WriteString(strings.Join([]string{line, "\n"}, ""))
	if err != nil {
		core.PrintAndExit(err)
	}

	return nil
}

// UpdateLine will update an existing line from an existing TMP file
func (w *Writer) UpdateLine(fd *os.File, cursorPos int64, newValue string) error {
	err := writeFileAt(fd, newValue, cursorPos)
	if err != nil {
		core.PrintAndExit(err)
	}
	return nil
}
// UpdateDeactivateDate will update the deactivate date of an existing line in TMP File
func (w *Writer) UpdateDeactivateDate(fd *os.File, cursorPos int64, newDeactivateDate string) error {
	err := writeFileAt(fd, newDeactivateDate, cursorPos+core.StartDeactivateDate)
	if err != nil {
		core.PrintAndExit(err)
	}
	return nil
}

// UpdateActivateDate will update the activate date of an existing line in TMP File
func (w *Writer) UpdateActivateDate(fd *os.File, cursorPos int64, newActivateDate string) error {
	err := writeFileAt(fd, newActivateDate, cursorPos+core.StartActivateDate)
	if err != nil {
		core.PrintAndExit(err)
	}
	return nil
}

// writeFileAt will write @line to @pos into file @fd
func writeFileAt(fd *os.File, line string, pos int64) error {
	fd.Seek(0, 0)
	_, err := fd.WriteAt([]byte(line), pos)
	if err != nil {
		return err
	}
	return err
}
