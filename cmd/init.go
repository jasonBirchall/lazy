package cmd

import (
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
		_ = createConfigFile(homeDir)
	},
}

func createConfigFile(homeDir string) string {
	configFilePath := filepath.Join(homeDir, ".lazycommit.yaml")
	err := os.WriteFile(configFilePath, []byte(configContent), 0644)
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
