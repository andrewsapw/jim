package jimgit

import (
	"fmt"
	"jim/execute"
	"strings"
)

func GetStashIndex(stashName string) (int, error) {
	getStashesCommand := "stash list"
	stashesStr := execute.RunGitCommand(getStashesCommand, true)

	// find stash index by name
	for index, stash := range strings.Split(stashesStr, "\n") {
		if strings.Contains(stash, stashName) {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Stash not found")
}

func CreateStash(stashName string) {
	createStashCommand := fmt.Sprintf("stash save %s", stashName)
	execute.RunGitCommand(createStashCommand, true)
}

func DeleteStashByName(stashName string) {
	stashIndex, err := GetStashIndex(stashName)
	if err != nil {
		return
	}

	DeleteStashByIndex(stashIndex)
}

func DeleteStashByIndex(stashIndex int) {
	deleteStashCommand := fmt.Sprintf("stash drop stash@{%d}", stashIndex)
	execute.RunGitCommand(deleteStashCommand, true)
}

func PopStashByIndex(stashIndex int) {
	popStashCommand := fmt.Sprintf("stash pop stash@{%d}", stashIndex)
	execute.RunGitCommand(popStashCommand, true)
}

func PopStashByName(stashName string) {
	stashIndex, err := GetStashIndex(stashName)
	if err != nil {
		return
	}

	PopStashByIndex(stashIndex)
}
