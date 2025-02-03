package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input", err)
		}

		user_cmd := command[:len(command)-1]

		cmds := strings.Fields(user_cmd)

		pathEnv := os.Getenv("PATH")
		paths := strings.Split(pathEnv, ":")

		switch cmds[0] {
		case "exit":
			if cmds[1] == "0" {
				os.Exit(0)
			}
		case "echo":
			fmt.Println(strings.Join(cmds[1:], " "))
		case "type":
			if cmds[1] == "type" || cmds[1] == "exit" || cmds[1] == "echo" {
				fmt.Printf("%s is a shell builtin\n", cmds[1])
			} else {
				found := false
				for _, path := range paths {
					exec := filepath.Join(path, cmds[1])

					if _, err := os.Stat(exec); err == nil {
						fmt.Printf("%s is %s\n", cmds[1], exec)
						found = true
					}
				}

				if !found {

					fmt.Printf("%s not found\n", cmds[1])
				}
			}
		default:
			fmt.Println(user_cmd + ": command not found")
		}
	}
}
