package main

import "fmt"

func main() {
	cmd := parseCmd()

	if cmd.showVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.showHelp || cmd.mainClass == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("Classpath: %s, mainClass: %s, arguments: %v\n", cmd.classPath, cmd.mainClass, cmd.arguments)
}