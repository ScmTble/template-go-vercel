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
		fmt.Println("Error:", err)
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Hostname:", hostname)
	
	fmt.Fprint(w, string(output))
}
