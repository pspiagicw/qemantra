package runner

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/dirs"
)

const MOST_RECENT_FILE = "recentf"

func FindMachine(name string, useCache bool) *Runner {
	if useCache {
		if name == "" {
			name = findMostRecentMachine()
		}
		storeMostRecentMachine(name)
	}
	for _, file := range dirs.ListDirs(ConfigProvider.GetMachineDir()) {
		filepath := getFileName(file)
		runner, ok := checkName(filepath, name)
		if ok {
			return runner
		}
	}
	return nil
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
	machines := getMachineList()
	for i, runner := range machines {
		fmt.Printf("%d) Name: %s\n", i+1, runner.Name)
		if verbose {
			fmt.Printf("    MemSize: %s\n", runner.MemSize)
			fmt.Printf("    CpuCores: %s\n", runner.CpuCores)
			fmt.Printf("    DrivePath: %s\n", runner.DrivePath)
		}
	}

}
func ListImages(verbose bool) {
	images := getImageList()
	for i, image := range images {
		fmt.Printf("%d) Path: %s\n", i, image)
	}
}

func getMachineList() []Runner {
	runners := make([]Runner, 0)
	for _, file := range dirs.ListDirs(ConfigProvider.GetMachineDir()) {
		filepath := getFileName(file)
		runner, err := decodeFileToRunner(filepath)
		if err != nil {
			log.Fatalf("Can't parse %s , %v", filepath, err)
		}
		runners = append(runners, *runner)

	}
	return runners
}

func getImageList() []string {
	paths := make([]string, 0)
	for _, file := range dirs.ListDirs(ConfigProvider.GetImageDir()) {
		filepath := getFileName(file)
		paths = append(paths, filepath)
	}
	return paths

}
