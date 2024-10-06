package services

// DhanuEmailServiceInterface defines the interface for sending Dhanu emails.
type DhanuEmailServiceInterface interface {
	// SendDhanuEmail sends a plain text or HTML email based on the isHTML flag.
	SendDhanuEmail(to []string, subject, body string, isHTML bool) error

	// SendDhanuEmailWithAttachments sends an email with or without HTML and includes attachments.
	SendDhanuEmailWithAttachments(to []string, subject, body string, isHTML bool, attachments []string) error
}
