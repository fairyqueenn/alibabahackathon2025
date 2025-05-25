package constant

import "os"

var (
	OSS_ENDPOINT    string
	OSS_ACCESS_KEY  string
	OSS_SECRET_KEY  string
	OSS_BUCKET_NAME string
)

func InitOSSConstant() {
	OSS_ENDPOINT = os.Getenv("OSS_ENDPOINT")
	OSS_ACCESS_KEY = os.Getenv("OSS_ACCESS_KEY_ID")
	OSS_SECRET_KEY = os.Getenv("OSS_SECRET_KEY")
	OSS_BUCKET_NAME = os.Getenv("OSS_BUCKET_NAME")
}
