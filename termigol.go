package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const ExitedWithSuccessCode int = 1
const CouldNotGetCurrentPathCode int = 2
const CouldNotReadCurrentPathCode int = 3

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
		exitHandler(ExitedWithSuccessCode)
	case HELP:
		handleHelp()
	case LS:
		handleLs()
	default:
		handleCommandNotFound()
	}
	fmt.Println("")
}

func exitHandler(errorCode int) {
	ExitedWithSuccessMessage := "Exiting..."
	CouldNotGetCurrentPathMessage := "Could not get current path."
	CouldNotReadCurrentPathMessage := "Could not read current path."

	switch errorCode {
	case ExitedWithSuccessCode:
		fmt.Println(ExitedWithSuccessMessage)
	case CouldNotGetCurrentPathCode:
		fmt.Println(CouldNotGetCurrentPathMessage)
	case CouldNotReadCurrentPathCode:
		fmt.Println(CouldNotReadCurrentPathMessage)
	}

	os.Exit(errorCode)
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		exitHandler(CouldNotGetCurrentPathCode)
	}
	return path
}

func handleLs() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		exitHandler(CouldNotReadCurrentPathCode)
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
