package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const EXIT string = "exit"
const HELP string = "help"
const LS string = "ls"
const PWD string = "pwd"
const CAT string = "cat"
const ECHO string = "echo"

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
	commands := strings.Fields(command)
	command = commands[0]
	switch command {
	case EXIT:
		exitHandler()
	case HELP:
		handleHelp()
	case LS:
		handleLs()
	case PWD:
		fmt.Println(getCurrentPath())
	case CAT:
		handleCat(commands)
	case ECHO:
		handleEcho(commands)
	default:
		handleCommandNotFound()
	}
	fmt.Println("")
}

func handleEcho(commands []string) {
	if len(commands) < 2 {
		fmt.Println("You need to give a second argumento to 'echo' command")
		return
	}
	content := commands[1]
	fmt.Println(content)
}

func handleCat(commands []string) {
	if len(commands) < 2 {
		fmt.Println("You need to give a second argumento to 'cat' command")
		return
	}
	fileName := commands[1]

	byteContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		CouldNotReadFileMessage := "Could not read " + fileName + "file."
		fmt.Println(CouldNotReadFileMessage)
	}

	text := string(byteContent)
	fmt.Println(text)
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

	availableCommands := [6]string{EXIT, HELP, LS, PWD, CAT, ECHO}
	for _, command := range availableCommands {
		fmt.Println("> $", command)
	}
}

func handleCommandNotFound() {
	fmt.Println("Command not found")
	fmt.Println("try `help`")
}
