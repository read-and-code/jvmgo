package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	ShowHelp bool
	ShowVersion bool
	JrePath string
	Classpath string
	ClassName string
	Arguments []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printCmdUsage
	flag.BoolVar(&cmd.ShowHelp, "help", false, "Print help message")
	flag.BoolVar(&cmd.ShowHelp, "?", false, "Print help message")
	flag.BoolVar(&cmd.ShowVersion, "version", false, "Print version and exit")
	flag.StringVar(&cmd.JrePath, "jre", "", "jre path")
	flag.StringVar(&cmd.Classpath, "classpath", "", "Classpath")
	flag.StringVar(&cmd.Classpath, "cp", "", "Classpath")
	flag.StringVar(&cmd.ClassName, "class", "", "Class name")
	flag.Parse()

	cmd.Arguments = flag.Args()

	return cmd
}

func printCmdUsage() {
	fmt.Printf("Usage: %s -classpath class/path -class ClassName [arguments...]\n", os.Args[0])
}