package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

// DhanuEmailService is responsible for handling email sending with various functionalities
// such as sending plain text, HTML, attachments, and more.
type DhanuEmailService struct {
	smtpHost    string
	smtpPort    string
	fromEmail   string
	credentials string
}

// NewDhanuEmailService creates a new instance of DhanuEmailService with the given SMTP configurations.
// Parameters:
// - smtpHost: The host of the SMTP server.
// - smtpPort: The port of the SMTP server.
// - fromEmail: The sender's email address.
// - credentials: The sender's credentials (e.g., password or app-specific token).
func NewDhanuEmailService(smtpHost, smtpPort, fromEmail, credentials string) DhanuEmailServiceInterface {
	return &DhanuEmailService{
		smtpHost:    smtpHost,
		smtpPort:    smtpPort,
		fromEmail:   fromEmail,
		credentials: credentials,
	}
}

// SendDhanuEmail sends a plain text or HTML email based on the isHTML flag.
// Parameters:
// - to: The list of recipients.
// - subject: The subject of the email.
// - body: The content of the email.
// - isHTML: Flag to specify whether the email is in HTML format or plain text.
func (es *DhanuEmailService) SendDhanuEmail(to []string, subject, body string, isHTML bool) error {
	// Build the message
	msg, err := es.buildMessage(to, subject, body, isHTML, nil)
	if err != nil {
		return fmt.Errorf("failed to build email message: %v", err)
	}

	// Send the email
	return es.send(msg, to)
}

// SendDhanuEmailWithAttachments sends an email with or without HTML and includes attachments.
// Parameters:
// - to: The list of recipients.
// - subject: The subject of the email.
// - body: The content of the email.
// - isHTML: Flag to specify whether the email is in HTML format or plain text.
// - attachments: A list of file paths to attach to the email.
func (es *DhanuEmailService) SendDhanuEmailWithAttachments(to []string, subject, body string, isHTML bool, attachments []string) error {
	// Build the message with attachments
	msg, err := es.buildMessage(to, subject, body, isHTML, attachments)
	if err != nil {
		return fmt.Errorf("failed to build email message with attachments: %v", err)
	}

	// Send the email
	return es.send(msg, to)
}

// buildMessage constructs the email message.
// Parameters:
// - to: The list of recipients.
// - subject: The subject of the email.
// - body: The content of the email.
// - isHTML: Flag to specify whether the email is in HTML format or plain text.
// - attachments: A list of file paths to attach to the email (optional).
func (es *DhanuEmailService) buildMessage(to []string, subject, body string, isHTML bool, attachments []string) (string, error) {
	var buffer bytes.Buffer

	// Create a new MIME multipart writer.
	writer := multipart.NewWriter(&buffer)

	// Set email headers.
	contentType := "text/plain"
	if isHTML {
		contentType = "text/html"
	}

	// Write headers: From, To, Subject.
	buffer.WriteString("MIME-Version: 1.0\r\n")
	buffer.WriteString(fmt.Sprintf("From: %s\r\n", es.fromEmail))
	buffer.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ",")))
	buffer.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", writer.Boundary()))

	// Add the email body part.
	bodyPartHeader := make(textproto.MIMEHeader)
	bodyPartHeader.Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", contentType))
	bodyPartHeader.Set("Content-Transfer-Encoding", "quoted-printable")

	bodyPart, err := writer.CreatePart(bodyPartHeader)
	if err != nil {
		return "", fmt.Errorf("failed to create email body part: %v", err)
	}
	bodyPart.Write([]byte(body)) // Add the body content.

	// Handle attachments if any.
	for _, attachment := range attachments {
		err = es.addAttachment(writer, attachment)
		if err != nil {
			return "", fmt.Errorf("failed to add attachment: %v", err)
		}
	}

	// Close the multipart writer to finalize the message.
	writer.Close()

	return buffer.String(), nil
}

// addAttachment adds a file as an attachment to the email.
// Parameters:
// - writer: The MIME multipart writer.
// - filePath: Path of the file to attach.
func (es *DhanuEmailService) addAttachment(writer *multipart.Writer, filePath string) error {
	// Open the file to attach.
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open attachment: %v", err)
	}
	defer file.Close()

	// Get the file name and its MIME type.
	_, fileName := filepath.Split(filePath)
	mimeType := "application/octet-stream" // Generic binary stream for the zip file.

	// Create a header for the attachment part.
	attachmentHeader := make(textproto.MIMEHeader)
	attachmentHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	attachmentHeader.Set("Content-Type", mimeType)
	attachmentHeader.Set("Content-Transfer-Encoding", "base64")

	// Create the attachment part in the writer.
	part, err := writer.CreatePart(attachmentHeader)
	if err != nil {
		return fmt.Errorf("failed to create attachment part: %v", err)
	}

	// Copy the file content to the attachment part.
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to write attachment content: %v", err)
	}

	return nil
}

// send handles the actual sending of the email through SMTP.
// Parameters:
// - msg: The constructed email message.
// - to: The list of recipients.
func (es *DhanuEmailService) send(msg string, to []string) error {
	auth := smtp.PlainAuth("", es.fromEmail, es.credentials, es.smtpHost)

	// Use SMTPS (secure SMTP) to send the email with strong authentication
	addr := fmt.Sprintf("%s:%s", es.smtpHost, es.smtpPort)
	return smtp.SendMail(addr, auth, es.fromEmail, to, []byte(msg))
}
