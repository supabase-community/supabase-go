# Contributing to Supabase Go Client

Thank you for your interest in contributing to the Supabase Go Client! This document provides guidelines and instructions for contributing to this project.

## Getting Started

### Prerequisites

- Go 1.24.0 or later
- Git
- Python 3.6+ (for pre-commit hooks)

### Setting Up Your Development Environment

1. **Fork and clone the repository:**

   ```bash
   git clone https://github.com/YOUR_USERNAME/supabase-go.git
   cd supabase-go
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Install pre-commit hooks:**

   Pre-commit hooks help maintain code quality by running automated checks before each commit.

   First, install pre-commit:

   ```bash
   pip install pre-commit
   ```

   Then install the hooks for this repository:

   ```bash
   pre-commit install
   ```

   This will set up the following checks that run automatically before each commit:
   - **Go formatting** with `gofumpt` and `goimports`
   - **Go vetting** to catch common Go mistakes
   - **Go mod tidy** to ensure dependencies are properly managed
   - **YAML validation** for configuration files
   - **Trailing whitespace removal**
   - **End-of-file fixing**
   - **Security scanning** with Gitleaks to prevent credential leaks

4. **Run tests to verify your setup:**

   ```bash
   go test ./...
   ```

## Development Workflow

### Before You Start

1. Check existing issues to see if your bug/feature is already being tracked
2. For new features, consider opening an issue first to discuss the approach
3. **Create a new branch from `main` for your changes:**

   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

### Making Changes

1. **Write your code** following Go best practices
2. **Add or update tests** for your changes
3. **Run tests locally:**

   ```bash
   go test ./...
   ```

4. **Run pre-commit checks manually** (optional, they'll run automatically on commit):

   ```bash
   pre-commit run --all-files
   ```

### Code Style Guidelines

- Follow standard Go formatting (enforced by `gofumpt`)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and reasonably sized
- Follow Go naming conventions (CamelCase for exported, camelCase for unexported)

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting
- Include integration tests where appropriate
- Test files should be named `*_test.go`

### Commit Messages

Use clear and descriptive commit messages:

- Start with a short summary (50 characters or less)
- Use the imperative mood ("Add feature" not "Added feature")
- Reference issues when applicable ("Fixes #123")

Example:

```text
Add support for custom timeout configuration

- Add Timeout field to ClientOptions
- Update client initialization to respect timeout
- Add tests for timeout behavior

Fixes #456
```

## Pre-commit Hooks Details

Our pre-commit configuration includes several important checks:

### Go-specific hooks

- **go-vet-mod**: Runs `go vet` to catch common Go mistakes
- **go-imports**: Automatically manages import statements
- **go-fumpt**: More strict version of `go fmt` for consistent formatting
- **go-mod-tidy-repo**: Ensures `go.mod` and `go.sum` are clean

### General hooks

- **end-of-file-fixer**: Ensures files end with a newline
- **trailing-whitespace**: Removes trailing whitespace
- **check-yaml**: Validates YAML syntax
- **gitleaks**: Scans for accidentally committed secrets/credentials

### Bypassing Pre-commit Hooks

In rare cases where you need to bypass pre-commit hooks (not recommended), you can use:

```bash
git commit --no-verify -m "Your commit message"
```

## Submitting Changes

1. **Push your branch** to your fork:

   ```bash
   git push origin your-branch-name
   ```

2. **Create a Pull Request** with:
   - A clear title and description
   - Reference to any related issues
   - Description of changes made
   - Any breaking changes highlighted

3. **Respond to feedback** from maintainers and update your PR as needed

## Project Structure

- `client.go` - Main client implementation
- `client_test.go` - Client tests
- `test/` - Additional test utilities
- `.pre-commit-config.yaml` - Pre-commit hook configuration

## Integration Libraries

This client integrates with several Supabase community Go libraries:

- [postgrest-go](https://github.com/supabase-community/postgrest-go) - Database REST API
- [gotrue-go](https://github.com/supabase-community/gotrue-go) - Authentication
- [storage-go](https://github.com/supabase-community/storage-go) - File storage
- [functions-go](https://github.com/supabase-community/functions-go) - Edge functions

## Getting Help

- Check existing [GitHub Issues](https://github.com/givr-eng/supabase-go/issues)
- Create a new issue for bugs or feature requests
- Join the [Supabase Discord](https://discord.supabase.com) community

## Code of Conduct

Please be respectful and inclusive in all interactions. We aim to create a welcoming environment for all contributors.

Thank you for contributing to the Supabase Go Client! 🚀
