package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

type EmailData struct {
	Name string
	Code string
}

func main() {
	from := "XXX@gmail.com"
	password := "XXXX"
	to := []string{"YYYY@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Load and parse HTML template
	tmpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Dynamic data
	data := EmailData{
		Name: "Ranjan Kumar",
		Code: "123456",
	}

	// Execute template to buffer
	var body bytes.Buffer
	body.WriteString("Subject: Welcome to Go Mail Templates!\n")
	body.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")

	err = tmpl.Execute(&body, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	log.Println("Email sent successfully!")
}
