package manage

import (
	"flag"

	"github.com/pspiagicw/qemantra/help"
	"github.com/pspiagicw/qemantra/vm"
)

func parseEditArgs(args []string) {
	flag := flag.NewFlagSet("qemantra edit", flag.ExitOnError)

	flag.Usage = help.HelpEdit

	flag.Parse(args)
}
func EditVM(args []string) {

	parseEditArgs(args)

	name, selected := selectMachine()

	newMachine := vm.PromptMachine(selected)

	if newMachine.DiskName != selected.DiskName {
		checkImage(newMachine)
	}

	newMachine.Name = name

	// Save new machine to old machine path
	saveToDisk(newMachine, getMachinePath(selected))
}
