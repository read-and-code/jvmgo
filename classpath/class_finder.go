package classpath

import (
	"os"
	"path/filepath"
)

type ClassFinder struct {
	bootstrapClasspathEntry  ClasspathEntry
	extensionsClasspathEntry ClasspathEntry
	userClasspathEntry       ClasspathEntry
}

func Parse(jrePath, classpath string) *ClassFinder {
	classFinder := &ClassFinder{}
	classFinder.parseBootstrapAndExtensionsClasspathEntry(jrePath)
	classFinder.parseUserClasspathEntry(classpath)

	return classFinder
}

func (classFinder *ClassFinder) parseBootstrapAndExtensionsClasspathEntry(jrePath string) {
	jreDirectory := getJreDirectory(jrePath)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDirectory, "lib", "*")
	classFinder.bootstrapClasspathEntry = NewWildcardClasspathEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtensionsPath := filepath.Join(jreDirectory, "lib", "ext", "*")
	classFinder.extensionsClasspathEntry = NewWildcardClasspathEntry(jreExtensionsPath)
}

func (classFinder *ClassFinder) parseUserClasspathEntry(classpath string) {
	if classpath == "" {
		classpath = "."
	}

	classFinder.userClasspathEntry = NewClasspathEntry(classpath)
}

func getJreDirectory(jrePath string) string {
	if jrePath != "" && isDirectoryExists(jrePath) {
		return jrePath
	}

	if isDirectoryExists("./jre") {
		return ".jre"
	}

	javaHome := os.Getenv("JAVA_HOME")

	if javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}

	panic("Cannot find jre folder!")
}

func isDirectoryExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func (classFinder *ClassFinder) ReadClass(className string) ([]byte, ClasspathEntry, error) {
	className = className + ".class"

	data, classpathEntry, err := classFinder.bootstrapClasspathEntry.ReadClass(className)

	if err == nil {
		return data, classpathEntry, err
	}

	data, classpathEntry, err = classFinder.extensionsClasspathEntry.ReadClass(className)

	if err == nil {
		return data, classpathEntry, err
	}

	return classFinder.userClasspathEntry.ReadClass(className)
}
