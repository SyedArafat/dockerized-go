package utils

import "os"

func Config(key string) string {
	return os.Getenv(key)
}
