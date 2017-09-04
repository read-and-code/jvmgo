package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	showHelp    bool
	showVersion bool
	jrePath     string
	classpath   string
	className   string
	arguments   []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printCmdUsage
	flag.BoolVar(&cmd.showHelp, "help", false, "Print help message")
	flag.BoolVar(&cmd.showHelp, "?", false, "Print help message")
	flag.BoolVar(&cmd.showVersion, "version", false, "Print version and exit")
	flag.StringVar(&cmd.jrePath, "jre", "", "jre path")
	flag.StringVar(&cmd.classpath, "classpath", "", "Classpath")
	flag.StringVar(&cmd.classpath, "cp", "", "Classpath")
	flag.StringVar(&cmd.className, "class", "", "Class name")
	flag.Parse()

	cmd.arguments = flag.Args()

	return cmd
}

func printCmdUsage() {
	fmt.Printf("Usage: %s -classpath class/path -class ClassName [arguments...]\n", os.Args[0])
}
