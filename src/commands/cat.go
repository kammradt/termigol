package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	cat(args)
}

func cat(args []string) {
	if len(args) < 1 {
		fmt.Println("You need to give a second argumento to 'cat' command")
		return
	}
	fileName := args[0]

	byteContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		CouldNotReadFileMessage := fmt.Sprintf("Could not read %s file", fileName)
		fmt.Println(CouldNotReadFileMessage)
	}

	text := string(byteContent)
	fmt.Println(text)
}