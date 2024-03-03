package blueprint

import (
	"os"
	"path"

	"github.com/atmatm9182/pomegranate/gitapi"
	"github.com/atmatm9182/pomegranate/util"
)

func FromRepo(url string, blueprintPath string) (Blueprint, error) {
	tmpDir := os.TempDir()
	if len(tmpDir) == 0 {
		panic("you don't have a temporary directory!")
	}

	repoFolderName := util.RepoUrlToFolderName(url)
	repoFolderName += "pomegranate"
	folderName := path.Join(tmpDir, repoFolderName)

	err := gitapi.Clone(url, folderName)
	if err != nil {
		return Blueprint{}, err
	}

	blueprintPath = path.Join(folderName, blueprintPath)
	return Parse(blueprintPath)
}
