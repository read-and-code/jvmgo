package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirectoryClasspathEntry struct {
	absoluteDirectory string
}

func NewDirectoryClasspathEntry(path string) DirectoryClasspathEntry {
	absoluteDirectory, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return DirectoryClasspathEntry{absoluteDirectory}
}

func (directoryClasspathEntry DirectoryClasspathEntry) ReadClass(className string) ([]byte, ClasspathEntry, error) {
	fileName := filepath.Join(directoryClasspathEntry.absoluteDirectory, className)
	data, err := ioutil.ReadFile(fileName)

	return data, directoryClasspathEntry, err
}

func (directoryClasspathEntry DirectoryClasspathEntry) ToString() string {
	return directoryClasspathEntry.absoluteDirectory
}
