package manage

import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/machine"
)

func EditMachine(newMachine *machine.Machine) {
	err := saveToDisk(newMachine)

	if err != nil {
		log.Fatalf("Error writing to file , %v", err)
	}
}
