package blueprint

import (
	"errors"
	"fmt"
	"path"

	"github.com/atmatm9182/pomegranate/util"
)

func FromCache(repoUrl string, configName string) (Blueprint, error) {
	cacheDir := util.GetCacheDirPath()
	repoDir := util.RepoUrlToFolderName(repoUrl)
	repoDir = path.Join(cacheDir, repoDir)
	fmt.Println(repoDir)

	if !util.FileExists(repoDir) {
		return Blueprint{}, errors.New("could not find cached blueprint")
	}

	blueprintPath := path.Join(repoDir, configName)
	return Parse(blueprintPath)
}
