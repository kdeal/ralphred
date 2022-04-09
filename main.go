package main

import (
	"flag"
    ralphred "github.com/kdeal/ralphred/src"
)

func main() {
	cmdPtr := flag.String("command", "commands", "What alfred command is being called")

	flag.Parse()

	cmdStr := flag.Args()
    ralphred.Run(*cmdPtr, cmdStr)
}
