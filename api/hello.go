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

	output,err = exec.Command("ps", "-ef").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	
	resp += string(output)
	fmt.Fprint(w, resp)
}
