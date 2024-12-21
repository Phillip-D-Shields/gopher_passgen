package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	specials  = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func generatePassword(length int, useLower, useUpper, useNumbers, useSpecial bool) (string, error) {
	var chars string
	if useLower {
		chars += lowercase
	}
	if useUpper {
		chars += uppercase
	}
	if useNumbers {
		chars += numbers
	}
	if useSpecial {
		chars += specials
	}

	if chars == "" {
		return "", fmt.Errorf("no character sets selected")
	}

	password := make([]byte, length)
	max := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}
		password[i] = chars[num.Int64()]
	}

	return string(password), nil
}

func main() {
	length := flag.Int("length", 16, "Password length")
	noLower := flag.Bool("no-lower", false, "Exclude lowercase letters")
	noUpper := flag.Bool("no-upper", false, "Exclude uppercase letters")
	noNumbers := flag.Bool("no-numbers", false, "Exclude numbers")
	noSpecial := flag.Bool("no-special", false, "Exclude special characters")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	// Generate specified number of passwords
	for i := 0; i < *count; i++ {
		password, err := generatePassword(*length, !*noLower, !*noUpper, !*noNumbers, !*noSpecial)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(password)
	}
}
