package manage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/dir"
	"github.com/pspiagicw/qemantra/pkg/run"
)

const MOST_RECENT_FILE = "recentf"

func FindMachine(name string, useCache bool) *run.Runner {
	if useCache {
		if name == "" {
			name = findMostRecentMachine()
		}
		storeMostRecentMachine(name)
	}
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		runner, matches := ifNameMatches(filepath, name)
		if matches {
			return runner
		}
	}
	return nil
}

func decodeBytesToRunner(contents []byte) (*run.Runner, error) {
	var runner run.Runner
	err := json.Unmarshal(contents, &runner)
	if err != nil {
		return nil, err
	}
	return &runner, nil
}
func LoadRunnerFromDisk(filepath string) (*run.Runner, error) {
	contents, err := readFile(filepath)
	if err != nil {
		return nil, err
	}
	runner, err := decodeBytesToRunner(contents)
	if err != nil {
		return nil, err
	}
	return runner, nil
}
func ifNameMatches(filepath string, name string) (*run.Runner, bool) {
	runner, err := LoadRunnerFromDisk(filepath)

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
	ioutil.WriteFile(filename, []byte(name), 0644)

}
func findMostRecentMachine() string {
	configdir := ConfigProvider.GetConfigDir()
	filename := filepath.Join(configdir, MOST_RECENT_FILE)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Can't read the recent file %s , %v", filename, err)
	}
	return string(contents)

}
func ListMachines(verbose bool) {
	machines := getRunnerList()

	for i, runner := range machines {
		fmt.Printf("%d) Name: %s\n", i+1, runner.Name)
		if verbose {
			fmt.Printf("    MemSize: %s\n", runner.MemSize)
			fmt.Printf("    CpuCores: %s\n", runner.CpuCores)
			fmt.Printf("    DrivePath: %s\n", runner.DrivePath)
		}
	}

}

func getRunnerList() []run.Runner {
	runners := make([]run.Runner, 0)
	for _, file := range dir.ListDir(ConfigProvider.GetMachineDir()) {
		filepath := getRunnerPath(file.Name())
		runner, err := LoadRunnerFromDisk(filepath)
		if err != nil {
			log.Fatalf("Can't parse %s , %v", filepath, err)
		}
		runners = append(runners, *runner)

	}
	return runners
}
func readFile(file string) ([]byte, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return contents, nil
}
