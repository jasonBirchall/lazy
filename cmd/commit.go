package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var openAIAPIKey = os.Getenv("OPENAI_TOKEN")

const systemMessage = `You are intelligent, helpful and an expert developer, who always gives the correct answer and only does what instructed. You always answer truthfully and don't make things up. (When responding to the following prompt, please make sure to properly style your response using Github Flavored Markdown. Use markdown syntax for things like headings, lists, colored text, code blocks, highlights etc. Make sure not to mention markdown or styling in your actual response.)`

const userMessage = `Suggest a precise and informative commit message based on the following diff. Do not use markdown syntax in your response.

The commit message should have description with a short title that follows emoji commit message format like <emoji> <description>.

Examples:
- :refactor: Change log format for better visibility
- :sparkles: Introduce new logging class

Diff: `

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Prints out the diff in the current directory and suggests commit messages",
	Long: `This command prints out the diff of the changes in the current directory.
It helps you see what changes have been made before committing them. It also suggests commit messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		diff := gitDiff()
		if diff == nil {
			return
		}

		commitMessages := getCommitMessages(diff)
		if len(commitMessages) == 0 {
			fmt.Println("No commit messages generated.")
			return
		}

		prompt := promptui.Select{
			Label: "Select a commit message",
			Items: commitMessages,
		}

		_, selectedMessage, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		commitChanges(selectedMessage)
	},
}

func gitDiff() []byte {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running git diff:", err)
		return nil
	}
	return output
}

func getCommitMessages(diff []byte) []string {
	// Call OpenAI API to get commit messages using gpt-3.5-turbo
	url := "https://api.openai.com/v1/chat/completions"
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": systemMessage},
			{"role": "user", "content": userMessage + string(diff)},
		},
		"max_tokens": 150,
		"n":          5,
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	choices, ok := result["choices"].([]interface{})
	if !ok {
		fmt.Println("Error parsing response.")
		return nil
	}

	var messages []string
	for _, choice := range choices {
		text, ok := choice.(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
		if ok {
			messages = append(messages, text)
		}
	}

	return messages
}

func commitChanges(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}
	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
