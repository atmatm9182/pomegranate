package cmd

import (
	"encoding/json"
	"flag"
	"os"
	"path"

	"github.com/atmatm9182/pomegranate/blueprint"
	"github.com/atmatm9182/pomegranate/util"
)

var (
	silentFlag bool
	nameFlag string
)

var cmds = map[string]*flag.FlagSet{
	"scaffold": scaffoldCmd,
	"cache":    cacheCmd,
    "aliases":  aliasesCmd,
}

func init() {
	for _, cmd := range cmds {
		cmd.BoolVar(&silentFlag, "silent", false, "disable all logging")
	}

	for _, cmd := range []*flag.FlagSet{scaffoldCmd, cacheCmd} {
		cmd.StringVar(&nameFlag, "name", blueprint.DefaultBlueprintPath, "the name of the blueprint file in the git repository")
	}
}

const aliasesFileName = "aliases.json"

func getAliasesFile() (*os.File, error) {
    configDir := util.GetConfigDirPath()
    err := os.MkdirAll(configDir, 0777)
    if err != nil {
        return nil, err
    }

	aliasesPath := path.Join(configDir, aliasesFileName)
	return os.OpenFile(aliasesPath, os.O_CREATE | os.O_RDWR, 0666)
}

func readAliasesFile() (map[string]string, error) {
    aliasesFile, err := getAliasesFile()
    if err != nil {
        return nil, err
    }
	defer aliasesFile.Close()

	decoder := json.NewDecoder(aliasesFile)
    if !decoder.More() {
        return make(map[string]string, 5), nil
    }

	var aliases map[string]string
	err = decoder.Decode(&aliases)
    return aliases, err
}

func writeAliasesToFile(aliases map[string]string) error {
    aliasesFile, err := getAliasesFile()
    if err != nil {
        return err
    }
    defer aliasesFile.Close()

    encoder := json.NewEncoder(aliasesFile)
    return encoder.Encode(aliases)
}
