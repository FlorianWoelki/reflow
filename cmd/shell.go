package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/florianwoelki/reflow/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the reflow programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Run(os.Stdin, os.Stdout)
}
