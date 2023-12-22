package jimgit

import (
	"fmt"
	"jim/run"
)

func CheckoutBranch(branchName string) {
	checkoutCommand := fmt.Sprintf("checkout %s", branchName)
	run.RunGitCommand(checkoutCommand, false)
}
