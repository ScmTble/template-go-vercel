package handler

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	var resp string

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	resp += fmt.Sprintf("Hello from %s\n", hostname)

	output,err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	resp += string(output)

	output,err = exec.Command("whoami").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	
	resp += string(output)


	outputs,err := os.Executable()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	resp += outputs

	fmt.Fprint(w, resp)
}
