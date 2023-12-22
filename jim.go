package main

import (
	"jim/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "jim",
		Usage: "jim is a git wrapper",
		Action: func(cCtx *cli.Context) error {
			commands.Run(cCtx)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "init repository in current directory",
				Action: func(cCtx *cli.Context) error {
					commands.Init(cCtx)
					return nil
				},
			},
			{
				Name:  "ignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					commands.IgnoreFiles(cCtx)
					return nil
				},
			},
			{
				Name:  "unignore",
				Usage: "remove repository in current directory",
				Action: func(cCtx *cli.Context) error {
					commands.UnIgnoreFiles(cCtx)
					return nil
				},
			},
			{
				Name:  "checkout",
				Usage: "git checkout",
				Action: func(cCtx *cli.Context) error {
					commands.Checkout(cCtx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
