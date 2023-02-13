package config

import "os"

var (
	Locale              string
	GgStorageCredential string
	BucketName          string
	PublicUrlGgStorage  string
)

func getConfig() {
	Locale = os.Getenv("PJ_LOCALE")
	GgStorageCredential = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	BucketName = os.Getenv("BUCKET_NAME")
	PublicUrlGgStorage = os.Getenv("GOOGLE_STORAGE_PUBLIC")

}

func ContentPath() string {
	return "public/content/"
}
func ThumbnailPath() string {
	return "public/thumbnail/"
}
