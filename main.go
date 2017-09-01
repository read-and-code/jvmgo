package main

import (
	"strings"
	"fmt"
	"github.com/Frederick-S/jvmgo/classpath"
)

func main() {
	cmd := parseCmd()

	if cmd.ShowVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.ShowHelp || cmd.ClassName == "" {
		printCmdUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("Classpath: %s, className: %s, arguments: %v\n", cmd.Classpath, cmd.ClassName, cmd.Arguments)

	classloader := classpath.Parse(cmd.JrePath, cmd.Classpath)
	className := strings.Replace(cmd.ClassName, ".", "/", -1)
	classData, _, err := classloader.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.ClassName)

		return
	}

	fmt.Printf("Class data: %v\n", classData)
}