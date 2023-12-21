package jimgit

import (
	"fmt"
	"jim/errutils"
	"jim/path"

	git "github.com/libgit2/git2go/v34"
)

func InitRepository() {
	var currentPath = path.CurrentPath()
	var _, err = git.InitRepository(currentPath, false)

	errutils.ProcessError(err)
	fmt.Printf("Created new repository at: %s", currentPath)
}

func GetRepository() *git.Repository {
	repo, err := git.OpenRepository(path.CurrentPath())
	errutils.ProcessError(err)

	return repo
}
