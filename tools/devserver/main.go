package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func run_cmd(w http.ResponseWriter, req *http.Request) {
	log.Printf("Recieved a request. [%s] %s", req.Method, req.URL)

	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse form. Error: %s", err)
		log.Printf("Failed to parse form. Error: %s", err)
		return
	}

	log.Printf("Form Data: %s", req.Form)
	formDate := req.Form
	cmdStr, cmd_set := formDate["command"]
	query, query_set := formDate["query"]

	if !(cmd_set && query_set) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing query or command")
		log.Printf("Missing query or command. command=%s, query=%s", cmdStr, query)
		return
	}

	cmd := exec.Command("go", "run", ".", "--command", cmdStr[0], "--query", query[0])
	stdout, err := cmd.Output()

	if err == nil {
		w.Write(stdout)
	} else if exit_err, ok := err.(*exec.ExitError); ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(exit_err.Stderr)
		log.Printf("Command failed. Error: %s. StdError: %s", err, exit_err.Stderr)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Command failed. Error: %s", err)
		log.Printf("Command failed. Error: %s", err)
	}
}

func main() {
	http.HandleFunc("/", run_cmd)
	http.ListenAndServe(":8080", nil)
}
