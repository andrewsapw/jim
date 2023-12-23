package execute

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
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

func RunGitCommand(command string, allowFail bool, cCtx *cli.Context) string {
	verbose := cCtx.Bool("verbose")
	dryRun := cCtx.Bool("dry-run")

	if verbose {
		if dryRun {
			fmt.Printf("jim (dry-run): git %s\n", command)
		} else {
			fmt.Printf("jim: git %s\n", command)
		}
	}
	commandArgs, _ := Split(command)

	cmd := exec.Command("git", commandArgs...)

	var outBuffer, errBuffer bytes.Buffer

	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	if dryRun {
		return ""
	}

	err := cmd.Run()

	if err != nil && !allowFail {
		fmt.Println(string(errBuffer.Bytes()))
		os.Exit(1)
	}

	return strings.Trim(string(outBuffer.Bytes()), " \n")
}
