package cmd

import (
	"os"
	"testing"
)

func TestInitCmd(t *testing.T) {
	// Create a temporary directory to simulate the user's home directory
	tempDir := t.TempDir()

	// Execute the createConfigFile function
	configFilePath := createConfigFile(tempDir)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected configuration file to be created at %s", configFilePath)
	}

	// Check the content of the configuration file
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		t.Fatalf("Unexpected error reading config file: %v", err)
	}
	if string(content) != configContent {
		t.Fatalf("Expected config content to be %q, got %q", configContent, string(content))
	}
}
