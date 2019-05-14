package git

import (
	"bytes"
)

// Git Namespace for git command execution
type Git struct {
	Output   string
	Error    error
	ErrorMsg string

	directory string
	exec      Executor
}

// NewGit Constructor for git executor
func NewGit(directory string, exec Executor) Git {
	return Git{directory: directory, exec: exec}
}

func (g *Git) execGitCommand(args []string) {
	argsWithDirectory := append([]string{"-C", g.directory}, args...)
	cmd := g.exec.Command("git", argsWithDirectory...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := cmd.Run(&stdout, &stderr)

	g.setOutputAndError(stdout, err, stderr)
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
func (g *Git) Delete(branch string, force bool) error {
	deleteArg := getDeleteArg(force)
	args := []string{"branch", deleteArg, branch}
	g.execGitCommand(args)
	return g.gitErrorIfSet()
}

// Fetch executes git fetch command
func (g *Git) Fetch() error {
	g.execGitCommand([]string{"fetch"})
	return g.gitErrorIfSet()
}

// ListRemoteBranches executes git branch -vv
func (g *Git) ListRemoteBranches() error {
	args := []string{"branch", "-vv"}
	g.execGitCommand(args)
	return g.gitErrorIfSet()
}

type GitError struct {
	msg string
}

func (g *Git) gitErrorIfSet() error {
	var msg string
	if g.ErrorMsg != "" {
		msg = g.ErrorMsg
	} else if g.Error != nil {
		msg = g.Error.Error()
	}

	if msg != "" {
		return GitError{msg: msg}
	}

	return nil
}

func (e GitError) Error() string {
	return e.msg
}

// Prune executes git remote prune origin
func (g *Git) Prune() error {
	args := []string{"remote", "prune", "origin"}
	g.execGitCommand(args)
	return g.gitErrorIfSet()
}
