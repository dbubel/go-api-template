// Package cli provides a command-line interface framework with no external dependencies.
package cli

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
)

// Command is an interface that implements a CLI command.
type Command interface {
	// Help returns the help text for the command
	Help() string

	// Synopsis returns a brief description of the command
	Synopsis() string

	// Run executes the command with the given arguments
	// and returns an exit code (0 for success)
	Run(args []string) int
}

// CommandFactory is a function that creates a Command
type CommandFactory func() (Command, error)

// CLI is the main CLI struct that manages commands
type CLI struct {
	// Name is the name of the CLI application
	Name string

	// Version is the version of the CLI application
	Version string

	// Args are the command-line arguments
	Args []string

	// Commands is a map of command names to their factories
	Commands map[string]CommandFactory

	// HelpFunc is a function that displays help for a command
	HelpFunc HelpFunc

	// HelpWriter is where help text is written to
	HelpWriter io.Writer

	once sync.Once
	isHelp bool
}

// HelpFunc is a function that displays help for a command
type HelpFunc func(map[string]CommandFactory) string

// NewCLI creates a new CLI instance
func NewCLI(name, version string) *CLI {
	return &CLI{
		Name:       name,
		Version:    version,
		HelpWriter: os.Stdout,
		HelpFunc:   DefaultHelpFunc,
	}
}

// IsHelp returns whether the CLI is displaying help
func (c *CLI) IsHelp() bool {
	c.once.Do(func() {
		c.isHelp = c.isHelpRequest()
	})
	return c.isHelp
}

// isHelpRequest returns true if the CLI arguments suggest we're asking for help
func (c *CLI) isHelpRequest() bool {
	if len(c.Args) == 0 {
		return true
	}

	arg := c.Args[0]
	return arg == "-h" || arg == "--help" || arg == "help"
}

// Run executes the CLI with the arguments provided
func (c *CLI) Run() (int, error) {
	// If we're in help mode, run help
	if c.IsHelp() {
		cmd, err := c.helpCommand()
		if err != nil {
			return 1, err
		}
		return cmd.Run(c.Args), nil
	}

	// Get the command name
	cmdName := c.Args[0]

	// Check if the command exists
	factory, ok := c.Commands[cmdName]
	if !ok {
		fmt.Fprint(c.HelpWriter, "Unknown command: ")
		fmt.Fprintln(c.HelpWriter, cmdName)
		fmt.Fprintln(c.HelpWriter, "")
		fmt.Fprint(c.HelpWriter, c.HelpFunc(c.Commands))
		return 1, nil
	}

	// Create the command
	cmd, err := factory()
	if err != nil {
		fmt.Fprint(c.HelpWriter, "Error instantiating ")
		fmt.Fprint(c.HelpWriter, cmdName)
		fmt.Fprint(c.HelpWriter, ": ")
		fmt.Fprintln(c.HelpWriter, err)
		return 1, err
	}

	// Run the command
	return cmd.Run(c.Args[1:]), nil
}

// helpCommand returns a Command that shows help
func (c *CLI) helpCommand() (Command, error) {
	return &helpCommand{
		cli:    c,
		append: true,
	}, nil
}

// helpCommand is a Command that shows help
type helpCommand struct {
	cli    *CLI
	append bool
}

func (c *helpCommand) Help() string {
	return "Shows help for a command"
}

func (c *helpCommand) Synopsis() string {
	return "Shows help"
}

func (c *helpCommand) Run(args []string) int {
	// If no args, show all commands
	if len(args) <= 1 {
		fmt.Fprint(c.cli.HelpWriter, c.cli.HelpFunc(c.cli.Commands))
		return 0
	}

	// Show help for a specific command
	cmdName := args[1]
	factory, ok := c.cli.Commands[cmdName]
	if !ok {
		fmt.Fprint(c.cli.HelpWriter, "Unknown command: ")
		fmt.Fprintln(c.cli.HelpWriter, cmdName)
		fmt.Fprintln(c.cli.HelpWriter, "")
		fmt.Fprint(c.cli.HelpWriter, c.cli.HelpFunc(c.cli.Commands))
		return 1
	}

	cmd, err := factory()
	if err != nil {
		fmt.Fprint(c.cli.HelpWriter, "Error instantiating ")
		fmt.Fprint(c.cli.HelpWriter, cmdName)
		fmt.Fprint(c.cli.HelpWriter, ": ")
		fmt.Fprintln(c.cli.HelpWriter, err)
		return 1
	}

	fmt.Fprint(c.cli.HelpWriter, "Usage: ")
	fmt.Fprint(c.cli.HelpWriter, c.cli.Name)
	fmt.Fprint(c.cli.HelpWriter, " ")
	fmt.Fprintln(c.cli.HelpWriter, cmdName)
	fmt.Fprintln(c.cli.HelpWriter, "")
	fmt.Fprint(c.cli.HelpWriter, cmd.Help())
	return 0
}

// DefaultHelpFunc returns help text for all commands
func DefaultHelpFunc(commands map[string]CommandFactory) string {
	var buf strings.Builder

	// Get a list of all command names
	commandNames := make([]string, 0, len(commands))
	for name := range commands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	buf.WriteString("Available commands:\n\n")

	// For each command, add its synopsis
	for _, name := range commandNames {
		factory := commands[name]
		cmd, err := factory()
		if err != nil {
			buf.WriteString("  ")
			buf.WriteString(name)
			buf.WriteString("\n")
			continue
		}
		buf.WriteString("  ")
		buf.WriteString(name)
		buf.WriteString(" - ")
		buf.WriteString(cmd.Synopsis())
		buf.WriteString("\n")
	}

	return buf.String()
} 