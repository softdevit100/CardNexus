package helpers

import "os"

func Env(key string) string {
	return os.Getenv(key)
}

func IsDebugMode() bool {
	return Env("DEBUG_MODE") == "true"
}
