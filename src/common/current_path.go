package common

import (
	"fmt"
	"os"
)

func CurrentPath(args []string) string {
	path, err := os.Getwd()
	if err != nil {
		CouldNotGetCurrentPathMessage := "Could not get current path."
		fmt.Println(CouldNotGetCurrentPathMessage)
		const CouldNotGetCurrentPathCode int = 2
		os.Exit(CouldNotGetCurrentPathCode)
	}
	return path
}
