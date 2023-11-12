package manage

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/pkg/machine"
)

func EditMachine(newMachine *machine.Machine) {
	err := saveToDisk(newMachine)

	if err != nil {
		goreland.LogFatal("Error writing to file , %v", err)
	}
}
