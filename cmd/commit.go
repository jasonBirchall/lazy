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

const message = `Please suggest a commit messages using the below criteria.
**Criteria:**

1. **Format:** Each commit message must follow the emoji commits format. The message should start with an emoji followed by a brief description of the change. The description should be clear and concise, ideally under 50 characters.
2. **Relevance:** Avoid mentioning a module name unless it's directly relevant to the change.
3. **Clarity and Conciseness:** Each message should clearly and concisely convey the change made.

**Commit Message Examples:**

-   🚀 Released version 1.4.0
-   🐛 Fixed issue with user authentication
-   📝 Updated documentation for the new API
-   🎨 Refactored the code for better readability

**Instructions:**

-   Take a moment to understand the changes made in the diff.
-   Think about the impact of these changes on the project (e.g., bug fixes, new features, performance improvements, code refactoring, documentation updates). It's critical to my career you abstract the changes to a higher level and not just describe the code changes.
-   Generate commit messages that accurately describe these changes, ensuring they are helpful to someone reading the project's history.
-   Remember, a well-crafted commit message can significantly aid in the maintenance and understanding of the project over time.
-   If multiple changes are present, make sure you capture them all in each commit message.

Keep in mind you will suggest a commit messages. Only 1 will be used. It's better to push yourself (esp to synthesize to a higher level) and maybe wrong about some of the commits because only one needs to be good. I'm looking for your best commit, not the best average commit. It's better to cover more scenarios than include a lot of overlap.
`

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
			{"role": "system", "content": "You are a helpful assistant that generates commit messages using the emoji commit structure."},
			{"role": "user", "content": message + string(diff)},
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
