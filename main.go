package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func execInput(input string, work_dir_cache map[int]string, counter int) (error, map[int]string, int) {

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
		if len(args) < 2 || args[1] == "" {
			counter += 1
			work_dir_cache[counter] = home_dir
			return os.Chdir(home_dir), work_dir_cache, counter
		} else if args[1] == "-" {
			if len(work_dir_cache) < 2 {
				fmt.Println("There is no old working directory")
				return nil, work_dir_cache, counter
			}
			arg_int := counter
			old_work_dir := work_dir_cache[counter]
			if len(args) > 2 {
				arg_int, _ = strconv.Atoi(args[2])
				old_work_dir = work_dir_cache[arg_int]
			} else {
				old_work_dir = work_dir_cache[arg_int-1]
			}
			counter += 1
			work_dir_cache[counter] = old_work_dir
			return os.Chdir(old_work_dir), work_dir_cache, counter
		} else {
			if os.Chdir(args[1]) == nil {
				counter += 1
				work_dir_cache[counter] = args[1]

			}
			return os.Chdir(args[1]), work_dir_cache, counter
		}

	case "\f":
		args[0] = "clear"
	case "wdcache":
		fmt.Println(work_dir_cache)
		return nil, work_dir_cache, counter
	}

	return cmd.Run(), work_dir_cache, counter
}

func main() {
	fmt.Println("Hello to Gshell!!")
	homedir, _ := os.UserHomeDir()
	if err := os.Chdir(homedir); err != nil {

	}

	workDir, _ := os.Getwd()
	counter := 1
	work_dir_cache := map[int]string{counter: workDir}

	reader := bufio.NewReader(os.Stdin)
	for {
		hostname, _ := os.Hostname()
		user, _ := user.Current()

		fmt.Printf("[%s~%s~%s]> ", user.Username, hostname, work_dir_cache[counter])

		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\nGoodbye!!")
			os.Exit(0)
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err, work_dir_cache, counter = execInput(input, work_dir_cache, counter); err != nil {
			if err.Error() == io.EOF.Error() {
				os.Exit(0)
			} else if err.Error() != "exec: no command" {
				fmt.Fprintln(os.Stderr, err)
			}

		}
		//fmt.Println(work_dir_cache)
	}
}
