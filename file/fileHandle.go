package file

import (
	"fmt"
	"os"
	"strings"
)

type FileHandle struct {
	seedName string // Used to create tmp filename
	pageSize uint   // Maximum size for TMP file
	files    []fileDetails
}

// New create and init a new FilePool
func New(seedName string, pageSize uint) *FileHandle {
	fp := FileHandle{seedName, pageSize, []fileDetails{}}
	return &fp
}

// ConcatFilenames will concat file names currently used into a single string
func (fp *FileHandle) ConcatFilenames() string {
	a := make([]string, len(fp.files))
	for i, fd := range fp.files {
		a[i] = fd.name
	}
	return strings.Join(a, " ")
}

// GetFilenames return names of all TMP files
func (fp *FileHandle) GetFilenames() []string {
	a := make([]string, len(fp.files))
	for i, file := range fp.files {
		a[i] = file.name
	}
	return a
}

// GetFdList return list of active file descriptors
func (fp *FileHandle) GetFdList() []*os.File {
	a := make([]*os.File, len(fp.files))
	for i, file := range fp.files {
		a[i] = file.fd
	}
	return a
}

// CleanPool will remove tmp files
func (fp *FileHandle) CleanPool() {
	for _, file := range fp.files {
		file.fd.Close()
		os.Remove(file.name)
	}
}

// GetCurrent will return the most available file to write in
func (fp *FileHandle) GetCurrent() (*os.File, error) {
	if len(fp.files) == 0 {
		// We have no file in the pool. Create new one
		err := fp.create()
		if err != nil {
			return nil, err
		}
		return fp.files[0].fd, nil
	}
	last := fp.files[len(fp.files)-1]
	fileInfo, err := last.fd.Stat()
	if err != nil {
		return nil, err
	}
	if uint(fileInfo.Size()) > fp.pageSize {
		// All our files are full. Create new one
		err := fp.create()
		if err != nil {
			return nil, err
		}
	}

	return fp.files[len(fp.files)-1].fd, nil
}

// create will create a new TMP file and insert its details in the pool
func (fp *FileHandle) create() error {
	filename := getTmpName(fp.seedName)
	_, err := os.Create(filename)
	if err != nil {
		return err
	}
	fd, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	fp.files = append(fp.files, fileDetails{filename, fd})
	return nil
}

// getTmpName generate unique name for TMP files based on @seed
func getTmpName(seed string) string {
	name := fmt.Sprintf("%s", seed)
	return name
}