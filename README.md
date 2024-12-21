# Password Generator CLI Tool

A secure command-line password generator written in Go. This tool generates cryptographically secure random passwords with customizable length and character sets.

## Features

- Cryptographically secure random password generation using Go's `crypto/rand`
- Customizable password length
- Flexible character set options (lowercase, uppercase, numbers, special characters)
- Generate multiple passwords at once
- Command-line flags for easy customization

## Installation

### Prerequisites

- Go 1.23 or higher
- Git (optional)

### Installing from source

1. Clone the repository (or create a new directory with the provided `main.go`):
```bash
git clone https://github.com/yourusername/passgen
cd passgen
```

2. Install the binary:
```bash
go install
```

3. Add Go's bin directory to your PATH if you haven't already:

For bash/zsh (add to `~/.bashrc` or `~/.zshrc`):
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Then reload your shell configuration:
```bash
source ~/.bashrc  # or source ~/.zshrc
```

## Usage

Basic usage:
```bash
passgen
```

This generates a 16-character password with all character sets enabled.

### Command-line Options

```bash
passgen [flags]

Flags:
  -length int
        Password length (default 16)
  -no-lower
        Exclude lowercase letters
  -no-upper
        Exclude uppercase letters
  -no-numbers
        Exclude numbers
  -no-special
        Exclude special characters
  -count int
        Number of passwords to generate (default 1)
```

### Examples

Generate a 20-character password:
```bash
passgen -length 20
```

Generate a password without special characters:
```bash
passgen -no-special
```

Generate 5 passwords:
```bash
passgen -count 5
```

Generate a numbers-only password:
```bash
passgen -no-lower -no-upper -no-special
```

Generate a 12-character password with only letters and numbers:
```bash
passgen -length 12 -no-special
```

## Security

This tool uses Go's `crypto/rand` package for generating random numbers, which is cryptographically secure. The character selection process ensures uniform distribution across the selected character sets.

## Character Sets

- Lowercase letters: a-z
- Uppercase letters: A-Z
- Numbers: 0-9
- Special characters: !@#$%^&*()_+-=[]{}|;:,.<>?

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Acknowledgments

- Built with Go's standard library
- Inspired by the need for secure password generation
