package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/lordofthemind/dhanu/internals/utils"
	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/lordofthemind/mygopher/gophersmtp"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an email with optional attachments, CC, and BCC",
	Long: `Send an email to a recipient with a subject, body, CC, and BCC. 
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
	sendCmd.Flags().StringSliceP("cc", "C", []string{}, "List of CC recipients")   // Uppercase C for CC
	sendCmd.Flags().StringSliceP("bcc", "B", []string{}, "List of BCC recipients") // Uppercase B for BCC
}

func sendEmail(cmd *cobra.Command) {
	// Check if any flags were provided
	if cmd.Flags().NFlag() == 0 {
		// Show help and return if no flags are passed
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

	// Get CC and BCC recipients from flags
	cc, _ := cmd.Flags().GetStringSlice("cc")
	bcc, _ := cmd.Flags().GetStringSlice("bcc")

	// Validate CC and BCC emails
	for _, email := range append(cc, bcc...) {
		if !utils.IsValidEmail(email) {
			log.Printf("Error: Invalid email address in CC or BCC: %s\n", email)
			return
		}
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

	// Get attachments if any, and handle folders (zip them)
	attachments, _ := cmd.Flags().GetStringSlice("attachments")
	var finalAttachments []string
	for _, attachment := range attachments {
		fileInfo, err := os.Stat(attachment)
		if err != nil {
			log.Printf("Error reading attachment: %v\n", err)
			return
		}

		if fileInfo.IsDir() {
			// Zip the folder
			zipFilePath := fmt.Sprintf("%s.zip", attachment)
			err := zipFolder(attachment, zipFilePath)
			if err != nil {
				log.Printf("Error zipping folder: %v\n", err)
				return
			}
			finalAttachments = append(finalAttachments, zipFilePath)
		} else {
			finalAttachments = append(finalAttachments, attachment)
		}
	}

	// Ensure attachments don't exceed 25 MB in total size
	if totalSize, err := checkTotalAttachmentSize(finalAttachments); err != nil {
		log.Println("Error checking attachment sizes:", err)
		return
	} else if totalSize > 25*1024*1024 {
		log.Println("Error: Attachments exceed 25 MB limit.")
		return
	}

	// Initialize the Gopher SMTP email service with configuration values
	emailService := gophersmtp.NewEmailService(
		config.SMTP.Host,
		fmt.Sprintf("%d", config.SMTP.Port),
		config.SMTP.FromEmail,
		config.SMTP.Credentials,
	)

	// Determine how to send the email based on the presence of CC, BCC, and attachments
	if len(finalAttachments) > 0 {
		// If there are attachments, send using the appropriate method
		if len(cc) > 0 || len(bcc) > 0 {
			// If there are CC or BCC recipients
			err = emailService.SendEmailWithCCAndBCCAndAttachments(
				[]string{to},     // To recipients
				cc,               // CC recipients
				bcc,              // BCC recipients
				subject,          // Subject
				body,             // Body
				finalAttachments, // Attachments
				false,            // isHtml flag (set to false for plain text)
			)
		} else {
			// No CC or BCC, send with attachments only
			err = emailService.SendEmail(
				[]string{to}, // To recipients
				subject,      // Subject
				body,         // Body
				false,        // isHtml flag (set to false for plain text)
			)
		}
	} else {
		// No attachments, send with CC and BCC if provided
		if len(cc) > 0 || len(bcc) > 0 {
			err = emailService.SendEmailWithCCAndBCC(
				[]string{to}, // To recipients
				cc,           // CC recipients
				bcc,          // BCC recipients
				subject,      // Subject
				body,         // Body
				false,        // isHtml flag (set to false for plain text)
			)
		} else {
			// Send a simple email without CC, BCC, or attachments
			err = emailService.SendEmail(
				[]string{to}, // To recipients
				subject,      // Subject
				body,         // Body
				false,        // isHtml flag (set to false for plain text)
			)
		}
	}

	// Handle sending errors
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return
	}

	log.Println("Email sent successfully.")
}

// Zip a folder to a specified destination zip file
func zipFolder(folderPath string, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Preserve the folder structure in the zip file
		header.Name, _ = filepath.Rel(folderPath, path)
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// Check the total size of the attachments to ensure it does not exceed the limit
func checkTotalAttachmentSize(attachments []string) (int64, error) {
	var totalSize int64
	for _, attachment := range attachments {
		fileInfo, err := os.Stat(attachment)
		if err != nil {
			return 0, err
		}
		totalSize += fileInfo.Size()
	}
	return totalSize, nil
}
