package main

import (
	"quiz2/cli"
	"quiz2/router"
)

func main() {
	var cliOnly bool

	cliOnly = cli.CheckFlags()

	if !cliOnly {
		router.InitRouter()
	}
}
