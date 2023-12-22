package main

import (
	"jim/jim"
	"log"
	"os"

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
					jim.Init(cCtx)
					return nil
				},
			},
			{
				Name:  "ignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					jim.IgnoreFiles(cCtx)
					return nil
				},
			},
			{
				Name:  "unignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					jim.UnIgnoreFiles(cCtx)
					return nil
				},
			},
			{
				Name:  "checkout",
				Usage: "git checkout",
				Action: func(cCtx *cli.Context) error {
					jim.Checkout(cCtx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
