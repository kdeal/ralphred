package main

import (
	"flag"
	"fmt"
	"strings"
)

func stringCommand(args []string) {
	if len(args) < 2 {
		return
	}

	subcmd := args[0]
	input_string := strings.Join(args[1:], " ")
	switch subcmd {
	case "lower":
		fmt.Println(strings.ToLower(input_string))
	case "title":
		fmt.Println(strings.ToTitle(input_string))
	case "upper":
		fmt.Println(strings.ToUpper(input_string))
	case "pymod":
		no_slashes := strings.Replace(input_string, "/", ".", -1)
		fmt.Println(strings.TrimSuffix(no_slashes, ".py"))
	case "unpymod":
		fmt.Println(strings.Replace(input_string, ".", "/", -1) + ".py")
	default:
		fmt.Println("Unknown string subcommand")
	}
}

func main() {
	cmdPtr := flag.String("command", "commands", "What alfred command is being called")

	flag.Parse()

	cmdStr := flag.Args()

	switch *cmdPtr {
	case "strings":
		stringCommand(cmdStr)
	}
}
