package utils

import "regexp"

// Function to validate email format
func IsValidEmail(email string) bool {
	// Simple regex for validating an email format
	// This regex may not cover all edge cases, but it works for most common email formats.
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
