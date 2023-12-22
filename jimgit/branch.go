package jimgit

import (
	"fmt"
	"jim/errutils"
	"jim/run"

	git "github.com/libgit2/git2go/v34"
)

func GetBranch(branchName string) *git.Branch {
	repo := GetRepository()

	branch, err := repo.LookupBranch(branchName, git.BranchLocal)
	errutils.ProcessError(err)

	return branch
}

func GetHead() *git.Reference {
	repo := GetRepository()

	head, err := repo.Head()
	errutils.ProcessError(err)

	return head
}

func CheckoutBranch(branchName string) {
	checkoutCommand := fmt.Sprintf("checkout %s", branchName)
	run.RunGitCommand(checkoutCommand, false)
}
