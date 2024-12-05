package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configContent = `
# Configuration file for lazycommit
# Add your configuration settings here
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the configuration file in the user's Linux profile",
	Long: `This command initializes a configuration file named .lazycommit.yaml
in the user's home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		configFilePath := filepath.Join(homeDir, ".lazycommit.yaml")
		if !checkForExistingConfig(configFilePath) {
			return
		}
		apiKey := promptForAPIKey()
		_ = createConfigFile(configFilePath, apiKey)
	},
}

func promptForAPIKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your OpenAI API Key: ")
	apiKey, _ := reader.ReadString('\n')
	return apiKey
}

func promptForOverwrite() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Configuration file already exists. Overwrite? (y/n): ")
	overwrite, _ := reader.ReadString('\n')
	return overwrite == "y\n"
}

func checkForExistingConfig(configFilePath string) bool {
	if _, err := os.Stat(configFilePath); err == nil {
		if !promptForOverwrite() {
			fmt.Println("Configuration file not overwritten.")
			return false
		}
	}
	return true
}

func createConfigFile(configFilePath, apiKey string) string {
	configWithAPIKey := configContent + "\nopenai_api_key: " + apiKey
	err := os.WriteFile(configFilePath, []byte(configWithAPIKey), 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		return ""
	}
	fmt.Println("Configuration file created at:", configFilePath)
	return configFilePath
}

func init() {
	rootCmd.AddCommand(initCmd)
}
