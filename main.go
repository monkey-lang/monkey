package main

import (
	"fmt"
	"github.com/monkey-lang/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type commands")
	repl.Start(os.Stdin, os.Stdout)
}
