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
		Usage: "jim is a git wrapper. if jim does not known command you entered, jim will run git command.",
		Action: func(cCtx *cli.Context) error {
			commands.Run(cCtx)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "show verbose output",
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "do not actually run command. use with --verbose flag",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "sw",
				Usage: "switch branch",
				Action: func(cCtx *cli.Context) error {
					commands.Checkout(cCtx)
					return nil
				},
			},
			{
				Name:  "sync",
				Usage: "sync branch",
				Action: func(cCtx *cli.Context) error {
					commands.Sync(cCtx)
					return nil
				},
			},
			{
				Name:  "ignore",
				Usage: "ignore files (git update-index --assume-unchanged)",
				Action: func(cCtx *cli.Context) error {
					commands.IgnoreFiles(cCtx)
					return nil
				},
			},
			{
				Name:  "unignore",
				Usage: "unignore files (git update-index --no-assume-unchanged)",
				Action: func(cCtx *cli.Context) error {
					commands.UnIgnoreFiles(cCtx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
