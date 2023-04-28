package filecopy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFiles copies files from sourceDir to targetDir that match fileMask
func CopyFiles(sourceDir, fileMask, targetDir string) error {
	// Validate sourceDir path
	if err := validatePath(sourceDir); err != nil {
		return err
	}

	// Validate targetDir path
	if err := validatePath(targetDir); err != nil {
		return err
	}

	// Walk through the file tree starting at sourceDir and execute the function for each file and directory.
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// Use filepath.Match to match the file name against the file mask
		matched, err := filepath.Match(fileMask, info.Name())
		if err != nil {
			return err
		}

		// If the file matches the file mask and is not a directory
		if !info.IsDir() && matched {
			// Create target path for copy of file
			targetPath := filepath.Join(targetDir, info.Name())

			// Copy file
			if err := copyFile(path, targetPath); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// copyFile copies the file at sourcePath to targetPath. It returns an error if it fails.
func copyFile(sourcePath, targetPath string) error {

	// Open source file.
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create target file.
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	// Copy contents of source to target.
	_, err = io.Copy(target, source)
	if err != nil {
		return err
	}

	return nil
}

// validatePath checks if the input path is a valid directory path.
// It returns an error if the path is empty or not a directory.
func validatePath(path string) error {
	// Get file info of the path.
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("invalid path: %s", path)
	}

	// Check if the path is a directory
	if !fileInfo.IsDir() {
		return fmt.Errorf("path is not a directory: %s", path)
	}

	// Return nil if the path is valid.
	return nil
}
