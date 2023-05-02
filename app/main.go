package main

import (
	"fmt"
	"log"

	"github.com/foxiam/filecopy/pkg/filecopy"
)

func inputWithMessage(message string) string {
	fmt.Print(message)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func main() {
	sourceDir := inputWithMessage("Enter source directory: ")
	fileMask := inputWithMessage("Enter file mask: ")
	targetDir := inputWithMessage("Enter target directory: ")

	err := filecopy.CopyFiles(sourceDir, fileMask, targetDir)

	if err != nil {
		log.Fatal(err)
	}
}
