package core

import "os"

// IsFileExisting return true if the file exists
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
