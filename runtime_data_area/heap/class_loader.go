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
	assignInstanceFieldsVariableIndices(class)
	assignStaticFieldsVariableIndices(class)
	initializeStaticFinalVariables(class)
}

func assignInstanceFieldsVariableIndices(class *Class) {
	index := uint(0)

	if class.superClass != nil {
		index = class.superClass.instanceVariablesCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.variableIndex = index

			if field.IsLongOrDouble() {
				index += 2
			} else {
				index++
			}
		}
	}

	class.instanceVariablesCount = index
}

func assignStaticFieldsVariableIndices(class *Class) {
	index := uint(0)

	for _, field := range class.fields {
		if field.IsStatic() {
			field.variableIndex = index

			if field.IsLongOrDouble() {
				index += 2
			} else {
				index++
			}
		}
	}

	class.staticVariablesCount = index
}

func initializeStaticFinalVariables(class *Class) {
	class.staticVariables = newVariables(class.staticVariablesCount)

	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initializeStaticFinalVariable(class, field)
		}
	}
}

func initializeStaticFinalVariable(class *Class, field *Field) {
	staticVariables := class.staticVariables
	constantPool := class.constantPool
	constantValueIndex := field.GetConstantValueIndex()
	variableIndex := field.GetVariableIndex()

	if constantValueIndex > 0 {
		switch field.GetDescriptor() {
		case "Z", "B", "C", "S", "I":
			value := constantPool.GetConstant(constantValueIndex).(int32)
			staticVariables.SetIntegerValue(variableIndex, value)
		case "J":
			value := constantPool.GetConstant(constantValueIndex).(int64)
			staticVariables.SetLongValue(variableIndex, value)
		case "F":
			value := constantPool.GetConstant(constantValueIndex).(float32)
			staticVariables.SetFloatValue(variableIndex, value)
		case "D":
			value := constantPool.GetConstant(constantValueIndex).(float64)
			staticVariables.SetDoubleValue(variableIndex, value)
		case "Ljava/lang/String;":
			panic("TODO")
		}
	}
}
