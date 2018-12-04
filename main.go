package main

import (
	"fmt"
	"os"
	"strings"

	bip39 "github.com/tyler-smith/go-bip39"
)

func main() {
	seed := strings.Join(os.Args[1:], " ")

	fmt.Printf("seed : %s\n", seed)

	v := bip39.IsMnemonicValid(seed)

	result := "invalid"

	if v {
		result = "valid"
	}

	fmt.Printf("phrase is %s\n", result)

	if !v {
		os.Exit(1)
	}
}
