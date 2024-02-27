package cmd

import (
	"flag"

	"github.com/atmatm9182/pomegranate/blueprint"
)

var (
	silentFlag bool
	nameFlag string
)

var cmds = map[string]*flag.FlagSet{
	"scaffold": scaffoldCmd,
	"cache":    cacheCmd,
}

func init() {
	for _, cmd := range cmds {
		cmd.BoolVar(&silentFlag, "silent", false, "disable all logging")
	}

	for _, cmd := range []*flag.FlagSet{scaffoldCmd, cacheCmd} {
		cmd.StringVar(&nameFlag, "name", blueprint.DefaultBlueprintPath, "the name of the blueprint file in the git repository")
	}
}
