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
	mainMethod := getMainMethod(classFile)

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.className)
	}
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

func getMainMethod(classFile *classfile.ClassFile) *classfile.MemberInfo {
	for _, method := range classFile.GetMethods() {
		if method.GetName() == "main" && method.GetDescriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}

	return nil
}
