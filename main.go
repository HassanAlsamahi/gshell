package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	switch args[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		return os.Chdir(args[1])
	}

	return cmd.Run()
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		hostname, _ := os.Hostname()

		fmt.Printf("[~%s~]> ", hostname)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			if err.Error() == io.EOF.Error() {
				os.Exit(0)
			} else if err.Error() != "exec: no command" {
				fmt.Fprintln(os.Stderr, err)
			}

		}
	}
}
