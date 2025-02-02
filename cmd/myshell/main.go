package main

import (
	"bufio"
	"fmt"
	"os"
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

    if user_cmd == "exit 0" {
      os.Exit(0)
    }

		fmt.Println(user_cmd + ": command not found")
	}
}
