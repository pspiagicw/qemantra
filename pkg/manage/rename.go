package manage

import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/dir"
)

func RenameMachine(oldname string, newname string) {
	run := FindMachine(oldname, false)

	if run == nil {
		log.Fatalf("Machine %s not found! ", oldname)
	}

	newRun := FindMachine(newname, false)
	if newRun != nil {
		log.Fatalf("Machine with name %s already exists!", newname)
	}

	filepath := FindMachineFile(oldname)

	ReplaceName(filepath, newname)
	// Add RenameFile(filepath , newname)
	// os.Remove(filepath)
}

func FindMachineFile(name string) string {
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		_, ok := ifNameMatches(filepath, name)
		if ok {
			return filepath
		}
	}
	return ""
}

func ReplaceName(path string, newname string) {
	runner, err := LoadRunnerFromDisk(path)
	if err != nil {
		log.Fatalf("Error reading %s file", path)

	}
	runner.Name = newname
	err = SaveRunnerToDisk(runner)
	if err != nil {
		log.Fatalf("Error updating %q file with new name", err)
	}
}
