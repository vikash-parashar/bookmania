package utils

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type EmailData struct {
	UserName string
	Subject  string
	HTMLBody string
}

// SendGreetingEmail sends a greeting email to the user.
func SendGreetingEmail(userName, userEmail, emailTemplatePath string) error {
	// Load the email template from file
	emailTemplate, err := template.ParseFiles(emailTemplatePath)
	if err != nil {
		return err
	}

	// Define the email data
	emailData := EmailData{
		UserName: userName,
		Subject:  "Welcome to Our Service",
	}

	// Execute the template to generate the HTML body
	var bodyBuffer bytes.Buffer
	if err := emailTemplate.Execute(&bodyBuffer, emailData); err != nil {
		return err
	}
	emailData.HTMLBody = bodyBuffer.String()

	// Configure the SMTP client
	smtpServer := "smtp.example.com"
	smtpPort := 587
	smtpUsername := "your-smtp-username"
	smtpPassword := "your-smtp-password"

	// Compose the email
	email := gomail.NewMessage()
	email.SetHeader("From", smtpUsername)
	email.SetHeader("To", userEmail)
	email.SetHeader("Subject", emailData.Subject)
	email.SetBody("text/html", emailData.HTMLBody)

	// Send the email
	dialer := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)
	if err := dialer.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
