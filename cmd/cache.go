package cmd

import (
	"errors"
	"flag"
	"os"
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
	cacheDir := getCacheDir()
	destDir := path.Join(cacheDir, "pomegranate", util.RepoNameToFolderName(repoUrl))

	if util.FileExists(destDir) {
		return nil
	}

	return gitapi.Clone(repoUrl, destDir)
}

func getCacheDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic("Your system does not support caching!")
	}

	return dir
}
