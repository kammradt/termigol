package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const LS string = "ls"
const EXIT string = "exit"
const HELP string = "help"

func main() {
	for {
		command := readCommand()
		commandExecutionHandler(command)
	}
}

func readCommand() string {
	path := getCurrentPath()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(path + " > $ ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func commandExecutionHandler(command string) {
	switch command {
	case EXIT:
		exitHandler()
	case HELP:
		handleHelp()
	case LS:
		handleLs()
	default:
		handleCommandNotFound()
	}
	fmt.Println("")
}

func exitHandler() {
	ExitedWithSuccessMessage := "Exiting..."
	fmt.Println(ExitedWithSuccessMessage)
	const ExitedWithSuccessCode int = 1
	os.Exit(ExitedWithSuccessCode)
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		CouldNotGetCurrentPathMessage := "Could not get current path."
		fmt.Println(CouldNotGetCurrentPathMessage)
		const CouldNotGetCurrentPathCode int = 2
		os.Exit(CouldNotGetCurrentPathCode)
	}
	return path
}

func handleLs() {
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

func handleHelp() {
	fmt.Println("Hi there!")
	fmt.Println("At this moment, we have these commands: ")

	availableCommands := [3]string{LS, EXIT, HELP}
	for _, command := range availableCommands {
		fmt.Println("> $", command)
	}
}

func handleCommandNotFound() {
	fmt.Println("Command not found")
	fmt.Println("try `help`")
}
