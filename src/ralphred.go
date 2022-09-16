package ralphred

import (
	"log"
	"strings"
)

func extract_args(query string) []string {
	if query == "" {
		return []string{}
	}

	args := []string{}
	args_raw := strings.Split(query, " ")
	for _, arg := range args_raw {
		if arg != "" {
			args = append(args, arg)
		}
	}
	return args
}

func Run(cmd string, query string) {
	args := extract_args(query)
	if query != "" {
	}
	log.Printf("cmd: %s, args: [%s]\n", cmd, strings.Join(args, ", "))

	var items []AlfredItem
	var err error

	switch cmd {
	case "strings":
		items, err = stringCommand(args)
	case "convert":
		items, err = convertCommand(args)
	case "datetimemath":
		items, err = dateTimeMathCommand(args)
	}

	if err != nil {
		items = errorAlfredItems(err.Error())
	}
	resp := AlfredResponse{Items: items}
	resp.Print()
}
