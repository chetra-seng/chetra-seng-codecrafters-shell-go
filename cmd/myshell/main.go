package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input", err)
		}

		args := strings.Fields(command[:len(command)-1])

		pathEnv := os.Getenv("PATH")
		paths := strings.Split(pathEnv, ":")

		switch args[0] {

		// Exit command
		case "exit":
			if args[1] == "0" {
				os.Exit(0)
			}

		// Print command
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))

		// Type command
		case "type":
			if isBuiltin(args[1]) {
				fmt.Printf("%s is a shell builtin\n", args[1])
			} else {
				found := isExec(paths, args[1])

				if !found {
					fmt.Printf("%s not found\n", args[1])
				}
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", args[0])
			}
		}
	}
}

func isBuiltin(cmd string) bool {
	builtIns := []string{"echo", "type", "exit"}

	for _, c := range builtIns {
		if cmd == c {
			return true
		}
	}

	return false
}

func isExec(paths []string, cmd string) bool {
	for _, path := range paths {
		execPath := filepath.Join(path, cmd)

		if _, err := os.Stat(execPath); err == nil {
			fmt.Printf("%s is %s\n", cmd, execPath)
			return true
		}
	}
	return false
}
