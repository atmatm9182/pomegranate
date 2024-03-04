package cmd

import (
	"errors"
	"flag"
	"path"

	"github.com/atmatm9182/pomegranate/gitapi"
	"github.com/atmatm9182/pomegranate/util"
)

var (
	cacheCmd   = flag.NewFlagSet("cache", flag.ExitOnError)
	cacheAlias = cacheCmd.String("alias", "", "cache the remote blueprint with a specified alias")
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
        if len(*cacheAlias) != 0 {
            return createAlias(repoUrl, *cacheAlias)
        }

        return nil
	}

    err := gitapi.Clone(repoUrl, destDir)
    if err != nil {
        return err
    }

    if len(*cacheAlias) == 0 {
        return nil
    }

    return createAlias(repoUrl, *cacheAlias)
}
