package path

import (
	"fmt"
	"os"
)

func CurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path
}
