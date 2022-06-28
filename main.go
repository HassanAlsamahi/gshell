package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
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
		home_dir, _ := os.UserHomeDir()
		if len(args) < 2 {
			return os.Chdir(home_dir)
		} else {
			return os.Chdir(args[1])
		}

	case "\f":
		args[0] = "clear"
	}

	return cmd.Run()
}

func main() {
	fmt.Println("Hello to Gshell!!")

	reader := bufio.NewReader(os.Stdin)
	for {
		hostname, _ := os.Hostname()
		workDir, _ := os.Getwd()
		user, _ := user.Current()

		fmt.Printf("[%s~%s~%s]> ", user.Username, hostname, workDir)

		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\nGoodbye!!")
			os.Exit(0)
		} else if err != nil {
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
