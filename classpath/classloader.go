package classpath

import (
	"os"
	"path/filepath"
)

type Classloader struct {
	BootstrapClasspathEntry ClasspathEntry
	ExtensionsClasspathEntry ClasspathEntry
	UserClasspathEntry ClasspathEntry
}

func Parse(jrePath, classpath string) *Classloader {
	classloader := &Classloader{}
	classloader.parseBootstrapAndExtensionsClasspathEntry(jrePath)
	classloader.parseUserClasspathEntry(classpath)

	return classloader
}

func (classloader *Classloader) parseBootstrapAndExtensionsClasspathEntry(jrePath string) {
	jreDirectory := getJreDirectory(jrePath)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDirectory, "lib", "*")
	classloader.BootstrapClasspathEntry = NewWildcardClasspathEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtensionsPath := filepath.Join(jreDirectory, "lib", "ext", "*")
	classloader.ExtensionsClasspathEntry = NewWildcardClasspathEntry(jreExtensionsPath)
}

func (classloader *Classloader) parseUserClasspathEntry(classpath string) {
	if classpath == "" {
		classpath = "."
	}

	classloader.UserClasspathEntry = NewClasspathEntry(classpath)
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

func (classloader *Classloader) ReadClass(className string) ([]byte, ClasspathEntry, error) {
	className = className + ".class"

	data, classpathEntry, err := classloader.BootstrapClasspathEntry.ReadClass(className)

	if err == nil {
		return data, classpathEntry, err
	}

	data, classpathEntry, err = classloader.ExtensionsClasspathEntry.ReadClass(className)

	if err == nil {
		return data, classpathEntry, err
	}

	return classloader.UserClasspathEntry.ReadClass(className)
}