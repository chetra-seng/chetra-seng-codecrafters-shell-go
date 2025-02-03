package main

import (
	"bufio"
	"fmt"
	"os"
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

		switch cmds[0] {
		case "exit":
			if cmds[1] == "0" {
				os.Exit(0)
			}
		case "echo":
			fmt.Println(cmds[1:])
		default:
			fmt.Println(user_cmd + ": command not found")
		}
	}
}
