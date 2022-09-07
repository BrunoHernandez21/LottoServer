package config

var DB = ConfigDB{}

type ConfigDB struct {
	User         string
	Password     string
	Soc          string
	DatabaseName string
}

var Mail = ConfigEmail{}

type ConfigEmail struct {
	Email      string
	ServerName string
	Host       string
	Password   string
}

var Rest_Port = ""
var JwtKey = []byte("TestForFasthttpWithJWT")

const Rest_version string = "/v3/"
