package heap

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/classfile"
	"github.com/Frederick-S/jvmgo/classpath"
)

type ClassLoader struct {
	classFinder   *classpath.ClassFinder
	loadedClasses map[string]*Class
}

func NewClassLoader(classFinder *classpath.ClassFinder) *ClassLoader {
	return &ClassLoader{
		classFinder:   classFinder,
		loadedClasses: make(map[string]*Class),
	}
}

func (classLoader *ClassLoader) LoadClass(className string) *Class {
	class, ok := classLoader.loadedClasses[className]

	if ok {
		return class
	}

	return classLoader.LoadNonArrayClass(className)
}

func (classLoader *ClassLoader) LoadNonArrayClass(className string) *Class {
	classData, classpathEntry := classLoader.ReadClass(className)
	class := classLoader.DefineClass(classData)

	linkClass(class)

	fmt.Printf("[Loaded class %s from %s]\n", className, classpathEntry)

	return class
}

func (classLoader *ClassLoader) ReadClass(className string) ([]byte, classpath.ClasspathEntry) {
	classData, classpathEntry, err := classLoader.classFinder.ReadClass(className)

	if err != nil {
		panic("java.lang.ClassNotFoundException: " + className)
	}

	return classData, classpathEntry
}

func (classLoader *ClassLoader) DefineClass(classData []byte) *Class {
	class := parseClassData(classData)
	class.classLoader = classLoader

	resolveSuperClass(class)
	resolveSuperClass(class)

	classLoader.loadedClasses[class.name] = class

	return class
}

func parseClassData(classData []byte) *Class {
	classFile, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	return newClass(classFile)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.classLoader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfacesCount := len(class.interfaceNames)

	if interfacesCount > 0 {
		class.interfaces = make([]*Class, interfacesCount)

		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
	}
}

func linkClass(class *Class) {
	verifyClass(class)
	prepareClass(class)
}

func verifyClass(class *Class) {

}

func prepareClass(class *Class) {

}
