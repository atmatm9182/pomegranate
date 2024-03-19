package util

import (
	"os"
	"strings"
)

func RepoUrlToFolderName(url string) string {
	urlParts := strings.Split(url, "/")

	idx := max(len(urlParts)-2, 0)
	return strings.Join(urlParts[idx:], "-")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
