package blueprint

import (
	"os"
	"path"
	"strings"
	
	"github.com/atmatm9182/pomegranate/gitapi"
)

func FromRepo(url string, blueprintPath string) (Blueprint, error) {
	tmpDir := os.TempDir()
	if len(tmpDir) == 0 {
		panic("you don't have a temporary directory!")
	}
	
	urlParts := strings.Split(url, "/")
	
	idx := max(len(urlParts) - 2, 0)
	repoFolderName := "pomegranate-"
	repoFolderName += strings.Join(urlParts[idx:], "-")
	folderName := path.Join(tmpDir, repoFolderName)

	err := gitapi.Clone(url, folderName)
	if err != nil {
		return Blueprint{}, err
	}

	return Parse(blueprintPath)
}
