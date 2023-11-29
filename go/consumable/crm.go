package consumable

import (
	"fmt"
	"net/smtp"
	"os"
)

func Crm(to string) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASS")

	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	subject := "MANUALL: Voce recebeu uma mensagem!"
	body := "Nosso assistente virtual Manuel te enviou uma mensagem"

	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
}
