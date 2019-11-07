package main

import (
	cli "github.com/urfave/cli/v2"
)

func startup(c *cli.Context) {
	go sigHandler()
	go serviceInit(c)
}
