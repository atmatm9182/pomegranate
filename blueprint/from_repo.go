package blueprint

import (
	"os"
	"path"
	"strings"

	"github.com/atmatm9182/pomegranate/gitapi"
)

func FromRepo(url string, bluepintPath string) (Blueprint, error) {
	tmpDir := os.TempDir()
	if len(tmpDir) == 0 {
		panic("you don't have a temporary directory!")
	}

	urlParts := strings.Split(url, "/")
	repoFolderName := "pomegranate-"
	repoFolderName += strings.Join(urlParts[len(urlParts) - 2:], "-")
	folderName := path.Join(tmpDir, repoFolderName)

	err := gitapi.Clone(url, folderName)
	if err != nil {
		return Blueprint{}, err
	}

	return Parse(bluepintPath)
}
