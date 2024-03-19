package cmd

import (
	"errors"
	"flag"
	"fmt"
)

var (
	aliasesCmd = flag.NewFlagSet("alias", flag.ExitOnError)
)

func execAliases() error {
	args := aliasesCmd.Args()
	if len(args) == 0 {
		printAliases()
		return nil
	}

	switch args[0] {
	case "list":
		printAliases()
		return nil
	case "add":
		if len(args) < 3 {
			return errors.New("Not enough arguments")
		}

		return createAlias(args[2], args[1])
	}

	return nil
}

func printAliases() {
	aliasesMap, err := readAliasesFile()
	if err != nil {
		return
	}

	// TODO: make this pretty
	for alias, url := range aliasesMap {
		fmt.Printf("%s => %s\n", alias, url)
	}
}
