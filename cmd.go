package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	showHelp bool
	showVersion bool
	jrePath string
	classPath string
	mainClass string
	arguments []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.showHelp, "help", false, "Print help message")
	flag.BoolVar(&cmd.showHelp, "?", false, "Print help message")
	flag.BoolVar(&cmd.showVersion, "version", false, "Print version and exit")
	flag.StringVar(&cmd.jrePath, "jre", "", "jre path")
	flag.StringVar(&cmd.classPath, "classpath", "", "Classpath")
	flag.StringVar(&cmd.classPath, "cp", "", "Classpath")
	flag.StringVar(&cmd.mainClass, "class", "", "Main class")
	flag.Parse()

	cmd.arguments = flag.Args()

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s -classpath class/path -class ClassName [arguments...]\n", os.Args[0])
}