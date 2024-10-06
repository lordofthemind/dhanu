package utils

import (
	"fmt"
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

// HandleAttachments checks the list of attachment paths for restricted files and denies sending them.
// Parameters:
// - attachments: List of file paths to be checked.
// Returns an error if any attachment is restricted.
func HandleAttachments(attachments []string) error {
	for _, attachment := range attachments {
		// Check if the file is restricted.
		if isRestricted(attachment) {
			return fmt.Errorf("attachment '%s' is restricted and cannot be sent", attachment)
		}
	}

	// All attachments are valid.
	return nil
}
