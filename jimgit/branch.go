package jimgit

import (
	"fmt"
	"jim/execute"
)

func CheckoutBranch(branchName string) {
	checkoutCommand := fmt.Sprintf("checkout %s", branchName)
	execute.RunGitCommand(checkoutCommand, false)
}

func SyncCurrentBranch() {
	// stash current changes
	stashName := "jims branch sync"
	CreateStash(stashName)
	defer PopStashByName(stashName)

	// fetch all branches
	fetchCommand := "fetch"
	execute.RunGitCommand(fetchCommand, false)

	// sync --ff-only
	pullCommand := "pull --ff-only"
	execute.RunGitCommand(pullCommand, false)
}
