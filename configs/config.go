package configs

import (
	"github.com/inact25/E-WarungApi/utils"
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
		DbUser:               utils.GetCustomConf("DbUser", "root"),
		DbPass:               utils.GetCustomConf("DbPass", "yourPass"),
		DbHost:               utils.GetCustomConf("DbHost", "localhost"),
		DbPort:               utils.GetCustomConf("DbPort", "8080"),
		DbSchema:             utils.GetCustomConf("DbSchema", "yourSchema"),
		AllowNativePasswords: true,
	}}
}
