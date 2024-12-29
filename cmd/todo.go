package cmd

import (
	"log"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "A CLI todo app navigable with Vim commands",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		list := tview.NewList()

		// Sample todo items
		todoItems := []string{"Buy groceries", "Read a book", "Write code"}

		for _, item := range todoItems {
			list.AddItem(item, "", 0, nil)
		}

		list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			list.SetItemText(index, mainText+" [x]", secondaryText)
		})

		list.SetDoneFunc(func() {
			app.Stop()
		})

		if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
			log.Fatalf("Error running application: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
