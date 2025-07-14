package config

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", username, password, host)
	msg := "From: " + username + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	addr := fmt.Sprintf("%s:%s", host, port)
	err := smtp.SendMail(addr, auth, username, []string{to}, []byte(msg))
	if err != nil {
		log.Println("SMTP Error:", err)
	}
	return err
}

func SendBulkEmail(toList []string, subject, body string) {
	for _, to := range toList {
		_ = SendEmail(strings.TrimSpace(to), subject, body)
	}
}
