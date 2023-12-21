package errutils

import (
	"fmt"
	"os"
)

func ProcessError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
