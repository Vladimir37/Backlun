package main

import (
	"bufio"
	"fmt"
	"os"

	"Backlun/back/blog"
	"Backlun/back/chat"
	"Backlun/back/forum"
	"Backlun/back/geopos"
	"Backlun/back/market"
	"Backlun/back/oauth"
	"Backlun/back/todo"
	"strings"
)

func main() {
	startInit()
}

func startInit() {
	fmt.Println("Enter the command using the following format:")
	fmt.Println("<command> <platform> [<port>]")
	fmt.Println("Add the host and the key at the end for \"oauth\" platform.")
	fmt.Println("Enter \"help\" to display all commands and platforms:")
	readInput()
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		commandRouter(args)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Internal error. Please, try again.")
		fmt.Println("---------------")
		os.Exit(1)
	}
}

func commandRouter(args []string) {
	if ((len(args) == 2) || (len(args) == 3)) && args[0] == "start" {
		startServer(args)
	} else if len(args) == 2 && args[0] == "help" {
		printPlatformHelp(args[1])
	} else if len(args) == 1 && args[0] == "help" {
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
	fmt.Println("Format for any platform other than \"oauth\":")
	fmt.Println("\"<command> <platform> [<port>]\"")
	fmt.Println("Default port: 8000")
	fmt.Println("")
	fmt.Println("Format for \"oauth\" platform")
	fmt.Println("\"start oauth <port> <host> <key>\"")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("List of commands:")
	fmt.Println("— help")
	fmt.Println("— start")
	fmt.Println("")
	fmt.Println("List of platforms:")
	fmt.Println("— todo")
	fmt.Println("— blog")
	fmt.Println("— market")
	fmt.Println("— forum")
	fmt.Println("— oauth")
	fmt.Println("— chat")
	fmt.Println("")
	fmt.Println("You can run \"help\" and \"start\" for any platform")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("\"help market\"")
	fmt.Println("\"start todo\"")
	fmt.Println("\"start blog 8080\"")
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
	case "chat":
		//
	case "oauth":
		//
	case "geopos":
		//
	default:
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Incorrect platform")
		fmt.Println("---------------")
	}
}

func startServer(args []string) {
	switch args[1] {
	case "todo":
		todo.Start(args)
	case "blog":
		blog.Start(args)
	case "market":
		market.Start(args)
	case "forum":
		forum.Start(args)
	case "oauth":
		oauth.Start(args)
	case "chat":
		chat.Start(args)
	case "geopos":
		geopos.Start(args)
	default:
		fmt.Println("---------------")
		fmt.Println("ERROR")
		fmt.Println("Incorrect platform")
		fmt.Println("---------------")
	}
}
