package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Prints out the diff in the current directory",
	Long: `This command prints out the diff of the changes in the current directory.
It helps you see what changes have been made before committing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		printDiff()
	},
}

func printDiff() {
	cmd := exec.Command("git", "diff")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running git diff:", err)
		return
	}
	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
