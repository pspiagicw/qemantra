package runner

import (
	"fmt"
	"log"

	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/dirs"
)

func FindMachine(name string) *Runner {
	config := config.GetConfig()
	for _, file := range dirs.ListDirs(config.MachineDir) {
		filepath := getFileName(file)
		runner , ok := checkName(filepath , name)
		if ok {
			return runner
		}
	}
	return nil
}
func ListMachines() {
	machines := getMachineList()
	for i , runner := range machines {
		fmt.Printf("%d) Name: %s\n" , i , runner.Name)
	}
}
func getMachineList() []Runner {
	config := config.GetConfig()
	runners := make([]Runner , 0)
	for _ , file := range dirs.ListDirs(config.MachineDir) {
		filepath := getFileName(file)
		runner , err := decodeFileToRunner(filepath)
		if err != nil {
			log.Fatalf("Can't parse %s , %v" , filepath , err)
		}
		runners = append(runners , *runner)

	}
	return runners
}
