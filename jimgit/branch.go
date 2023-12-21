package jimgit

import (
	"fmt"
	"jim/errutils"

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
	repo := GetRepository()
	branch := GetBranch(branchName)

	localCommit, err := repo.LookupCommit(branch.Target())
	errutils.ProcessError(err)

	fmt.Println(localCommit.Message())

	tree, err := repo.LookupTree(localCommit.TreeId())
	errutils.ProcessError(err)

	// checkout branch
	repo.CheckoutTree(tree, &git.CheckoutOptions{
		Strategy: git.CheckoutSafe | git.CheckoutRecreateMissing | git.CheckoutAllowConflicts | git.CheckoutUseTheirs,
	})
	// repo.SetHead(branch.Reference.Name())
}
