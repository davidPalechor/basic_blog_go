package utils

import (
	"os"

	"github.com/astaxie/beego"
)

func GetAppConfig(key string) string {
	conf_key := beego.AppConfig.String(key)
	return os.Getenv(conf_key)
}
