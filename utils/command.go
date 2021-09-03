package utils

import (
	"errors"
	"fmt"
)

// CommandRunner is a function signature expected
// for modules that need to register new commands
type CommandRunner func(args ...string) error

// CommandOptions allows services to registers
// commands available in the terminal
type CommandOptions struct {
	Cmd    string
	Runner CommandRunner
}

// cli contains all registered executable commands
type cli struct {
	Commands map[string]CommandRunner
}

// Service is used to manage cli commands
var Service *cli

// RunCommand processes commands executed in terminal
// and the args is a slice of arguments passed in terminal
// where the first element is the CommandOptions.Cmd which
// will
func (c *cli) RunCommand(args ...string) error {

	if len(args) == 0 {
		return errors.New("command is required")
	}

	cmd := args[0]
	runner := c.Commands[cmd]
	if runner == nil {
		return errors.New(fmt.Sprintf(
			"invalid command %s",
			cmd,
		))
	}

	return runner(args[1:]...)
}

// RegisterCommand allows each service to create
// an executable command which can be triggered
// in terminal
func (c *cli) RegisterCommand(command CommandOptions) error {

	if c.Commands[command.Cmd] != nil {
		return errors.New(fmt.Sprintf(
			"comamnd already registered: [command: %s]",
			command.Cmd,
		))
	}

	c.Commands[command.Cmd] = command.Runner

	return nil
}

// init instantiates command Service
func init() {
	Service = &cli{
		Commands: make(map[string]CommandRunner),
	}
}
