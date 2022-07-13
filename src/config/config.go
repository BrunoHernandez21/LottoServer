package config

var DB = configDB{
	User:      "root",
	Password:  "",
	Soc:       "@tcp(127.0.0.1:3306)/",
	TableName: "lotto_music",
}

type configDB struct {
	User      string
	Password  string
	Soc       string
	TableName string
}

var Mail = configEmail{
	Email:      "lotto@inclusive.com.mx",
	ServerName: "mail.inclusive.com.mx:465",
	Host:       "mail.inclusive.com.mx",
	Password:   "fayGYp81lRt$",
}

type configEmail struct {
	Email      string
	ServerName string
	Host       string
	Password   string
}
