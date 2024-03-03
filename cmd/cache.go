package cmd

import (
	"errors"
	"flag"
	"path"

	"github.com/atmatm9182/pomegranate/gitapi"
	"github.com/atmatm9182/pomegranate/util"
)

var (
	cacheCmd = flag.NewFlagSet("cache", flag.ExitOnError)
)

func execCache() error {
	args := cacheCmd.Args()
	if len(args) == 0 {
		cacheCmd.Usage()
		return errors.New("Not enough arguments")
	}
	
	repoUrl := args[0]
	cacheDir := util.GetCacheDirPath()
	destDir := path.Join(cacheDir, util.RepoUrlToFolderName(repoUrl))

	if util.FileExists(destDir) {
		return nil
	}

	return gitapi.Clone(repoUrl, destDir)
}
