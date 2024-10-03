package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/lordofthemind/mygopher/gophersmtp"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an email with optional attachments",
	Long: `Send an email to a recipient with a subject and body. You can specify 
a recipient, subject, body text, or body file, and attach files to the email.`,
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
	sendCmd.Flags().StringSliceP("attachments", "a", []string{}, "List of file paths to attach to the email")
}

// Function to send an email
func sendEmail(cmd *cobra.Command) {
	// Load configuration to get default recipient
	config, _, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Get the recipient from the flag or use the default recipient
	to, _ := cmd.Flags().GetString("to")
	if to == "" {
		to = config.DefaultRecipient
	}
	if to == "" {
		fmt.Println("Error: No recipient specified and no default recipient found.")
		return
	}

	// Get the subject from the flag or use the current date and time as the subject
	subject, _ := cmd.Flags().GetString("subject")
	if subject == "" {
		subject = "Email sent on " + time.Now().Format("2006-01-02 15:04:05")
	}

	// Get the body from the flag or read the body from the file
	body, _ := cmd.Flags().GetString("body")
	bodyFile, _ := cmd.Flags().GetString("body-file")
	if body == "" && bodyFile != "" {
		bodyBytes, err := os.ReadFile(bodyFile)
		if err != nil {
			fmt.Printf("Error reading body file: %v\n", err)
			return
		}
		body = string(bodyBytes)
	}

	// Check if body is empty
	if body == "" {
		fmt.Println("Error: Email body cannot be empty. Please provide text or a file.")
		return
	}

	// Get attachments if any
	attachments, _ := cmd.Flags().GetStringSlice("attachments")

	// Initialize the Gopher SMTP email service with configuration values
	emailService := gophersmtp.NewEmailService(
		config.SMTP.Host,
		fmt.Sprintf("%d", config.SMTP.Port),
		config.SMTP.Username,
		config.SMTP.Password,
	)

	// Send the email (without or with attachments)
	if len(attachments) == 0 {
		err = emailService.SendEmail([]string{to}, subject, body, false)
	} else {
		err = emailService.SendEmailWithAttachments([]string{to}, subject, body, attachments, false)
	}

	// Handle sending errors
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}

	fmt.Println("Email sent successfully.")
}
