package main

import (
	"fmt"
	"strings"

	"github.com/Frederick-S/jvmgo/classfile"
	"github.com/Frederick-S/jvmgo/classpath"
)

func main() {
	//cmd := parseCmd()
	cmd := &Cmd{
		jrePath:   "/Library/Java/JavaVirtualMachines/jdk1.8.0_92.jdk/Contents/Home/jre",
		className: "./java/Sum",
	}

	if cmd.showVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.showHelp || cmd.className == "" {
		printCmdUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	classFinder := classpath.Parse(cmd.jrePath, cmd.classpath)
	className := strings.Replace(cmd.className, ".", "/", -1)
	classFile := loadClass(className, classFinder)
	mainMethod := getMainMethod(classFile)

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.className)
	}
}

func loadClass(className string, classFinder *classpath.ClassFinder) *classfile.ClassFile {
	classData, _, err := classFinder.ReadClass(className)

	if err != nil {
		panic(err)
	}

	classFile, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	return classFile
}

func getMainMethod(classFile *classfile.ClassFile) *classfile.MemberInfo {
	for _, method := range classFile.GetMethods() {
		if method.GetName() == "main" && method.GetDescriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}

	return nil
}
