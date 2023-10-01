package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Check if the user provided an argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <string>")
		return
	}

	// Get the string argument from the command line
	dockerImageName := os.Args[1]

	pullCmd(dockerImageName)
}

func pullCmd(imageName string) error {
	start := time.Now()
	cmd := exec.Command("docker", "pull", imageName)

	err := cmd.Run()
	if err != nil {
		return err
	}

	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Printf("Baseline Docker pull took %s\n", elapsedTime)

	return nil
}
