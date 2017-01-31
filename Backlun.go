package main

import (
	"fmt"
	"os"

	"Backlun/back/blog"
	"Backlun/back/oauth"
	"Backlun/back/todo"
)

func main() {
	args := os.Args
	if ((len(args) == 3) || (len(args) == 4)) && args[1] == "start" {
		startServer(args)
	} else if len(args) == 3 && args[1] == "help" {
		printPlatformHelp(args[2])
	} else if len(args) == 2 && args[1] == "help" {
		printFullHelp()
	} else {
		incorrectCommand()
	}
}

func incorrectCommand() {
	fmt.Println("---------------")
	fmt.Println("ERROR")
	fmt.Println("Incorrect command")
	fmt.Println("For help run \"./Backlun help\"")
	fmt.Println("---------------")
}

func printFullHelp() {
	//
}

func printPlatformHelp(platform string) {
	switch platform {
	case "todo":
		//
	case "blog":
		//
	case "market":
		//
	case "forum":
		//
	default:
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Incorrect platform")
		fmt.Println("---------------")
	}
}

func startServer(args []string) {
	switch args[2] {
	case "todo":
		todo.Start(args)
	case "blog":
		blog.Start(args)
	case "market":
		// market.Start(args)
	case "forum":
		//
	case "oauth":
		oauth.Start(args)
	default:
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Incorrect platform")
		fmt.Println("---------------")
	}
}
