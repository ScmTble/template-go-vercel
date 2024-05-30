package handler

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	output,err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprint(w, string(output))
}
