package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	showHelp bool
	showVersion bool
	classPath string
	className string
	arguments []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.showHelp, "help", false, "Print help message")
	flag.BoolVar(&cmd.showHelp, "?", false, "Print help message")
	flag.BoolVar(&cmd.showVersion, "version", false, "Print version and exit")
	flag.StringVar(&cmd.classPath, "classpath", "", "Classpath")
	flag.StringVar(&cmd.classPath, "cp", "", "Classpath")
	flag.Parse()

	arguments := flag.Args()

	if len(arguments) > 0 {
		cmd.className = arguments[0]
		cmd.arguments = arguments[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] className [arguments...]\n", os.Args[0])
}