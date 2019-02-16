package parser

import (
	"github.com/Haleluak/home-test-tiki/core"
	"github.com/Haleluak/home-test-tiki/util"
	"os"
	"sync"
)

func updateIfDuplicateAsync(w *util.Writer, phoneNumber string, newLine string) (bool, error) {
	fds := w.File.GetFdList()

	duplicateChan := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(len(fds))
	for _, fd := range fds {
		// Process job
		go parseFileAndUpdate(w, fd, phoneNumber, newLine, wg, duplicateChan)
	}

	// Wait for all routines to be done
	go waitResult(wg, duplicateChan)

	// Parse routines result and return true if there is a duplicate
	return readResult(duplicateChan), nil
}

// parseFileAndUpdate will parse a single file and update it if required
// return true if the file contains a duplicate of @phoneNumber
func parseFileAndUpdate(w *util.Writer, fd *os.File, phoneNumber string, newLine string, wg *sync.WaitGroup, duplicateChan chan bool) {
	defer wg.Done()

	line, lineIndex, err := indexOf(fd, phoneNumber)
	if err != nil {
		core.PrintAndExit(err)
	}
	if lineIndex >= 0 {
		// The phone number is existing in our TMP files
		err = updateIfLastAction(w, fd, lineIndex, line, newLine)
		if err != nil {
			core.PrintAndExit(err)
		}
		duplicateChan <- true
	} else {
		duplicateChan <- false
	}
}

// waitResult wait for all goroutines to be done and close the channel when done
func waitResult(wg *sync.WaitGroup, duplicateChan chan bool) {
	wg.Wait()
	close(duplicateChan)
}

// readResult parse channel and return true if we have a positive result
func readResult(duplicateChan chan bool) bool {
	for isDup := range duplicateChan {
		if isDup {
			return true
		}
	}
	return false
}
