package ralphred

func Run(cmd string, args []string) {
	switch cmd {
	case "strings":
		stringCommand(args)
	}
}
