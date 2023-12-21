package main

import (
	"jim/jim"
	"jim/path"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/urfave/cli/v2"
)

func main() {
	var currentPath = path.CurrentPath()
	var wt, dot billy.Filesystem

	wt = osfs.New(currentPath)
	dot, _ = wt.Chroot(".git")

	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "init repository in current directory",
				Action: func(cCtx *cli.Context) error {
					jim.Init(storage, dot)
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
