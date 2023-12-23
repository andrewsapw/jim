package commands

import (
	"fmt"
	"jim/execute"
	"jim/jimgit"
	"jim/path"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func IgnoreFiles(cCtx *cli.Context) {
	ignorePath := cCtx.Args().Get(0)
	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --assume-unchanged %s", configPath)

	execute.RunGitCommand(command, false, cCtx)
}

func UnIgnoreFiles(cCtx *cli.Context) {
	ignorePath := cCtx.Args().Get(0)

	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --no-assume-unchanged %s", configPath)

	execute.RunGitCommand(command, false, cCtx)
}

func createStashName(branchName string) string {
	return fmt.Sprintf(`jims checkout: %s`, branchName)
}

func Checkout(cCtx *cli.Context) {
	targetBranch := cCtx.Args().Get(0)
	targetBranch = strings.Trim(targetBranch, " \n")

	currentBranch := execute.RunGitCommand("branch --show-current", false, cCtx)
	stashName := createStashName(currentBranch)

	jimgit.CreateStash(stashName, cCtx)
	jimgit.CheckoutBranch(targetBranch, cCtx)

	// // maybe we can reset previously saved stash
	prevStashName := createStashName(targetBranch)
	prevStashIndex, err := jimgit.GetStashIndex(prevStashName, cCtx)

	if err == nil {
		jimgit.PopStashByIndex(prevStashIndex, cCtx)
	}
}

func Sync(cCtx *cli.Context) {
	jimgit.SyncCurrentBranch(cCtx)
}

func Run(cCtx *cli.Context) {
	args := cCtx.Args().Slice()
	command := strings.Join(args, " ")
	output := execute.RunGitCommand(command, false, cCtx)
	fmt.Print(output)
}
