package main

import (
	"fmt"
	"os"

	"github.com/lucas-campelo/go-jenkins/command"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		command.Help()
		os.Exit(0)
	}

	switch args[0] {
	case "up":
		if err := command.Up(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "down":
		if err := command.Down(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		command.Help()
	default:
		fmt.Println("Invalid command line argument")
		os.Exit(1)
	}

	os.Exit(0)
}
