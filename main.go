package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

// flags
var (
	force bool
)

func checkForError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	git := Git{}

	checkForError(git.Fetch())

	fmt.Println("Done.")
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
}
