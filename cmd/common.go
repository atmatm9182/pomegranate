package cmd

import "flag"

var (
	silentFlag bool
)

var cmds = map[string]*flag.FlagSet {
	"scaffold": scaffoldCmd,
}

func init() {
	for _, cmd := range cmds {
		cmd.BoolVar(&silentFlag, "silent", false, "disable all logging")
	}
}
