package jimgit

import (
	"fmt"
	"jim/execute"
)

func CheckoutBranch(branchName string) {
	checkoutCommand := fmt.Sprintf("checkout %s", branchName)
	execute.RunGitCommand(checkoutCommand, false)
}
