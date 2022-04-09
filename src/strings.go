package ralphred

import (
	"strings"
)

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
