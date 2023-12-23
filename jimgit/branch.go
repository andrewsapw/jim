package jimgit

import (
	"fmt"
	"jim/execute"

	"github.com/urfave/cli/v2"
)

func CheckoutBranch(branchName string, cCtx *cli.Context) {
	checkoutCommand := fmt.Sprintf("checkout %s", branchName)
	execute.RunGitCommand(checkoutCommand, false, cCtx)
}

func SyncCurrentBranch(cCtx *cli.Context) {
	// stash current changes
	stashName := "jims branch sync"
	CreateStash(stashName, cCtx)
	defer PopStashByName(stashName, cCtx)

	// fetch all branches
	fetchCommand := "fetch"
	execute.RunGitCommand(fetchCommand, false, cCtx)

	// sync --ff-only
	pullCommand := "pull --ff-only"
	execute.RunGitCommand(pullCommand, false, cCtx)

	// push current branch
	pushCommand := "push -u"
	execute.RunGitCommand(pushCommand, false, cCtx)
}
