// Package enables running virtual machines
package run

/*
This package is incharge of finding and running virtual machines.
The most important function of this package is `RunMachine`.
This function initializes the `Runner` struct and then uses it to construct a command.

This command is executed using `executor` package.

*/
import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/execute"
	"github.com/pspiagicw/qemantra/pkg/image"
)

type argumentGenerator func(runner *Runner) []string

const menuBoot string = "menu=on"
const isoBoot string = "d"

const OVMF_PATH = "/usr/share/ovmf/x64/OVMF.fd"

var ExecProvider = execute.GetExecutor()

type Runner struct {
	Name          string `json:"name"`
	DrivePath     string `json:"drivePath"`
	SystemCommand string `json:"systemCommand"`
	MemSize       string `json:"memSize"`
	CpuCores      string `json:"cpuCores"`
	Iso           string `json:"-"`
	ExternalDisk  string `json:"-"`
	Boot          string `json:"-"`
	UEFI          bool   `json:"-"`
	KVM           bool   `json:"-"`
}

func RunMachine(runner *Runner) {
	checkExternalDisk(runner)
	startMachine(runner)
}
func checkExternalDisk(runner *Runner) {
	if runner.ExternalDisk != "" {
		fullpath := image.FindImage(runner.ExternalDisk)
		if fullpath == "" {
			log.Fatalf("Can't find disk with name: '%s'", runner.ExternalDisk)
		}
		log.Printf("Disk Found! Using '%s'", fullpath)
		runner.ExternalDisk = fullpath
	}
}
func getGenerators() []argumentGenerator {
	var argumentOrder []argumentGenerator

	argumentOrder = append(argumentOrder, generateMemArguments)
	argumentOrder = append(argumentOrder, generateKVMArguments)
	argumentOrder = append(argumentOrder, generateISOArguments)
	argumentOrder = append(argumentOrder, generateDriveArguments)
	argumentOrder = append(argumentOrder, generateBootArguments)
	argumentOrder = append(argumentOrder, generateCPUArguments)
	argumentOrder = append(argumentOrder, generateUEFIArguments)
	argumentOrder = append(argumentOrder, generateExternalDiskArguments)
	return argumentOrder
}
func startMachine(runner *Runner) {
	arguments := constructArguments(runner)

	err := ExecProvider.Execute(generateSystemCommand(runner), arguments)
	if err != nil {
		log.Printf("Some error occured: %v", err)
	}
}
func constructArguments(runner *Runner) []string {
	arguments := []string{}

	generators := getGenerators()
	for i := 0; i < len(generators); i++ {
		arguments = append(arguments, generators[i](runner)...)
	}
	return arguments
}
func generateUEFIArguments(runner *Runner) []string {
	if runner.UEFI {
		return []string{"-bios", OVMF_PATH}
	}
	return []string{}
}
func generateISOArguments(runner *Runner) []string {
	if runner.Iso != "" {
		option := []string{"-cdrom", runner.Iso}
		return option
	}
	return []string{}

}
func generateDriveArguments(runner *Runner) []string {
	if runner.DrivePath != "" {
		option := []string{"-hda", runner.DrivePath}
		return option
	}
	return []string{}
}
func generateMemArguments(runner *Runner) []string {
	if runner.MemSize != "" {
		return []string{"-m", runner.MemSize}
	}
	return []string{}
}
func generateCPUArguments(runner *Runner) []string {
	if runner.CpuCores != "" {
		return []string{"-cpu", "host", "-smp", runner.CpuCores}
	}
	return []string{"-cpu", "host"}
}
func generateKVMArguments(runner *Runner) []string {
	if runner.KVM {
		return []string{}
	}
	return []string{"-enable-kvm"}
}
func generateBootArguments(runner *Runner) []string {
	if runner.Boot == "menu" {
		return []string{"-boot", menuBoot}
	} else if runner.Boot == "iso" {
		return []string{"-boot", isoBoot}
	}
	return []string{}

}
func generateExternalDiskArguments(runner *Runner) []string {
	if runner.ExternalDisk != "" {
		return []string{"-hdb", runner.ExternalDisk}

	}
	return []string{}
}
func generateSystemCommand(runner *Runner) string {
	if runner.SystemCommand == "" {
		return "qemu-system-x86_64"
	}
	return runner.SystemCommand
}
