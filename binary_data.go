package main

import (
	"fmt"
	"os"
)

// printBinaryData prints data about the binary
func printBinaryData() error {

	// print current directory path
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("binary path: %s\n", path)

	// print commit
	fmt.Printf("commit version: %s\n", gitCommit)

	return nil
}
