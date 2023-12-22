package jimgit

import (
	"fmt"
	"jim/run"
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
	createStashCommand := fmt.Sprintf("stash save %s", stashName)
	run.RunGitCommand(createStashCommand, true)
}

func PopStash(stashIndex int) {
	popStashCommand := fmt.Sprintf("stash pop stash@{%d}", stashIndex)
	run.RunGitCommand(popStashCommand, true)
}
