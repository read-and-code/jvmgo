package main

import (
	"fmt"
	"strings"

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
	fmt.Printf("Classpath: %s, className: %s, arguments: %v\n", cmd.classpath, cmd.className, cmd.arguments)

	classloader := classpath.Parse(cmd.jrePath, cmd.classpath)
	className := strings.Replace(cmd.className, ".", "/", -1)
	classData, _, err := classloader.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.className)

		return
	}

	fmt.Printf("Class data: %v\n", classData)
}
