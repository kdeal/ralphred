package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type AlfredItem struct {
	UID string `json:"uid,omitempty"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Arg []string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
}

func alfredItemFromString(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID: uid,
		Title: str,
		Subtitle: "",
		Arg: []string{str},
		Autocomplete: str,
	}
}


func alfredItemFromStringForwarded(str string, set_uid bool) AlfredItem {
	uid := ""
	if set_uid {
		uid = str
	}

	return AlfredItem{
		UID: uid,
		Title: str,
		Subtitle: "",
		Arg: []string{fmt.Sprintf("%s ", str)},
		Autocomplete: str,
	}
}

type AlfredResponse struct {
	Items []AlfredItem `json:"items"`
}

func (resp AlfredResponse) Print() {
	json_data, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error formatting string commands json")
		return
	}
	fmt.Println(string(json_data))
}

func stringCommands() {
	resp := AlfredResponse {
		Items: []AlfredItem{
			alfredItemFromStringForwarded("lower", true),
			alfredItemFromStringForwarded("title", true),
			alfredItemFromStringForwarded("upper", true),
			alfredItemFromStringForwarded("pymod", true),
			alfredItemFromStringForwarded("unpymod", true),
		},
	}
	resp.Print()
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
	result := ""
	switch subcmd {
	case "lower":
		result = strings.ToLower(input_string)
	case "title":
		result = strings.ToTitle(input_string)
	case "upper":
		result = strings.ToUpper(input_string)
	case "pymod":
		no_slashes := strings.Replace(input_string, "/", ".", -1)
		result = strings.TrimSuffix(no_slashes, ".py")
	case "unpymod":
		result = strings.Replace(input_string, ".", "/", -1) + ".py"
	default:
		result = "Unknown string subcommand"
	}

	resp := AlfredResponse {
		Items: []AlfredItem{
			alfredItemFromString(result, false),
		},
	}
	resp.Print()
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
