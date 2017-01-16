package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 3 && args[1] == "start" {
		startServer(args[2])
	} else if len(args) == 3 && args[1] == "help" {
		printPlatformHelp(args[2])
	} else if len(args) == 2 && args[1] == "help" {
		printFullHelp()
	} else {
		incorrectCommand()
	}
}

func incorrectCommand() {
	fmt.Println("####")
	fmt.Println("Incorrect command")
	fmt.Println("For help run \"./Backlun help\"")
	fmt.Println("####")
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
		fmt.Println("####")
		fmt.Println("Incorrect platform")
		fmt.Println("####")
	}
}

func startServer(arg string) {
	switch arg {
	case "todo":
		//
	case "blog":
		//
	case "market":
		//
	case "forum":
		//
	default:
		fmt.Println("####")
		fmt.Println("Incorrect platform")
		fmt.Println("####")
	}
}
