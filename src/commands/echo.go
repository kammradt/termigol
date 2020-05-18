package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	echo(args)
}

func echo(args []string) {
	if len(args) < 1 {
		fmt.Println("You need to give a second argumento to 'echo' command")
		return
	}
	content := args[0]
	fmt.Println(content)
}
