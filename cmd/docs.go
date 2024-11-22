package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generates Markdown documentation for the CLI",
	Long: `This command generates Markdown documentation for the CLI commands
and saves it in the docs directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDocs()
	},
}

func generateDocs() {
	dir := "docs"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error generating docs:", err)
		return
	}

	err = doc.GenMarkdownTree(rootCmd, dir)
	if err != nil {
		fmt.Println("Error generating documentation:", err)
		return
	}

	fmt.Println("Documentation generated in the docs directory")
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
