package main

import (
	"bytes"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"

	"net/smtp"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"", "mostafa.omrani71@gmail.com", "gwowuxhoivzxzhbo", "smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail("smtp.gmail.com:587", auth, "mostafa.omrani71@gmail.com", to, []byte(msg))

	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHtml(subject string, templatePath string, to []string) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "mostafa"})

	auth := smtp.PlainAuth(
		"", "mostafa.omrani71@gmail.com", "gwowuxhoivzxzhbo", "smtp.gmail.com",
	)

	header := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + header + "\n\n" + body.String()

	err = smtp.SendMail("smtp.gmail.com:587", auth, "mostafa.omrani71@gmail.com", to, []byte(msg))

	if err != nil {
		fmt.Println(err)
	}
}

func sendGoMail(templatePath string) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "mostafa"})

	if err != nil {
		fmt.Println(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "mostafa.omrani71@gmail.com")
	m.SetHeader("To", "mostafa.omrani71@gmail.com")
	// m.SetAddressHeader("Cc", "mostafa.omrani71@gmail.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./404.png")

	d := gomail.NewDialer("smtp.gmail.com", 587, "mostafa.omrani71@gmail.com", "gwowuxhoivzxzhbo")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func main() {
	// sendMailSimple("salam", "my body", []string{"mostafa.omrani71@gmail.com"})
	// sendMailSimpleHtml("salam", "./test.html", []string{"mostafa.omrani71@gmail.com"})
	sendGoMail("./test.html")
}
