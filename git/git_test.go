package git_test

import (
	"bytes"
	g "github.com/drewwyatt/gitclean/git"
	"testing"
)

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
		t.Errorf("Send was called %d times", callsToCommand)
	}
	if cmdName != "fetch" {
		t.Errorf("unexpected command name: %s", cmdName)
	}
}
