package runner

import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/executor"
	"github.com/pspiagicw/qemantra/pkg/image"
)

const menuBoot string = "menu=on"
const isoBoot string = "d"
var ExecProvider = executor.GetExecutor()

type Runner struct {
	Name          string `json:"name"`
	DrivePath     string `json:"drivePath"`
	SystemCommand string `json:"systemCommand"`
	MemSize       string `json:"memSize"`
	CpuCores      string `json:"cpuCores"`
	Iso           string `json:"-"`
	ExternalDisk  string `json:"-"`
	Boot          string `json:"-"`
}

func RunMachine(runner *Runner) {
	if runner.ExternalDisk != ""{
		fullpath := image.FindImage(runner.ExternalDisk)
		if fullpath == "" {
			log.Fatalf("Can't find disk with name '%s'" , runner.ExternalDisk)
		}
		log.Printf("Disk Found! Using '%s'" , fullpath)
		runner.ExternalDisk = fullpath
	}

	startMachine(runner)
}
func addExternalDisk(runner *Runner) {
	
}
func startMachine(runner *Runner) {
	options := constructOptions(runner)
	// cmd := exec.Command(runner.SystemCommand, options...)

	// log.Printf("Executing '%s' command on your system", cmd)
	// var out bytes.Buffer
	// cmd.Stderr = &out

	// err := cmd.Run()
	err := ExecProvider.Execute(runner.SystemCommand , options)
	if err != nil {
		log.Printf("Some error occured %v", err)
	}
}
func constructOptions(runner *Runner) []string {
	options := []string{}
	options = append(options, getMemOptions(runner)...)
	options = append(options, getMiscOptions(runner)...)
	options = append(options, getIsoOptions(runner)...)
	options = append(options, getDriveOptions(runner)...)
	options = append(options, getBootOptions(runner)...)
	options = append(options ,getCpuOptions(runner)...)
	options = append(options, getExternalDiskOption(runner)...)
	return options
}
func getIsoOptions(runner *Runner) []string {
	if runner.Iso != "" {
		option := []string{"-cdrom", runner.Iso}
		return option
	}
	return []string{}

}
func getDriveOptions(runner *Runner) []string {
	if runner.DrivePath != "" {
		option := []string{"-hda", runner.DrivePath}
		return option
	}
	return []string{}
}
func getMemOptions(runner *Runner) []string {
	if runner.MemSize != "" {
		return []string{"-m", runner.MemSize}
	}
	return []string{}
}
func getCpuOptions(runner *Runner) []string {
	if runner.CpuCores != "" {
		return []string{"-smp", runner.CpuCores}
	}
	return []string{}
}
func getMiscOptions(runner *Runner) []string {
	return []string{"-enable-kvm"}
}
func getBootOptions(runner *Runner) []string {
	if runner.Boot == "menu" {
		return []string{"-boot", menuBoot}
	} else if runner.Boot == "iso" {
		return []string{"-boot", isoBoot}
	}
	return []string{}

}
func getExternalDiskOption(runner *Runner) []string {
	if runner.ExternalDisk != "" {
		return []string{"-hdb", runner.ExternalDisk}

	}
	return []string{}
}
