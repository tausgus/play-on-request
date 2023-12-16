package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const port string = ":1234"

var cmd = exec.Command("paplay", strings.Join(os.Args[1:], "")) // Get the file from command line arguments

func play(w http.ResponseWriter, r *http.Request) {
	if cmd.Start() != nil {
		fmt.Println("paplay failed to execute")
	} else {
		fmt.Println("playing audio")
	}
}

func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stopping command")
	cmd.Process.Kill()
}

func main() {
	// Check if the file argument was supplied and is valid
	if !strings.HasSuffix(strings.Join(os.Args[1:], ""), ".mp3") {
		fmt.Println("file argument missing or non-mp3 file provided")
		os.Exit(2)
	}

	http.HandleFunc("/play", play)
	http.HandleFunc("/stop", stop)
	http.ListenAndServe(port, nil)
}
