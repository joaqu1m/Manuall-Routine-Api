package consumable

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func Crm(to string) {
	from := "xxxxxxxxxx"
	password := "xxxxxxxxxx"

	smtpHost := "smtp.office365.com"
	smtpPort := 587

	subject := "Test Email"
	body := "This is a test email sent using Golang."

	message := "Subject: " + subject + "\r\n" +
		"\r\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	err := sendWithStartTLS(smtpHost, smtpPort, auth, from, []string{to}, []byte(message), tlsConfig)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func sendWithStartTLS(smtpHost string, smtpPort int, auth smtp.Auth, from string, to []string, message []byte, tlsConfig *tls.Config) error {
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", smtpHost, smtpPort))
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.StartTLS(tlsConfig); err != nil {
		return err
	}

	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			return err
		}
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}
	_, err = wc.Write(message)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	return nil
}
