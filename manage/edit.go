package manage

import (
	"flag"

	"github.com/pspiagicw/qemantra/vm"
)

func EditVM(args []string) {
	flag := flag.NewFlagSet("qemantra edit", flag.ExitOnError)

	flag.Parse(args)

	name, selected := selectMachine()

	newMachine := vm.GetMachine(selected)

	if newMachine.DiskName != selected.DiskName {
		checkImage(newMachine)
	}

	newMachine.Name = name

	// Save new machine to old machine path
	saveToDisk(newMachine, getMachinePath(selected))
}
