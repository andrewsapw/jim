package jimgit

import (
	"fmt"
	"jim/execute"
	"strings"

	"github.com/urfave/cli/v2"
)

func GetStashIndex(stashName string, cCtx *cli.Context) (int, error) {
	getStashesCommand := "stash list"
	stashesStr := execute.RunGitCommand(getStashesCommand, true, cCtx)

	// find stash index by name
	for index, stash := range strings.Split(stashesStr, "\n") {
		if strings.Contains(stash, stashName) {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Stash not found")
}

func CreateStash(stashName string, cCtx *cli.Context) {
	createStashCommand := fmt.Sprintf("stash save %s", stashName)
	execute.RunGitCommand(createStashCommand, true, cCtx)
}

func DeleteStashByName(stashName string, cCtx *cli.Context) {
	stashIndex, err := GetStashIndex(stashName, cCtx)
	if err != nil {
		return
	}

	DeleteStashByIndex(stashIndex, cCtx)
}

func DeleteStashByIndex(stashIndex int, cCtx *cli.Context) {
	deleteStashCommand := fmt.Sprintf("stash drop stash@{%d}", stashIndex)
	execute.RunGitCommand(deleteStashCommand, true, cCtx)
}

func PopStashByIndex(stashIndex int, cCtx *cli.Context) {
	popStashCommand := fmt.Sprintf("stash pop stash@{%d}", stashIndex)
	execute.RunGitCommand(popStashCommand, true, cCtx)
}

func PopStashByName(stashName string, cCtx *cli.Context) {
	stashIndex, err := GetStashIndex(stashName, cCtx)
	if err != nil {
		return
	}

	PopStashByIndex(stashIndex, cCtx)
}
