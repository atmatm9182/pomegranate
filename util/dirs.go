package util

import (
	"os"
	"path"
)

func GetCacheDirPath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic("Your system does not support caching!")
	}

	return path.Join(dir, "pomegranate")
}
