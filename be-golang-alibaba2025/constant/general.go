package constant

import (
	"os"
	"strconv"
)

var APP_PROD bool
var WEB_URL string

func InitGeneralConstant() {
	APP_PROD, _ = strconv.ParseBool(os.Getenv("APP_PROD"))

	WEB_URL = os.Getenv("DEV_WEB_URL")

	if APP_PROD {
		WEB_URL = os.Getenv("WEB_URL")
	}
}
