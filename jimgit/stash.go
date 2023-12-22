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

func GetStashIndex(stashName string) (int, error) {
	getStashesCommand := "stash list"
	stashesStr := run.RunGitCommand(getStashesCommand, true)

	fmt.Println(stashesStr)

	// find stash index by name
	for index, stash := range strings.Split(stashesStr, "\n") {
		fmt.Println(stash)
		if strings.Contains(stash, stashName) {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Stash not found")
}

func CreateStash(stashName string) {
	createStashCommand := fmt.Sprintf("stash save %s", stashName)
	run.RunGitCommand(createStashCommand, true)
}

func PopStash(stashIndex int) {
	popStashCommand := fmt.Sprintf("stash pop stash@{%d}", stashIndex)
	run.RunGitCommand(popStashCommand, true)
}
