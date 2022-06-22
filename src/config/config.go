package config

var DB = configDB{
	User:      "root",
	Soc:       "",
	Password:  "@tcp(127.0.0.1:3306)/",
	TableName: "lotto_music",
}

type configDB struct {
	User      string
	Password  string
	Soc       string
	TableName string
}

var Mail = configEmail{
	Email:      "ichimar21@gmail.com",
	ServerName: "smtp.gmail.com:587",
	Host:       "smtp.gmail.com",
	Password:   "xpxqhixlrsldedzq",
}

type configEmail struct {
	Email      string
	ServerName string
	Host       string
	Password   string
}
