package main

import (
	"blog-system/cmd"
	"blog-system/pkg/bootstrap"
)

func main() {
	// this is for command line application currently disabled
	cmd.Execute()
	serve()
}

func serve() {
	bootstrap.Serve()
}
