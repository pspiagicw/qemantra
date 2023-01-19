package manage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/dir"
	"github.com/pspiagicw/qemantra/pkg/machine"
)

const MOST_RECENT_FILE = "recentf"

func FindMachine(name string) *machine.Machine {
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		machine, matches := ifNameMatches(filepath, name)
		if matches {
			return machine
		}
	}
	return nil
}

func decodeBytesToRunner(contents []byte) (*machine.Machine, error) {

	var machine machine.Machine

	err := json.Unmarshal(contents, &machine)

	if err != nil {
		return nil, err
	}

	return &machine, nil
}
func loadMachine(filepath string) (*machine.Machine, error) {
	contents, err := readFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error reading contents: %q", err)
	}
	machine, err := decodeBytesToRunner(contents)
	if err != nil {
		return nil, fmt.Errorf("Can't decode read contents: %q", err)
	}
	return machine, nil
}
func ifNameMatches(filepath string, name string) (*machine.Machine, bool) {
	runner, err := loadMachine(filepath)

	if err != nil {
		log.Fatalf("Can't decode file %s , %v", filepath, err)
	}

	if runner.Name == name {
		return runner, true
	}
	return nil, false
}
func storeMostRecentMachine(name string) {
	configdir := ConfigProvider.GetConfigDir()
	filename := filepath.Join(configdir, MOST_RECENT_FILE)
	os.WriteFile(filename, []byte(name), 0644)

}
func findMostRecentMachine() string {
	configdir := ConfigProvider.GetConfigDir()
	filename := filepath.Join(configdir, MOST_RECENT_FILE)
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Can't read the recent file %s , %v", filename, err)
	}
	return string(contents)

}
func ListMachines(verbose bool) []machine.Machine {
	machines := getMachineList()

	return machines

}

func getMachineList() []machine.Machine {
	runners := make([]machine.Machine, 0)
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		machine, err := loadMachine(filepath)
		if err != nil {
			log.Fatalf("Can't parse %s , %v", filepath, err)
		}
		runners = append(runners, *machine)

	}
	return runners
}
func readFile(file string) ([]byte, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return contents, nil
}
