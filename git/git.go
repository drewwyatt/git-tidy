package git

import (
	"fmt"
	"os/exec"
)

const cmd = "git"

// Git Namespace for git command execution
type Git struct {
	Output    string
	RawOutput []byte
	Error     error
}

func reportProcess(name string) {
	fmt.Printf("Running '%s'...\n", name)
}

func (g *Git) setOutputAndError(output []byte, error error) {
	g.RawOutput = output
	g.Output = string(output)
	g.Error = error
}

// Fetch execture git fetch command
func (g *Git) Fetch() *Git {
	reportProcess("git fetch")
	g.setOutputAndError(exec.Command(cmd, "fetch").Output())
	return g
}

// ListRemoteBranches executes git branch -vv
func (g *Git) ListRemoteBranches() *Git {
	reportProcess("git branch -vv")
	args := []string{"branch", "-vv"}
	g.setOutputAndError(exec.Command(cmd, args...).Output())
	return g
}

// Prune executes git remote prune origin
func (g *Git) Prune() *Git {
	reportProcess("git remote prune origin")
	args := []string{"remote", "prune", "origin"}
	g.setOutputAndError(exec.Command(cmd, args...).Output())
	return g
}
