package main

import (
	"fmt"
	"strings"

	"github.com/Frederick-S/jvmgo/classfile"
	"github.com/Frederick-S/jvmgo/classpath"
)

func main() {
	cmd := parseCmd()

	if cmd.showVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.showHelp || cmd.className == "" {
		printCmdUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	classloader := classpath.Parse(cmd.jrePath, cmd.classpath)
	className := strings.Replace(cmd.className, ".", "/", -1)
	classFile := loadClass(className, classloader)

	printClassInfo(classFile)
}

func loadClass(className string, classloader *classpath.Classloader) *classfile.ClassFile {
	classData, _, err := classloader.ReadClass(className)

	if err != nil {
		panic(err)
	}

	classFile, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	return classFile
}

func printClassInfo(classFile *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", classFile.GetMajorVersion(), classFile.GetMinorVersion())
	fmt.Printf("Constants count: %v\n", len(classFile.GetConstantPool()))
	fmt.Printf("Access flags: 0x%x\n", classFile.GetAccessFlags())
	fmt.Printf("This class: %v\n", classFile.GetClassName())
	fmt.Printf("Super class: %v\n", classFile.GetSuperClassName())
	fmt.Printf("Interfaces: %v\n", classFile.GetInterfaceNames())
	fmt.Printf("Fields count: %v\n", len(classFile.GetFields()))

	for _, field := range classFile.GetFields() {
		fmt.Printf(" %s\n", field.GetName())
	}

	fmt.Printf("Methods count: %v\n", len(classFile.GetMethods()))

	for _, method := range classFile.GetMethods() {
		fmt.Printf(" %s\n", method.GetName())
	}
}
