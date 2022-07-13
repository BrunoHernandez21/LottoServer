package email

import (
	"crypto/tls"
	"lottomusic/src/config"
	"net"
	"net/mail"
	"net/smtp"
)

func Send_Recovery_Password(email string, password string) bool {

	from := mail.Address{Name: "Lotto", Address: config.Mail.Email}
	to := mail.Address{Name: "", Address: email}
	subj := "Lotto Music new password"
	body := "Lotto music " + "<" + config.Mail.Email + ">\n" + "Razón: " + "Nueva contraseña " + "\n" + "\nSu nueva contraseña se muestra a continuacion\nPassword: " + password

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += (k + ":" + v + "\r\n")
	}
	message += "\r\n" + body

	// Connect to the SMTP Server

	servername := config.Mail.ServerName

	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", config.Mail.Email, config.Mail.Password, config.Mail.Host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return false
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return false
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return false
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return false
	}

	if err = c.Rcpt(to.Address); err != nil {
		return false
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return false
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return false
	}

	err = w.Close()
	if err != nil {
		return false
	}

	c.Quit()
	return true
}
