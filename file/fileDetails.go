package file

import "os"

type fileDetails struct {
	name string
	fd   *os.File
}
