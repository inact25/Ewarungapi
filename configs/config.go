package configs

import (
	"github.com/inact25/E-WarungApi/utils/environtment"
)

type dbConf struct {
	DbUser               string
	DbPass               string
	DbHost               string
	DbPort               string
	DbSchema             string
	AllowNativePasswords bool
}

type Conf struct {
	Db dbConf
}

func NewAppConfig() *Conf {
	return &Conf{dbConf{
		DbUser:               environtment.Get("DbUser", "root"),
		DbPass:               environtment.Get("DbPass", "yourPass"),
		DbHost:               environtment.Get("DbHost", "localhost"),
		DbPort:               environtment.Get("DbPort", "8080"),
		DbSchema:             environtment.Get("DbSchema", "yourSchema"),
		AllowNativePasswords: true,
	}}
}
