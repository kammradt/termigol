package main

import (
	"fmt"
	"os"
)

const EXIT string = "exit"
const HELP string = "help"
const LS string = "ls"
const PWD string = "pwd"
const ECHO string = "echo"
const CAT string  = "cat"

func main() {
	args := os.Args[1:]
	help(args)
}

func help(args []string) {
	fmt.Println("Hi there!")
	fmt.Println("At this moment, we have these commands: ")

	availableCommands := [6]string{EXIT, HELP, LS, PWD, ECHO, CAT}
	for _, command := range availableCommands {
		fmt.Println("> $", command)
	}
}
