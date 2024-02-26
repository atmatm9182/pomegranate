package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/atmatm9182/pomegranate/blueprint"
)

var (
	scaffoldCmd = flag.NewFlagSet("scaffold", flag.ExitOnError)
	scaffoldRemote = scaffoldCmd.Bool("remote", false, "scaffold the project using remote git repository")
	scaffoldName = scaffoldCmd.String("name", blueprint.DefaultBlueprintPath, "the name of the file containing the blueprint")
)

var cmds = map[string]*flag.FlagSet {
	"scaffold": scaffoldCmd,
}

const pomegranateUsage = `pomegranate is a tool for project scaffolding.

Usage:
    pomegranate <command> [arguments]

The commands are:
    scaffold - scaffold the project using provided blueprint

Use pomegranate <command> --help for more information about a specific command.
`

func usage() {
	fmt.Println(pomegranateUsage)
}

func execScaffold() error {
	args := scaffoldCmd.Args()
	if len(args) == 0 {
		scaffoldCmd.Usage()
		return errors.New("Not enough arguments")
	}

	var (
		b blueprint.Blueprint
		err error
	)
	
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

	return b.Scaffold()
}

func Execute() error {
	flag.Usage = usage
	
	args := os.Args[1:]
	if len(args) == 0 {
		flag.Usage()
		return errors.New("No subcommand provided")
	}

	switch args[0] {
	case "scaffold":
		if err := scaffoldCmd.Parse(args[1:]); err == nil {
			return execScaffold()
		}
	case "--help", "-help", "-h":
		flag.Usage()
		os.Exit(0)
	default:
		return fmt.Errorf("Unknown command %s", args[0])
	}

	panic("should be unreachable")
}
