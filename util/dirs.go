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

func GetConfigDirPath() string {
    dir, err := os.UserConfigDir()
    if err != nil {
        panic(err)
    }

    return path.Join(dir, "pomegranate")
}
