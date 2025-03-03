package main

import (
	"flag"
)

var llpkgConfigPath = flag.String("config", "", "path to config file")

func main() {
	flag.Parse()
	if *llpkgConfigPath == "" {
		printUsage()
		return
	}

	println("TODO")
}

func printUsage() {
	println("Usage: llpkg -config <path to config file>")
}
