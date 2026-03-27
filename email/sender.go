package email

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

var (
	smtpHost     string
	smtpPort     string
	smtpUser     string
	smtpPassword string
	fromEmail    string
)

func init() {
	smtpHost = os.Getenv("SMTP_HOST")
	smtpPort = os.Getenv("SMTP_PORT")
	smtpUser = os.Getenv("SMTP_USER")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
	fromEmail = os.Getenv("SMTP_FROM")

	if smtpHost == "" {
		smtpHost = "smtp.gmail.com"
	}
	if smtpPort == "" {
		smtpPort = "587"
	}
	if fromEmail == "" {
		fromEmail = smtpUser
	}
}

func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func SendVerificationLink(toEmail, link, purpose string) error {
	if smtpUser == "" || smtpPassword == "" {
		log.Printf("EMAIL NOT CONFIGURED - Verification link for %s (%s): %s", toEmail, purpose, link)
		return nil
	}

	var subject, body string

	if purpose == "register" {
		subject = "JudoQuiz - Verify Your Email"
		body = fmt.Sprintf(
			"Welcome to JudoQuiz!\n\nPlease verify your email by clicking the link below:\n\n%s\n\nThis link expires in 24 hours.\n\nIf you didn't create an account, please ignore this email.",
			link,
		)
	} else {
		subject = "JudoQuiz - Confirm Password Change"
		body = fmt.Sprintf(
			"You requested a password change.\n\nClick the link below to confirm:\n\n%s\n\nThis link expires in 10 minutes.\n\nIf you didn't request this, please ignore this email.",
			link,
		)
	}

	msg := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=utf-8\r\n\r\n%s",
		fromEmail, toEmail, subject, body,
	)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	addr := smtpHost + ":" + smtpPort

	if err := smtp.SendMail(addr, auth, fromEmail, []string{toEmail}, []byte(msg)); err != nil {
		log.Printf("Failed to send email to %s: %v", toEmail, err)
		return fmt.Errorf("failed to send email")
	}

	log.Printf("Verification email sent to %s", toEmail)
	return nil
}
