package main

import (
	"bytes"
	"errors"
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
	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf
	stdout, err := cmd.Output()
	stderr := errbuf.Bytes()

	if err == nil {
		w.Write(stdout)
		log.Printf("Command Succeeded. StdError: \n%s", stderr)
	} else if errors.Is(err, err.(*exec.ExitError)) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(stderr)
		log.Printf("Command failed. Error: %s. StdError: %s", err, stderr)
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
