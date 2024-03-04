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

Commands:
    scaffold - scaffold the project using provided blueprint
    cache    - save a remote directory to local cache to use later
    alias    - manage aliases

Use pomegranate <command> --help for more information about a specific command.
`

func printUsage() {
	fmt.Println(pomegranateUsage)
}

func Execute() error {
	flag.Usage = printUsage
	
	args := os.Args[1:]
	if len(args) == 0 {
		flag.Usage()
		return errors.New("No subcommand provided")
	}

    if cmd, ok := cmds[args[0]]; ok {
        if err := cmd.flagSet.Parse(args[1:]); err == nil {
            return cmd.exec()
        }
    }

    printUsage()
    return nil
}
