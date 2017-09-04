package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func NewWildcardClasspathEntry(path string) CompositeClasspathEntry {
	// Remove * at the end of path
	baseDirectory := path[:len(path)-1]
	compositeClasspathEntry := []ClasspathEntry{}

	walk := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && path != baseDirectory {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarClasspathEntry := NewZipClasspathEntry(path)
			compositeClasspathEntry = append(compositeClasspathEntry, jarClasspathEntry)
		}

		return nil
	}

	filepath.Walk(baseDirectory, walk)

	return compositeClasspathEntry
}
