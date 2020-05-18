package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	ls(args)
}

func ls(args []string) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		CouldNotReadCurrentPathMessage := "Could not read current path."
		fmt.Println(CouldNotReadCurrentPathMessage)
		const CouldNotReadCurrentPathCode int = 3
		os.Exit(CouldNotReadCurrentPathCode)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
