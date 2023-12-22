package main

import (
	"jim/jim"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "init repository in current directory",
				Action: func(cCtx *cli.Context) error {
					jim.Init()
					return nil
				},
			},
			{
				Name:  "ignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					ignorePath := cCtx.Args().Get(0)
					jim.IgnoreFiles(ignorePath)
					return nil
				},
			},
			{
				Name:  "unignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					ignorePath := cCtx.Args().Get(0)
					jim.UnIgnoreFiles(ignorePath)
					return nil
				},
			},
			{
				Name:  "checkout",
				Usage: "git checkout",
				Action: func(cCtx *cli.Context) error {
					targetBranch := cCtx.Args().Get(0)
					targetBranch = strings.Trim(targetBranch, " \n")
					jim.Checkout(targetBranch)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
