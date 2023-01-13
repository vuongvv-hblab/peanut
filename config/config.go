package config

import "os"

var Locale string

func getConfig() {
	Locale = os.Getenv("PJ_LOCALE")
}

func ContentPath() string {
	return "public/content/"
}
func ThumbnailPath() string {
	return "public/thumbnail/"
}
