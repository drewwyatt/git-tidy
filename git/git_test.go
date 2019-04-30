package git_test

import (
	"bytes"
	"testing"

	g "github.com/drewwyatt/gitclean/git"
)

var errorString string = "This is an error"

type TestError struct{}

func (e TestError) Error() string {
	return errorString
}

func TestDelete(t *testing.T) {
	var cmdName string
	var deletedBranch string
	var deletionArg string
	branchName := "fix-the-bug"
	mockedExecutor := &g.ExecutorMock{
		CommandFunc: func(name string, otherArgs ...string) g.CmdRunner {
			cmdName = otherArgs[2]
			deletionArg = otherArgs[3]
			deletedBranch = otherArgs[4]
			return &g.CmdRunnerMock{
				RunFunc: func(o *bytes.Buffer, e *bytes.Buffer) error {
					return nil
				},
			}
		},
	}
	git := g.NewGit(".", mockedExecutor)
	git.Delete(branchName, false)
	callsToCommand := len(mockedExecutor.CommandCalls())

	if callsToCommand != 1 {
		t.Errorf("cmd was called %d times", callsToCommand)
	}
	if cmdName != "branch" {
		t.Errorf("unexpected command name: %s (expected 'branch')", cmdName)
	}
	if deletedBranch != branchName {
		t.Errorf("unexpected branch name: %s (expected %s)", deletedBranch, branchName)
	}
	if deletionArg != "-d" {
		t.Errorf("unexpected deletion arg: %s (expected '-d')", deletionArg)
	}

	if len(git.DeletedBranches) != 1 {
		t.Errorf("unexpected number of deleted branches: %d", len(git.DeletedBranches))
	}

	if git.DeletedBranches[0] != branchName {
		t.Errorf("Unexpected branch name: %s", git.DeletedBranches[0])
	}

	git.Delete(branchName, true)
	if deletionArg != "-D" {
		t.Errorf("unexpected deletion arg: %s (expected '-D')", deletionArg)
	}

	if len(git.DeletedBranches) != 2 {
		t.Errorf("unexpected number of deleted branches: %d", len(git.DeletedBranches))
	}
}

func TestDeleteError(t *testing.T) {
	mockedExecutor := &g.ExecutorMock{
		CommandFunc: func(name string, otherArgs ...string) g.CmdRunner {
			return &g.CmdRunnerMock{
				RunFunc: func(o *bytes.Buffer, e *bytes.Buffer) error {
					return TestError{}
				},
			}
		},
	}

	git := g.NewGit(".", mockedExecutor)
	branchName := "hello"
	git.Delete(branchName, false)

	if git.Error.Error() != errorString {
		t.Errorf("Unexpected error message: %s", git.Error)
	}

	if len(git.BranchDeletionErrors) != 1 {
		t.Errorf("Unexpected error count: %d", len(git.BranchDeletionErrors))
	}

	if git.BranchDeletionErrors[0].Branch != branchName {
		t.Errorf("Unexpected branch name: %s", git.BranchDeletionErrors[0].Branch)
	}
}

func TestFetch(t *testing.T) {
	var cmdName string
	mockedExecutor := &g.ExecutorMock{
		CommandFunc: func(name string, otherArgs ...string) g.CmdRunner {
			cmdName = otherArgs[2]
			return &g.CmdRunnerMock{
				RunFunc: func(o *bytes.Buffer, e *bytes.Buffer) error {
					return nil
				},
			}
		},
	}
	git := g.NewGit(".", mockedExecutor)
	git.Fetch()
	callsToCommand := len(mockedExecutor.CommandCalls())
	if callsToCommand != 1 {
		t.Errorf("cmd was called %d times", callsToCommand)
	}
	if cmdName != "fetch" {
		t.Errorf("unexpected command name: %s", cmdName)
	}
}

func TestListRemoteBranches(t *testing.T) {
	var cmdName string
	var arg string
	mockedExecutor := &g.ExecutorMock{
		CommandFunc: func(name string, otherArgs ...string) g.CmdRunner {
			cmdName = otherArgs[2]
			arg = otherArgs[3]
			return &g.CmdRunnerMock{
				RunFunc: func(o *bytes.Buffer, e *bytes.Buffer) error {
					return nil
				},
			}
		},
	}
	git := g.NewGit(".", mockedExecutor)
	git.ListRemoteBranches()
	callsToCommand := len(mockedExecutor.CommandCalls())
	if callsToCommand != 1 {
		t.Errorf("cmd was called %d times", callsToCommand)
	}
	if cmdName != "branch" {
		t.Errorf("unexpected command name: %s", cmdName)
	}

	if arg != "-vv" {
		t.Errorf("unexpected arg: %s", arg)
	}
}
