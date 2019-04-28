package git

import (
	"bytes"
	"os/exec"
)

//go:generate moq -out executor_moq_test.go . Executor
// Executor - exposes Command function
type Executor interface {
	Command(name string, args ...string) CmdRunner
}

//go:generate moq -out cmd_runnerr_moq_test.go . CmdRunner
// CmdRunner created by Command(), exposes Run function
type CmdRunner interface {
	Run(stdOut *bytes.Buffer, stdErr *bytes.Buffer) error
}

// NewExecutorWithExec Created executor using os/exec
func NewExecutorWithExec() ExecWrapper {
	return ExecWrapper{}
}

// ExecWrapper Wrapper on exec that conforms to generic Executor
type ExecWrapper struct{}

// Command Calls exec.Command and returns a wrapped *exec.Cmd
func (w ExecWrapper) Command(name string, args ...string) CmdRunner {
	return cmdWrapper{cmd: exec.Command(name, args...)}
}

type cmdWrapper struct {
	cmd *exec.Cmd
}

func (w cmdWrapper) Run(stdOut *bytes.Buffer, stdErr *bytes.Buffer) error {
	w.cmd.Stdout = stdOut
	w.cmd.Stderr = stdErr
	return w.cmd.Run()
}
