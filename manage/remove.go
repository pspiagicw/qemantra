package manage

import (
	"flag"

	"github.com/pspiagicw/qemantra/help"
)

func parseRemoveArgs(args []string) {
	flag := flag.NewFlagSet("qemantra remove", flag.ExitOnError)

	flag.Usage = help.HelpRemove

	flag.Parse(args)
}
func RemoveVM(args []string) {
	parseRemoveArgs(args)

	_, selected := selectMachine()

	deleteFile(getMachinePath(selected))
}
