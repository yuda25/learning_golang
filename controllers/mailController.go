package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func SendEmail(c *gin.Context) {
	var input struct {
		To string
	}
	c.Bind(&input)

	var body bytes.Buffer
	t, err := template.ParseFiles("./templates/email/email.html")
	t.Execute(&body, struct{Email string}{Email: input.To})

	if err != nil {
		fmt.Println(err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "yuda7246@gmail.com")
	m.SetHeader("To", input.To)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "yuda7246@gmail.com", os.Getenv("SMTP"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
