package cmd

import (
	"flag"
	"fmt"
)

var (
    aliasesCmd = flag.NewFlagSet("aliases", flag.ExitOnError)
)

func execAliases() error {
    aliasesMap, err := readAliasesFile()
    if err != nil {
        return err
    }

    // TODO: make this pretty
    for alias, url := range aliasesMap {
        fmt.Printf("%s => %s\n", alias, url)
    }

    return nil
}
