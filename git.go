package main

import (
	"fmt"
	"os/exec"
)

const cmd = "git"

// Git Namespace for git command execution
type Git struct {
}

// Fetch  execture git fetch command
func (g Git) Fetch() error {
	fmt.Println("Running git fetch...")
	return exec.Command(cmd, "fetch").Run()
}
