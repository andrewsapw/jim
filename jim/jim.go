package jim

import (
	"fmt"
	"jim/jimgit"
	"jim/path"
	"jim/run"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func Init(cCtx *cli.Context) {
	currentPath := path.CurrentPath()
	initRepoCommand := fmt.Sprintf("git init %s", currentPath)
	output := run.RunGitCommand(initRepoCommand, true)
	fmt.Print(output)
}

func IgnoreFiles(cCtx *cli.Context) {
	ignorePath := cCtx.Args().Get(0)
	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --assume-unchanged %s", configPath)

	run.RunGitCommand(command, false)
}

func UnIgnoreFiles(cCtx *cli.Context) {
	ignorePath := cCtx.Args().Get(0)

	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --no-assume-unchanged %s", configPath)

	run.RunGitCommand(command, false)
}

func createStashName(branchName string) string {
	return fmt.Sprintf(`jims checkout: %s`, branchName)
}

func Checkout(cCtx *cli.Context) {
	targetBranch := cCtx.Args().Get(0)
	targetBranch = strings.Trim(targetBranch, " \n")

	currentBranch := run.RunGitCommand("branch --show-current", false)
	stashName := createStashName(currentBranch)

	jimgit.CreateStash(stashName)
	jimgit.CheckoutBranch(targetBranch)

	// // maybe we can reset previously saved stash
	prevStashName := createStashName(targetBranch)
	prevStashIndex, err := jimgit.GetStashIndex(prevStashName)

	if err == nil {
		jimgit.PopStash(prevStashIndex)
	}
}
