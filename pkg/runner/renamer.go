package runner

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pspiagicw/qemantra/pkg/dirs"
)

func RenameMachine(oldname string, newname string) {
	run := FindMachine(oldname , false)
	if run == nil {
		log.Fatalf("Machine %s not found! ", oldname)
	}
	newRun := FindMachine(newname , false)
	if newRun != nil {
		log.Fatalf("Machine with name %s already exists!", newname)
	}
	filepath := FindMachineFile(oldname)
	ReplaceName(filepath, newname)
}
func FindMachineFile(name string) string {
	for _, file := range dirs.ListDirs(ConfigProvider.GetMachineDir()) {
		filepath := getFileName(file)
		_, ok := checkName(filepath, name)
		if ok {
			return filepath
		}
	}
	return ""
}

func ReplaceName(path string, newname string) {
	runner, err := decodeFileToRunner(path)
	if err != nil {
		log.Fatalf("Error reading %s file", path)

	}
	runner.Name = newname
	err = encodeJsonToFile(runner, path)
	if err != nil {
		log.Fatalf("Error updating %q file with new name", err)
	}
}

func encodeJsonToFile(runner *Runner, filepath string) error {
	contents, err := json.Marshal(runner)
	if err != nil {
		return err
	}
	writeFile(contents, filepath)
	return nil

}
func writeFile(contents []byte, filepath string) error {
	err := ioutil.WriteFile(filepath, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
