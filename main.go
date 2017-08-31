package main

import "fmt"

func main() {
	cmd := parseCmd()

	if cmd.showVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.showHelp || cmd.className == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("Classpath: %s, className: %s, arguments: %v\n", cmd.classPath, cmd.className, cmd.arguments)
}