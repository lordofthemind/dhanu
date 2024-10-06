package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Restricted file extensions as per Gmail's restrictions.
var restrictedExtensions = []string{
	".exe", ".bat", ".com", ".msi", ".cmd", ".sh", ".apk", ".js", ".vbs", ".ps1", ".jar", ".iso", ".dmg", ".zip", ".rar", ".tar", ".gz", ".docm", ".xlsm",
}

// isRestricted checks if the file has a restricted extension.
func isRestricted(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	for _, restrictedExt := range restrictedExtensions {
		if ext == restrictedExt {
			return true
		}
	}
	return false
}

// CreateZipFromAttachments takes the list of attachment paths, renames restricted files if necessary, and creates a zip file for email attachment.
// Parameters:
// - attachments: List of file paths to be attached.
// - zipFileName: Name of the resulting zip file.
// Returns the path to the zip file or an error.
func CreateZipFromAttachments(attachments []string, zipFileName string) (string, error) {
	// Create the zip file.
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return "", fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	// Initialize zip writer.
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, attachment := range attachments {
		// Open the file to be added to the zip.
		file, err := os.Open(attachment)
		if err != nil {
			return "", fmt.Errorf("failed to open attachment: %v", err)
		}
		defer file.Close()

		// Check if the file is restricted.
		_, fileName := filepath.Split(attachment)
		if isRestricted(fileName) {
			// Rename restricted file (add .safe to its name).
			fileName = fileName + ".safe"
		}

		// Create zip entry for the file.
		zipFileEntry, err := zipWriter.Create(fileName)
		if err != nil {
			return "", fmt.Errorf("failed to create zip entry for file: %v", err)
		}

		// Copy the file content to the zip entry.
		if _, err := io.Copy(zipFileEntry, file); err != nil {
			return "", fmt.Errorf("failed to copy file content to zip entry: %v", err)
		}
	}

	// Successfully created the zip.
	return zipFileName, nil
}
