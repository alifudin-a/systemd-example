package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// bash script path (root path)
	path := "./no-shadow.sh"

	// command to execute bash script
	cmd := exec.Command("/bin/bash", path)

	fiveMin := time.NewTicker(2 * time.Minute)

	// execute script for every 2 minutes
	for {
		select {
		case <-fiveMin.C:
			fmt.Println("2 minute passed: Script Executed")
			// start execute the bash script
			cmd.Run()
		}
	}
}
