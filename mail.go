package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func init(){
	if err := godotenv.Load(); err != nil {
		return 
	}
}

func sendMailSimple(subject, body string, to []string) {
	code := os.Getenv("code")

	auth := smtp.PlainAuth(
		"",
		"avinfajar6@gmail.com",
		code,
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"avinfajar6@gmail.com",
		to,
		[]byte(msg),
	); err != nil {
		log.Println(err)
	}

}

func main() {
	sendMailSimple("ini adalah subject", "ini adalah body", []string{"avinfajar6@gmail.com"})
}
