package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	ls(args)
}

func ls(args []string) {

	files := make([]os.FileInfo, 0)
	var err error
	if len(args) > 0 {
		files, err = ioutil.ReadDir("./" + args[0])
	} else {
		files, err = ioutil.ReadDir("./")
	}

	if err != nil {
		CouldNotReadCurrentPathMessage := "Could not read current path."
		fmt.Println(CouldNotReadCurrentPathMessage)
		const CouldNotReadCurrentPathCode int = 3
		os.Exit(CouldNotReadCurrentPathCode)
	}

	for _, f := range files {
		if f.IsDir() {
			_, _ = emoji.Print(":dividers: ")
		} else {
			_, _ = emoji.Print(":memo: ")
		}
		fmt.Println(f.Name())
	}
}
