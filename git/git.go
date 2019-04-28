package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

type branchDeletionError struct {
	Branch string
	Msg    string
}

// Git Namespace for git command execution
type Git struct {
	Output   string
	Error    error
	ErrorMsg string

	DeletedBranches      []string
	BranchDeletionErrors []branchDeletionError

	directory string
}

// NewExecutor Constructor for git executor
func NewExecutor(directory string) Git {
	return Git{directory: directory}
}

func reportProcess(name string) {
	fmt.Printf("Running '%s'...\n", name)
}

func (g *Git) execGitCommand(args []string) {
	argsWithDirectory := append([]string{"-C", g.directory}, args...)
	cmd := exec.Command("git", argsWithDirectory...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	g.setOutputAndError(out, err, stderr)
}

func (g *Git) setOutputAndError(output bytes.Buffer, error error, errMsg bytes.Buffer) {
	g.Output = output.String()
	g.Error = error
	g.ErrorMsg = errMsg.String()
}

func getDeleteArg(force bool) string {
	if force {
		return "-D"
	}

	return "-d"
}

// Delete executes git branch -d/D on a given branch
func (g *Git) Delete(branch string, force bool) *Git {
	deleteArg := getDeleteArg(force)
	reportProcess(fmt.Sprintf("git branch %s %s", deleteArg, branch))
	args := []string{"branch", deleteArg, branch}
	g.execGitCommand(args)
	if g.Error != nil {
		errs := append(g.BranchDeletionErrors, branchDeletionError{branch, g.ErrorMsg})
		g.BranchDeletionErrors = errs
	} else {
		g.DeletedBranches = append(g.DeletedBranches, branch)
	}
	return g
}

// Fetch executes git fetch command
func (g *Git) Fetch() *Git {
	reportProcess("git fetch")
	g.execGitCommand([]string{"fetch"})
	return g
}

// ListRemoteBranches executes git branch -vv
func (g *Git) ListRemoteBranches() *Git {
	reportProcess("git branch -vv")
	args := []string{"branch", "-vv"}
	g.execGitCommand(args)
	return g
}

// Prune executes git remote prune origin
func (g *Git) Prune() *Git {
	reportProcess("git remote prune origin")
	args := []string{"remote", "prune", "origin"}
	g.execGitCommand(args)
	return g
}
