package classpath

import (
	"errors"
	"strings"
)

type CompositeClasspathEntry []ClasspathEntry

func NewCompositeClasspathEntry(pathList string) CompositeClasspathEntry {
	compositeClasspathEntry := []ClasspathEntry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := NewClasspathEntry(path)
		compositeClasspathEntry = append(compositeClasspathEntry, entry)
	}

	return compositeClasspathEntry
}

func (compositeClasspathEntry CompositeClasspathEntry) ReadClass(className string) ([]byte, ClasspathEntry, error) {
	for _, classpathEntry := range compositeClasspathEntry {
		data, targetClasspathEntry, err := classpathEntry.ReadClass(className)

		if err == nil {
			return data, targetClasspathEntry, nil
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (compositeClasspathEntry CompositeClasspathEntry) ToString() string {
	value := make([]string, len(compositeClasspathEntry))

	for i, classpathEntry := range compositeClasspathEntry {
		value[i] = classpathEntry.ToString()
	}

	return strings.Join(value, pathListSeparator)
}