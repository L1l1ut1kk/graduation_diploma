package autoriz

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

func GeneratePassword(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(chars))
		password[i] = chars[randomIndex]
	}

	return string(password)
}

func SendPasswordResetEmail(mail string) error {
	from := mail
	password := GeneratePassword(8)

	subject := "Password Reset"
	body := fmt.Sprintf("Your new password: %s", password)
	message := "From: " + from + "\n" +
		"To: " + mail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, "your-password", "smtp.example.com") // Замените "your-password" на ваш реальный пароль

	err := smtp.SendMail("smtp.example.com:587", auth, from, []string{mail}, []byte(message)) // Замените "smtp.example.com:587" на реальный адрес SMTP-сервера и порт
	if err != nil {
		return err
	}

	return nil
}
