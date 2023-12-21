package run

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func safeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}

func RunGitCommand(command string, allowFail bool) string {
	fmt.Printf("Running: git %s\n", command)

	commandArgs, _ := Split(command)

	cmd := exec.Command("git", commandArgs...)

	var outBuffer, errBuffer bytes.Buffer

	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()

	if err != nil && !allowFail {
		fmt.Println(string(errBuffer.Bytes()))
		os.Exit(1)
	}

	return strings.Trim(string(outBuffer.Bytes()), " \n")
}
