package main

import (
	"flag"
	ralphred "github.com/kdeal/ralphred/src"
)

func main() {
	cmdPtr := flag.String("command", "commands", "What alfred command is being called")
	queryPtr := flag.String("query", "query", "Query to send to the command")

	flag.Parse()

	ralphred.Run(*cmdPtr, *queryPtr)
}
