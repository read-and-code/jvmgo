package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type ClasspathEntry interface {
	ReadClass(className string) ([]byte, ClasspathEntry, error)
	ToString() string
}

func NewClasspathEntry(path string) ClasspathEntry {
	if strings.Contains(path, pathListSeparator) {
		return NewCompositeClasspathEntry(path)
	}

	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return NewZipClasspathEntry(path)
	}

	return NewDirectoryClasspathEntry(path)
}