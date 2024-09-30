package main

import (
	"dumblang/dumblang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the DumbLang programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	// Start the REPL
	repl.Start(os.Stdin, os.Stdout)

}
