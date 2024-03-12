package handle

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/argparse"
	"github.com/pspiagicw/qemantra/help"
	"github.com/pspiagicw/qemantra/manage"
)

func HandleArgs(opts *argparse.Opts) {
	checkArgs(opts)

	handlers := map[string]func([]string){
		"create": manage.CreateVM,
		"run":    manage.RunVM,
		"edit":   manage.EditVM,
		"rename": manage.RenameVM,
		"list":   manage.ListVM,
		"delete": manage.RemoveVM,
		"version": func([]string) {
			help.PrintVersion(opts.Version)
		},
		"help": func(args []string) {
			help.HandleHelp(args, opts.Version)
		},
	}

	cmd := opts.Args[0]

	handleFunc, ok := handlers[cmd]

	if !ok {
		help.HelpUsage(opts.Version)
		goreland.LogFatal("No command named '%s'", cmd)
	} else {
		handleFunc(opts.Args[1:])
	}
}
func checkArgs(opts *argparse.Opts) {
	if len(opts.Args) == 0 {
		help.HelpUsage(opts.Version)
		goreland.LogFatal("No subcommand provided!")
	}
}
