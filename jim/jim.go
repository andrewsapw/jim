package jim

import (
	"fmt"
	"jim/errutils"
	"jim/jimgit"
	"jim/path"
	"jim/run"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
)

func Init(s storage.Storer, worktree billy.Filesystem) {
	var currentPath = path.CurrentPath()
	var _, err = gogit.Init(s, worktree)

	errutils.ProcessError(err)
	fmt.Printf("Created new repository at: %s", currentPath)
}

func IgnoreFiles(ignorePath string) {
	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --assume-unchanged %s", configPath)

	run.RunGitCommand(command, false)
}

func UnIgnoreFiles(ignorePath string) {
	currentPath := path.CurrentPath()

	configPath := filepath.Join(currentPath, ignorePath)
	command := fmt.Sprintf("update-index --no-assume-unchanged %s", configPath)

	run.RunGitCommand(command, false)
}

func createStashName(branchName string) string {
	return fmt.Sprintf(`jims checkout: %s`, branchName)
}

func Checkout(targetBranch string) {
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
