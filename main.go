package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/RPiAwesomeness/in.flux/repl"
)

const version = "v0.0.1 beta"

var greeting = `
╔══════════════════════════════════╗
║   _           ___ _              ║
║  (_)_ __     / __\ |_   ___  __  ║
║  | | '_ \   / _\ | | | | \ \/ /  ║
║  | | | | |_/ /   | | |_| |>  <   ║
║  |_|_| |_(_)/    |_|\__,_/_/\_\  ║
║                                  ║
║%s║
║%s║
╚══════════════════════════════════╝
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(greeting, center(34, version), center(34, "Welcome", user.Username))
	repl.Start(os.Stdin, os.Stdout)
}

func center(length int, value ...string) string {
	var lPad, rPad string
	// Can't center if what we're supposed to be centering is longer than the available space
	if len(value) > length {
		return ""
	}
	var valLen = len(strings.Join(value, " "))
	var available = length - valLen
	if available%2 != 0 {
		// it's an odd number, pad the left side more
		lPad += " "
		available--
	}
	available /= 2
	for i := 0; i < available; i++ {
		lPad += " "
		rPad += " "
	}

	return fmt.Sprint(lPad + strings.Join(value, " ") + rPad)
}
