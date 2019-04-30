package git_test

import (
	"bytes"
	"testing"

	g "github.com/drewwyatt/gitclean/git"
)

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

	git.Delete(branchName, true)
	if deletionArg != "-D" {
		t.Errorf("unexpected deletion arg: %s (expected '-D)", deletionArg)
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
