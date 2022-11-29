package config

import (
	"fmt"
	"time"
)

const (
	User           = "root"
	Password       = "123456"
	Host           = "172.27.106.74:3306"
	DataBaseName   = "testmes"
	TokenValidTime = time.Hour * 1
	HttpPort       = "5000"
)

var DSN string

func init() {
	DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User, Password, Host, DataBaseName)
}
