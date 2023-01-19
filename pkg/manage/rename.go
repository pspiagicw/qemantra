package manage

import (
	"log"
	"os"

	"github.com/pspiagicw/qemantra/pkg/dir"
)

func RenameMachine(oldname string, newname string) {
	run := FindMachine(oldname)

	if run == nil {
		log.Fatalf("Machine %s not found! ", oldname)
	}

	newRun := FindMachine(newname)
	if newRun != nil {
		log.Fatalf("Machine with name %s already exists!", newname)
	}

	filepath := findMachineFile(oldname)

	replaceName(filepath, newname)
	os.Remove(filepath)
}

func findMachineFile(name string) string {
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		_, ok := ifNameMatches(filepath, name)
		if ok {
			return filepath
		}
	}
	return ""
}

func replaceName(path string, newname string) {
	runner, err := loadMachine(path)
	if err != nil {
		log.Fatalf("Error reading %s file", path)

	}
	runner.Name = newname
	err = saveToDisk(runner)
	if err != nil {
		log.Fatalf("Error updating %q file with new name", err)
	}
}
