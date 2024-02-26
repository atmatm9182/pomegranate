package cmd

import (
	"errors"
	"flag"
	"os"

	"github.com/atmatm9182/pomegranate/blueprint"
	"github.com/atmatm9182/pomegranate/blueprint/options"
	"github.com/atmatm9182/pomegranate/gitapi"
)

const (
	defaultScaffoldDest = "./"
)

var (
	scaffoldCmd    = flag.NewFlagSet("scaffold", flag.ExitOnError)
	scaffoldRemote = scaffoldCmd.Bool("remote", false, "scaffold the project using remote git repository")
	scaffoldName   = scaffoldCmd.String("name", blueprint.DefaultBlueprintPath, "the name of the blueprint file in the git repository")

	scaffoldDest = scaffoldCmd.String("o", defaultScaffoldDest, "where to scaffold the project")
)

func execScaffold() error {
	args := scaffoldCmd.Args()
	if len(args) == 0 {
		scaffoldCmd.Usage()
		return errors.New("Not enough arguments")
	}

	var (
		b   blueprint.Blueprint
		err error
	)

	if silentFlag {
		gitapi.EnableLogging = false
		options.DisableLogging()
	}

	options.SetScaffoldPrefix(*scaffoldDest)
	if *scaffoldDest != defaultScaffoldDest {
		err = os.MkdirAll(*scaffoldDest, 0777)
		if err != nil {
			return err
		}

		defer func() {
			if err != nil {
				os.RemoveAll(*scaffoldDest)
			}
		}()
	}

	if *scaffoldRemote {
		b, err = blueprint.FromRepo(args[0], *scaffoldName)
		if err != nil {
			return err
		}
	} else {
		b, err = blueprint.Parse(args[0])
		if err != nil {
			return err
		}
	}

	err = b.Scaffold()
	return err
}
