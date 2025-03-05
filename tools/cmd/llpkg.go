package main

import (
	"flag"
	"fmt"

	"github.com/goplus/llpkg/tools/pkg/config"
)

var llpkgConfigPath = flag.String("config", "", "path to config file")

func main() {
	config, err := config.ParseLLpkgConfig("./llpkg-tool/_demo/.llpkg/llpkg.cfg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(config)
	fmt.Println(config.UpstreamConfig.InstallerConfig.Config["options"])
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
