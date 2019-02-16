package core

import "os"

const (
	// OK = all good
	OK = iota
	// WrongInput = data provided by user is wrong
	WrongInput = iota
	// Error = unexpected error
	Error = iota
)

// PrintError will write the error message in stderr
func PrintError(err error) {
	os.Stderr.Write([]byte("ERROR: "))
	os.Stderr.Write([]byte(err.Error()))
	os.Stderr.Write([]byte("\n"))
}

// PrintAndExit will display a friendly message to user before exit the program with error code
func PrintAndExit(err error) {
	PrintError(err)
	os.Exit(Error)
}
