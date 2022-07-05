package email

import (
	"lottomusic/src/config"
	"net/smtp"
)

func Send_Password(email string, password string) bool {

	from := smtp.PlainAuth("", config.Mail.Email, config.Mail.Password, config.Mail.Host)
	to := []string{email}

	titulo := "From: Lotto music "
	Subtitulo := "Nueva contraseña "
	msg := []byte(titulo + "<" + config.Mail.Email + ">\n" + "Subject: " + Subtitulo + "\n" + "\nSu nueva contraseña se muestra a continuacion\nPassword: " + password)
	erro := smtp.SendMail(config.Mail.ServerName, from, config.Mail.Email, to, msg)
	return erro != nil
}

func Send_Recovery_Password(email string, password string) bool {

	from := smtp.PlainAuth("", config.Mail.Email, config.Mail.Password, config.Mail.Host)
	to := []string{email}

	titulo := "From: Lotto music "
	Subtitulo := "Nueva contraseña "
	msg := []byte(titulo + "<" + config.Mail.Email + ">\n" + "Subject: " + Subtitulo + "\n" + "\nSu nueva contraseña se muestra a continuacion\nPassword: " + password)
	erro := smtp.SendMail(config.Mail.ServerName, from, config.Mail.Email, to, msg)
	return erro != nil
}
