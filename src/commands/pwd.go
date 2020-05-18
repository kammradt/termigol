package main

import (
	"awesomeProject/src/common"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	pwd(args)
}

func pwd(args []string) {
	currentPath := common.CurrentPath(args)
	fmt.Println(currentPath)
}