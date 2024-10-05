package unzipandrename

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// UnzipAndRename takes a zip file, extracts it, and renames any files with a ".rename" extension by removing that extension.
func UnzipAndRename(zipFile string, destDir string) error {
	// Open the zip file
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	// Iterate through each file in the zip archive
	for _, file := range zipReader.File {
		filePath := filepath.Join(destDir, file.Name)

		// Check if it's a directory or file
		if file.FileInfo().IsDir() {
			// Create the directory
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		} else {
			// Ensure the directory exists for the file
			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return fmt.Errorf("failed to create file directory: %v", err)
			}

			// Open the file inside the zip
			zipFileReader, err := file.Open()
			if err != nil {
				return fmt.Errorf("failed to open zip file content: %v", err)
			}
			defer zipFileReader.Close()

			// Create the destination file
			outFile, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create destination file: %v", err)
			}
			defer outFile.Close()

			// Copy the content to the destination file
			if _, err := io.Copy(outFile, zipFileReader); err != nil {
				return fmt.Errorf("failed to copy file content: %v", err)
			}
		}

		// Check for .rename extension and remove it
		if strings.HasSuffix(filePath, ".rename") {
			newFilePath := strings.TrimSuffix(filePath, ".rename")
			if err := os.Rename(filePath, newFilePath); err != nil {
				return fmt.Errorf("failed to rename file: %v", err)
			}
		}
	}

	return nil
}
