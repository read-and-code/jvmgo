package main

import (
	"fmt"
	"strings"

	"github.com/Frederick-S/jvmgo/classpath"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

func main() {
	// cmd := parseCmd()
	cmd := &Cmd{
		jrePath:   "/Library/Java/JavaVirtualMachines/jdk1.8.0_92.jdk/Contents/Home/jre",
		className: "./java/HelloWorld",
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
	classLoader := heap.NewClassLoader(classFinder)
	className := strings.Replace(cmd.className, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.className)
	}
}
