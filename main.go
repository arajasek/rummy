package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	cmds := []*cli.Command{}

	app := &cli.App{
		Name:     "rummy",
		Usage:    "A simulator for robber's rummy",
		Commands: cmds,
		Flags:    []cli.Flag{},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
		return
	}
}
