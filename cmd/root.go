package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

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

func Execute() error {
	flag.Usage = usage
	
	args := os.Args[1:]
	if len(args) == 0 {
		flag.Usage()
		return errors.New("No subcommand provided")
	}

	// TODO: refactor this
	switch args[0] {
	case "scaffold":
		if err := scaffoldCmd.Parse(args[1:]); err == nil {
			return execScaffold()
		}
	case "cache":
		if err := cacheCmd.Parse(args[1:]); err == nil {
			return execCache()
		}
	case "--help", "-help", "-h":
		flag.Usage()
		os.Exit(0)
	default:
		return fmt.Errorf("Unknown command %s", args[0])
	}

	panic("should be unreachable")
}
