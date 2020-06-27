package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// construct `sleep.sh` command
	cmd := &exec.Cmd{
		Path:   "/home/divierda/go/src/go-plasma-no-shadow/no-shadow.sh",
		Args:   []string{"/home/divierda/go/src/go-plasma-no-shadow/no-shadow.sh"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fiveMin := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-fiveMin.C:
			// run `cmd` in background
			fmt.Println("Executed")
			cmd.Run()

		}
	}
}
