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

func emailTemplate(title, content string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"></head>
<body style="margin:0;padding:0;background-color:#0f0f23;font-family:'Helvetica Neue',Arial,sans-serif;">
<table width="100%%" cellpadding="0" cellspacing="0" style="background-color:#0f0f23;padding:40px 20px;">
<tr><td align="center">
<table width="100%%" cellpadding="0" cellspacing="0" style="max-width:520px;background-color:#1a1a2e;border-radius:16px;border:1px solid rgba(255,255,255,0.1);overflow:hidden;">
<tr><td style="background:linear-gradient(135deg,#667eea 0%%,#764ba2 100%%);padding:30px;text-align:center;">
<h1 style="margin:0;color:#ffffff;font-size:22px;font-weight:700;letter-spacing:0.5px;">%s</h1>
</td></tr>
<tr><td style="padding:32px 32px 40px;">
%s
<p style="color:#555;font-size:12px;margin:32px 0 0;text-align:center;border-top:1px solid rgba(255,255,255,0.06);padding-top:20px;">
JudoQuiz &mdash; Judo Club Management
</p>
</td></tr>
</table>
</td></tr>
</table>
</body>
</html>`, title, content)
}

func button(label, href, color string) string {
	if color == "" {
		color = "#667eea"
	}
	return fmt.Sprintf(`<table width="100%%" cellpadding="0" cellspacing="0" style="margin:8px 0;"><tr><td align="center">
<a href="%s" style="display:inline-block;background:%s;color:#ffffff;text-decoration:none;padding:14px 32px;border-radius:10px;font-size:15px;font-weight:600;letter-spacing:0.3px;">%s</a>
</td></tr></table>`, href, color, label)
}

func text(s string) string {
	return fmt.Sprintf(`<p style="color:#cccccc;font-size:15px;line-height:1.6;margin:0 0 16px;">%s</p>`, s)
}

func smallText(s string) string {
	return fmt.Sprintf(`<p style="color:#888888;font-size:13px;line-height:1.5;margin:0 0 12px;">%s</p>`, s)
}

func SendVerificationLink(toEmail, link, purpose string) error {
	if smtpUser == "" || smtpPassword == "" {
		log.Printf("EMAIL NOT CONFIGURED - Verification link for %s (%s): %s", toEmail, purpose, link)
		return nil
	}

	var subject, html string

	switch purpose {
	case "register":
		subject = "Verify your email — JudoQuiz"
		html = emailTemplate("Welcome to JudoQuiz", ""+
			text("Thanks for signing up! Please verify your email address to get started.")+
			button("Verify Email", link, "")+
			smallText("This link expires in 24 hours.")+
			smallText("If you didn't create an account, you can safely ignore this email."),
		)
	default:
		subject = "Confirm password change — JudoQuiz"
		html = emailTemplate("Password Change", ""+
			text("You requested a password change for your JudoQuiz account.")+
			button("Confirm Change", link, "")+
			smallText("This link expires in 10 minutes.")+
			smallText("If you didn't request this, you can safely ignore this email."),
		)
	}

	return sendHTML(toEmail, subject, html)
}

func SendClubInvite(toEmail, clubName, acceptLink, denyLink string) error {
	subject := fmt.Sprintf("You're invited to join %s — JudoQuiz", clubName)
	html := emailTemplate(fmt.Sprintf("Join %s", clubName), ""+
		text(fmt.Sprintf("You've been invited to join <strong style=\"color:#fff;\">%s</strong> on JudoQuiz.", clubName))+
		text("Click below to accept or decline the invitation.")+
		button("Accept Invitation", acceptLink, "")+
		button("Decline", denyLink, "rgba(255,255,255,0.1)")+
		smallText("This invitation expires in 7 days."),
	)
	return sendHTML(toEmail, subject, html)
}

func SendNotification(toEmail, subject, body string) error {
	html := emailTemplate("JudoQuiz Notification", text(body))
	return sendHTML(toEmail, subject, html)
}

func SendEmail(toEmail, subject, body string) error {
	if smtpUser == "" || smtpPassword == "" {
		log.Printf("EMAIL NOT CONFIGURED - To: %s, Subject: %s", toEmail, subject)
		return nil
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

	log.Printf("Email sent to %s: %s", toEmail, subject)
	return nil
}

func sendHTML(toEmail, subject, html string) error {
	if smtpUser == "" || smtpPassword == "" {
		log.Printf("EMAIL NOT CONFIGURED - To: %s, Subject: %s", toEmail, subject)
		return nil
	}

	msg := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=utf-8\r\n\r\n%s",
		fromEmail, toEmail, subject, html,
	)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	addr := smtpHost + ":" + smtpPort

	if err := smtp.SendMail(addr, auth, fromEmail, []string{toEmail}, []byte(msg)); err != nil {
		log.Printf("Failed to send email to %s: %v", toEmail, err)
		return fmt.Errorf("failed to send email")
	}

	log.Printf("Email sent to %s: %s", toEmail, subject)
	return nil
}
