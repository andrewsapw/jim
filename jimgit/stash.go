package jimgit

import (
	"fmt"
	"jim/errutils"
	"strings"

	git "github.com/libgit2/git2go/v34"
)

func GetStashes() git.StashCollection {
	repo := GetRepository()
	stashes := repo.Stashes
	return stashes
}

func GetStashIndex(stashName string) int {
	// find stash by name
	stashes := GetStashes()
	var stashIndex = -1

	stashes.Foreach(func(index int, message string, id *git.Oid) error {
		fmt.Println(message)
		if strings.Contains(message, stashName) {
			stashIndex = index
			return fmt.Errorf("found")
		}
		return nil
	})

	fmt.Println(stashIndex)

	return stashIndex
}

func CreateStash(stashName string) {
	repo := GetRepository()

	stashIndex := GetStashIndex(stashName)

	if stashIndex != 0 {
		repo.Stashes.Drop(stashIndex)
	}
	// get signrautre
	signature, err := repo.DefaultSignature()
	errutils.ProcessError(err)

	repo.Stashes.Save(signature, stashName, git.StashDefault)
}

func PopStash(stashIndex int) {
	repo := GetRepository()

	repo.Stashes.Pop(stashIndex, git.StashApplyOptions{})
}
