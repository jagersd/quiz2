package main

import (
	"quiz2/cli"
)

func main() {
	var cliOnly bool

	cliOnly = cli.CheckFlags()

	if !cliOnly {
		InitRouter()
	}
}
