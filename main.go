package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/RPiAwesomeness/in.flux/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the in.Flux REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
