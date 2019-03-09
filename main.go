package main

import (
	"fmt"
	"os"
	"os/exec"

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

	cmd := "git"
	args := []string{"status"}

	out, err := exec.Command(cmd, args...).Output()
	checkForError(err)

	// Can I combine these?
	formattedOutput := fmt.Sprintf("%s", out)
	fmt.Println(formattedOutput)

	fmt.Println("Done.")
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
}
