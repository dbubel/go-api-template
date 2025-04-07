package cli

import (
	"bytes"
	"testing"
)

// TestCLI_Run tests the basic functionality of the CLI
func TestCLI_Run(t *testing.T) {
	// Create a test command
	testCmd := &testCommand{
		SynopsisText: "test command",
		HelpText:     "This is a test command",
		RunResult:    42,
	}

	// Create a CLI
	buf := new(bytes.Buffer)
	cli := &CLI{
		Args: []string{"test"},
		Commands: map[string]CommandFactory{
			"test": func() (Command, error) {
				return testCmd, nil
			},
		},
		HelpWriter: buf,
	}

	// Run the CLI
	exitCode, err := cli.Run()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Check the exit code
	if exitCode != 42 {
		t.Fatalf("expected exit code 42, got: %d", exitCode)
	}

	// Check the command was run
	if !testCmd.WasRun {
		t.Fatal("command was not run")
	}
}

// TestCLI_Help tests that help works correctly
func TestCLI_Help(t *testing.T) {
	// Create a test command
	testCmd := &testCommand{
		SynopsisText: "test command",
		HelpText:     "This is a test command",
	}

	// Create a CLI
	buf := new(bytes.Buffer)
	cli := &CLI{
		Args: []string{"help"},
		Commands: map[string]CommandFactory{
			"test": func() (Command, error) {
				return testCmd, nil
			},
		},
		HelpWriter: buf,
		HelpFunc:   BasicHelpFunc("test-app"),
	}

	// Run the CLI
	exitCode, err := cli.Run()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Check the exit code
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got: %d", exitCode)
	}

	// Check that help was output
	if buf.Len() == 0 {
		t.Fatal("expected help output, got none")
	}
}

// BasicHelpFunc returns a HelpFunc that outputs the command's help text
func BasicHelpFunc(app string) HelpFunc {
	return func(commands map[string]CommandFactory) string {
		return DefaultHelpFunc(commands)
	}
}

// testCommand is a test implementation of Command
type testCommand struct {
	SynopsisText string
	HelpText     string
	RunResult    int
	WasRun       bool
}

func (c *testCommand) Help() string {
	return c.HelpText
}

func (c *testCommand) Synopsis() string {
	return c.SynopsisText
}

func (c *testCommand) Run(args []string) int {
	c.WasRun = true
	return c.RunResult
} 