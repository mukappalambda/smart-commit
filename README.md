# Smart Commit

Smart Commit is a command-line interface tool that leverages Large Language Models (LLMs) to generate intelligent and descriptive Git commit messages automatically.

## Features

- **Automatic Commit Message Generation**: Analyzes your staged changes and suggests a concise yet informative commit message.
- **LLM Powered**: Integrates with LLMs (e.g., OpenAI) to provide high-quality suggestions.
- **Customizable**: Easily configure your preferred LLM provider and other settings.

## Installation

To install Smart Commit, make sure you have Go installed (version 1.18 or higher).

```bash
go install
```

## Usage

Navigate to your Git repository and stage your changes as usual.

```bash
git add .
```

Then, run the `smart-commit` command to generate a commit message:

```bash
smart-commit generate
```

The generated message will be displayed, and you will be prompted to confirm or edit it before committing.

## Configuration

The configuration file is typically located at `~/.config/smart-commit/config.yaml`.

Example `config.yaml`:

```yaml
openai_api_key: your_openai_api_key_here
model: gpt-4o
custom_prompt: "My custom prompt: %s"
```

Replace `your_openai_api_key_here` with your actual OpenAI API key.

The `custom_prompt` is optional. If you provide a custom prompt, it will be used to generate the commit message. The `%s` in the prompt will be replaced with the git diff.
