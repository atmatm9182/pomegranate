package cmd

import (
	"errors"
	"flag"
	"os"
	"path"

	"github.com/atmatm9182/pomegranate/blueprint"
	"github.com/atmatm9182/pomegranate/blueprint/options"
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
	opts := options.DefaultScaffolding()
	destDir := path.Join(cacheDir, "pomegranate", util.RepoNameToFolderName(repoUrl))

	if util.FileExists(destDir) {
		return nil
	}
	
	b, err := blueprint.FromRepo(repoUrl, nameFlag)
	if err != nil {
		return err
	}

	sc := blueprint.NewScaffolder(&opts)
	return sc.Scaffold(&b)
}

func getCacheDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic("Your system does not support caching!")
	}

	return dir
}
