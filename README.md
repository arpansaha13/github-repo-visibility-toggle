# GitHub Repo Visibility Toggle (Go Script)

This is a lightweight Go script that lets you **bulk update the visibility** of your GitHub repositories, switching them between `public` or `private`, using the GitHub REST API.

It's ideal for developers who want to automate visibility changes across multiple repositories without manually editing each one.

---

## 🚀 Quick Start

### 1. Clone the Repo

```bash
git clone https://github.com/arpansaha13/github-repo-visibility-toggle.git

cd github-repo-visibility-toggle
```

### 2. Set Your GitHub Token

- Generate a [fine-grained personal access token](https://github.com/settings/personal-access-tokens) with repository scope.
- Store it in a `.env` file beside the script.

```env
# .env

GITHUB_TOKEN=your_personal_access_token_here
```

### 3. Run the Script

```bash
go run update_visibility.go <visibility> <owner/repo> [owner/repo] ...
```

#### Example

```bash
go run update_visibility.go public username/my-repo username/another-repo
```

## How It Works

- Uses the GitHub REST API `PATCH` `/repos/{owner}/{repo}` endpoint.
- Sends a JSON payload with the desired visibility.
- Authenticates using your personal access token.

## Notes

- Valid visibility values: `public` or `private`.
- Make sure your token has permission to edit the target repositories.

## License

[MIT License](https://github.com/arpansaha13/github-repo-visibility-toggle/blob/main/LICENSE) - Made with ❤️
