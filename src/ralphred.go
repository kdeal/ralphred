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
	switch cmd {
	case "strings":
		stringCommand(args)
	case "convert":
		convertCommand(args)
	case "datetimemath":
		dateTimeMathCommand(args)
	}
}
