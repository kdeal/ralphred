package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type AlfredItem struct {
	Title string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Arg []string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
}

func alfredItemFromString(str string) AlfredItem {
	return AlfredItem{Title: str, Subtitle: "", Arg: []string{str}, Autocomplete: str}
}

type AlfredResponse struct {
	Items []AlfredItem `json:"items"`
}

func stringCommands() {
	resp := AlfredResponse {
		Items: []AlfredItem{
			alfredItemFromString("lower"),
			alfredItemFromString("title"),
			alfredItemFromString("upper"),
			alfredItemFromString("pymod"),
			alfredItemFromString("unpymod"),
		},
	}
	json_data, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error formatting string commands json")
		return
	}
	fmt.Println(string(json_data))
}

func stringCommand(args []string) {
	if len(args) == 0 {
		stringCommands()
		return
	}

	if len(args) == 1 {
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
