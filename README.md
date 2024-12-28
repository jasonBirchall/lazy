# Lazy CLI

Lazy is a simple and intelligent CLI tool designed to make your life easier by automating repetitive tasks. It uses OpenAI's GPT-3.5-turbo to provide commit message suggestions based on your staged changes. `lazy commit` is the first tool in the Lazy suite, with more tools to come in the future.

## Features

- **Automated Commit Message Suggestions**: Uses OpenAI's GPT-3.5-turbo to suggest commit messages based on your staged changes.
- **Emoji Commit Format**: Provides commit messages in a format that includes emojis for better visibility and categorization.
- **Interactive Selection**: Allows you to choose from multiple suggested commit messages using a simple interactive prompt.

## Installation

To install LazyCommit, use the following command:

```sh
go install github.com/jasonbirchall/lazy@latest
```

## Usage

1. **Set Up API Key**: Ensure you have an OpenAI API key. You can set it as an environment variable:

   ```sh
   export OPENAI_TOKEN=your_openai_api_key
   ```

   Alternatively, you can store it in a configuration file at `~/.lazycommit.yaml`:
   _you can run `lazy init` to create the configuration file_

   ```yaml
   openai_api_key: your_openai_api_key
   ```

2. **Stage Your Changes**: Use `git add` to stage the changes you want to commit.

3. **Run LazyCommit**: Execute the following command to get commit message suggestions:

   ```sh
   lazy commit
   ```

4. **Select a Commit Message**: Choose a commit message from the suggestions provided.

5. **Commit Your Changes**: LazyCommit will automatically commit your changes with the selected message.

## Example

Here's an example of how to use LazyCommit:

```sh
git add .
lazy commit

# Output:
# > 🚀 Added a new feature to improve user experience
#   🐛 Fixed a bug that caused the application to crash
#   📝 Updated the documentation to reflect the latest changes
#   🎨 Refactored the code to improve readability
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.