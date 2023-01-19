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
	"github.com/pspiagicw/qemantra/pkg/machine"
)

type argumentGenerator func(*machine.Machine) []string

const menuBoot string = "menu=on"
const isoBoot string = "d"

const OVMF_PATH = "/usr/share/ovmf/x64/OVMF.fd"

var ExecProvider = execute.GetExecutor()

func RunMachine(machine *machine.Machine) {
	checkExternalDisk(machine)
	startMachine(machine)
}
func checkExternalDisk(machine *machine.Machine) {
	if machine.ExternalDisk != "" {
		fullpath := image.FindImage(machine.ExternalDisk)
		if fullpath == "" {
			log.Fatalf("Can't find disk with name: '%s'", machine.ExternalDisk)
		}
		log.Printf("Disk Found! Using '%s'", fullpath)
		machine.ExternalDisk = fullpath
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
func startMachine(machine *machine.Machine) {
	arguments := constructArguments(machine)

	err := ExecProvider.Execute(generateSystemCommand(machine), arguments)
	if err != nil {
		log.Printf("Some error occured: %v", err)
	}
}
func constructArguments(machine *machine.Machine) []string {
	arguments := []string{}

	generators := getGenerators()
	for i := 0; i < len(generators); i++ {
		arguments = append(arguments, generators[i](machine)...)
	}
	return arguments
}
func generateUEFIArguments(machine *machine.Machine) []string {
	if machine.UEFI != "" {
		return []string{"-bios", machine.UEFI}
	}
	return []string{}
}
func generateISOArguments(machine *machine.Machine) []string {
	if machine.Iso != "" {
		option := []string{"-cdrom", machine.Iso}
		return option
	}
	return []string{}

}
func generateDriveArguments(machine *machine.Machine) []string {
	if machine.DrivePath != "" {
		option := []string{"-hda", machine.DrivePath}
		return option
	}
	return []string{}
}
func generateMemArguments(machine *machine.Machine) []string {
	if machine.MemSize != "" {
		return []string{"-m", machine.MemSize}
	}
	return []string{}
}
func generateCPUArguments(machine *machine.Machine) []string {
	if !machine.KVM {
		return []string{}
	}
	if machine.CpuCores != "" {
		return []string{"-cpu", "host", "-smp", machine.CpuCores}
	}
	return []string{"-cpu", "host"}
}
func generateKVMArguments(machine *machine.Machine) []string {
	if machine.KVM {
		return []string{"-enable-kvm"}
	}
	return []string{}
}
func generateBootArguments(machine *machine.Machine) []string {
	if machine.Boot == "menu" {
		return []string{"-boot", menuBoot}
	} else if machine.Boot == "iso" {
		return []string{"-boot", isoBoot}
	}
	return []string{}

}
func generateExternalDiskArguments(machine *machine.Machine) []string {
	if machine.ExternalDisk != "" {
		return []string{"-hdb", machine.ExternalDisk}

	}
	return []string{}
}
func generateSystemCommand(machine *machine.Machine) string {
	if machine.SystemCommand == "" {
		return "qemu-system-x86_64"
	}
	return machine.SystemCommand
}
