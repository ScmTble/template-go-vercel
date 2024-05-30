package handler

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	output,err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	fmt.Fprint(w, string(output),hostname)

	output,err = exec.Command("ps", "-ef").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprint(w, string(output))
}
