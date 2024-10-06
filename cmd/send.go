package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lordofthemind/dhanu/internals/services"
	"github.com/lordofthemind/dhanu/internals/utils"
	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an email with optional attachments",
	Long: `Send an email to a recipient with a subject, body. 
You can also specify attachments or folders to zip and attach.`,
	Run: func(cmd *cobra.Command, args []string) {
		sendEmail(cmd)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Define flags for sending email
	sendCmd.Flags().StringP("to", "t", "", "Recipient email address")
	sendCmd.Flags().StringP("subject", "s", "", "Email subject")
	sendCmd.Flags().StringP("body", "b", "", "Email body text")
	sendCmd.Flags().StringP("body-file", "f", "", "Path to text file for email body")
	sendCmd.Flags().StringSliceP("attachments", "a", []string{}, "List of file paths or folders to attach to the email")
}

func sendEmail(cmd *cobra.Command) {
	// Check if any flags were provided
	if cmd.Flags().NFlag() == 0 {
		_ = cmd.Help()
		return
	}

	// Load configuration to get default recipient
	config, _, err := configs.LoadConfig()
	if err != nil {
		log.Println("Error loading configuration:", err)
		return
	}

	// Get the recipient from the flag or use the default recipient
	to, _ := cmd.Flags().GetString("to")
	if to == "" {
		to = config.DefaultRecipient
	}
	if to == "" {
		log.Println("Error: No recipient specified and no default recipient found.")
		return
	}

	// Validate the recipient email
	if !utils.IsValidEmail(to) {
		log.Println("Error: Invalid recipient email address.")
		return
	}

	// Get the subject from the flag or use the current Unix timestamp as the subject
	subject, _ := cmd.Flags().GetString("subject")
	if subject == "" {
		subject = fmt.Sprintf("Email sent at %s", time.Now().Format(time.RFC1123))
	} else if len(subject) > 78 {
		log.Println("Error: Subject exceeds 78 character limit.")
		return
	}

	// Get the body from the flag or read the body from the file
	body, _ := cmd.Flags().GetString("body")
	bodyFile, _ := cmd.Flags().GetString("body-file")
	if body == "" && bodyFile != "" {
		bodyBytes, err := os.ReadFile(bodyFile)
		if err != nil {
			log.Printf("Error reading body file: %v\n", err)
			return
		}
		body = string(bodyBytes)
	}

	// Check if body is empty
	if body == "" {
		log.Println("Error: Email body cannot be empty. Please provide text or a file.")
		return
	}

	// Get attachments if any
	attachments, _ := cmd.Flags().GetStringSlice("attachments")

	// Initialize the Dhanu email service with configuration values
	emailService := services.NewDhanuEmailService(
		config.SMTP.Host,
		fmt.Sprintf("%d", config.SMTP.Port),
		config.SMTP.FromEmail,
		config.SMTP.Credentials,
	)

	// Determine how to send the email based on the presence of attachments
	if len(attachments) > 0 {
		// Send email with attachments
		err = emailService.SendDhanuEmailWithAttachments(
			[]string{to}, // To recipients
			subject,      // Subject
			body,         // Body
			false,        // isHtml flag (set to true for HTML content)
			attachments,  // Attachments
		)
	} else {
		// Send a simple email without attachments
		err = emailService.SendDhanuEmail(
			[]string{to}, // To recipients
			subject,      // Subject
			body,         // Body
			false,        // isHtml flag (set to true for HTML content)
		)
	}

	// Handle sending errors
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return
	}

	log.Println("Email sent successfully.")
}
