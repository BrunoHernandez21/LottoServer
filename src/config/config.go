package config

var DB configDB = configDB{}

type configDB struct {
	User      string
	Password  string
	Soc       string
	TableName string
}

var Mail = configEmail{}

type configEmail struct {
	Email      string
	ServerName string
	Host       string
	Password   string
}
