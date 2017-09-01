package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipClasspathEntry struct {
	absolutePath string
}

func NewZipClasspathEntry(path string) ZipClasspathEntry {
	absolutePath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return ZipClasspathEntry{ absolutePath }
}

func (zipClasspathEntry ZipClasspathEntry) ReadClass(className string) ([]byte, ClasspathEntry, error) {
	zipReader, err := zip.OpenReader(zipClasspathEntry.absolutePath)

	if err != nil {
		return nil, nil, err
	}

	defer zipReader.Close()

	for _, zipFile := range zipReader.File {
		if zipFile.Name != className {
			continue
		}

		file, err := zipFile.Open()

		if err != nil {
			return nil, nil, err
		}

		defer file.Close()

		data, err := ioutil.ReadAll(file)

		if err != nil {
			return nil, nil, err
		}

		return data, zipClasspathEntry, nil
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (zipClasspathEntry ZipClasspathEntry) ToString() string {
	return zipClasspathEntry.absolutePath
}