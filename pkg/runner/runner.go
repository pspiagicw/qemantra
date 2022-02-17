package runner

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/pspiagicw/lazyqemu/pkg/config"
)

type Runner struct {
	Name          string `json:"name"`
	DrivePath     string `json:"drivePath"`
	SystemCommand string `json:"systemCommand"`
	MemSize       string `json:"memSize"`
	IsoPath       string `json:"isoPath"`
}

func getFileName(machineDir string, file fs.FileInfo) string {
	path := filepath.Join(machineDir, file.Name())
	return path
}

func RunMachine(config *config.Config, file fs.FileInfo) {
	runner := Runner{}
	fileName := getFileName(config.MachineDir, file)
	filecontents := readFile(fileName)
	err := json.Unmarshal([]byte(filecontents), &runner)
	if err != nil {
		log.Fatalf("Cannot stored machine in %s", file.Name())
	}
	startMachine(&runner)
}
func startMachine(runner *Runner) {
	cmd := exec.Command(runner.SystemCommand, runner.DrivePath, getMemOptions(runner), getCpuOptions(runner), getMiscOptions(runner), getBoolOptions(runner))
	fmt.Println(cmd)
}
func getMemOptions(runner *Runner) string {
	return "-m 4G"
}
func getCpuOptions(runner *Runner) string {
	return ""
}
func getMiscOptions(runner *Runner) string {
	return ""
}
func getBoolOptions(runner *Runner) string {
	return ""
}
func readFile(file string) string {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot read file %s , %v", file, err)
	}
	return string(contents)

}
