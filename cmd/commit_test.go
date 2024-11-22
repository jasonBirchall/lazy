package cmd

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestCommitCmd(t *testing.T) {
	// Mock the exec.Command function
	originalCommand := execCommand
	defer func() { execCommand = originalCommand }()
	execCommand = mockExecCommand

	// Capture the output
	var output bytes.Buffer
	rootCmd.SetOut(&output)

	// Execute the commit command
	rootCmd.SetArgs([]string{"commit"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check the output
	expectedOutput := "mocked git diff output\n"
	if output.String() != expectedOutput {
		t.Fatalf("Expected output to be %q, got %q", expectedOutput, output.String())
	}
}

// Mock exec.Command function
var execCommand = exec.Command

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cmd := exec.Command("echo", "mocked git diff output")
	return cmd
}
