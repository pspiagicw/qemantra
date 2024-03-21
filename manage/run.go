package manage

import (
	"encoding/xml"
	"flag"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/config"
	"github.com/pspiagicw/qemantra/help"
	prompt "github.com/pspiagicw/qemantra/prompts"
	"github.com/pspiagicw/qemantra/vm"
)

type argumentGenerator func(*vm.VirtualMachine) []string

const menuBoot string = "menu=on"
const isoBoot string = "d"

type Opts struct {
	kvm   bool
	boot  string
	iso   string
	uefi  string
	edisk string
}

func parseRunArgs(args []string) *Opts {
	flag := flag.NewFlagSet("qemantra run", flag.ExitOnError)

	opts := new(Opts)

	flag.BoolVar(&opts.kvm, "kvm", true, "Enable KVM (default true)")
	flag.StringVar(&opts.boot, "boot", "iso", "Boot order")
	flag.StringVar(&opts.iso, "iso", "", "Path of ISO to boot")
	flag.StringVar(&opts.uefi, "uefi", "", "Path of OVMF (.fd) file.")
	flag.StringVar(&opts.edisk, "external-disk", "", "Path to external disk")

	flag.Usage = help.HelpRun

	flag.Parse(args)

	return opts

}

func RunVM(args []string) {

	opts := parseRunArgs(args)

	_, selected := selectMachine()

	selected.KVM = opts.kvm
	selected.Boot = opts.boot
	selected.ISO = opts.iso
	selected.UEFI = opts.uefi
	selected.ExternalDisk = opts.edisk

	runMachine(selected)
}
func selectMachine() (string, *vm.VirtualMachine) {
	machines := getMachines()

	choices := []string{}

	for name, _ := range machines {
		choices = append(choices, name)
	}

	selected := prompt.SelectionPrompt("Select Machine", choices)

	selectedMachine := machines[selected]

	return selected, selectedMachine

}
func runMachine(m *vm.VirtualMachine) {

	cmd := getMachineCommand(m)

	goreland.LogExec(cmd.String())

	executeCommand(cmd)
}
func getMachineCommand(m *vm.VirtualMachine) *exec.Cmd {
	args := getMachineArgs(m)

	cmd := exec.Command(m.Architecture, args...)

	return cmd
}
func getMachineArgs(m *vm.VirtualMachine) []string {
	arguments := []string{}

	generators := getGenerators()

	for i := 0; i < len(generators); i++ {
		arguments = append(arguments, generators[i](m)...)
	}
	return arguments
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
func getMachines() map[string]*vm.VirtualMachine {

	files, err := os.ReadDir(config.MachineDir())

	if err != nil {
		goreland.LogFatal("Error reading machines.: %v", err)
	}

	machines := map[string]*vm.VirtualMachine{}

	for _, file := range files {
		path := filepath.Join(config.MachineDir(), file.Name())
		if !file.IsDir() {
			machine := readMachine(path)
			machines[machine.Name] = machine
		}
	}

	return machines
}
func readFile(path string) []byte {
	contents, err := os.ReadFile(path)
	if err != nil {
		goreland.LogFatal("Error reading file '%s': %v", path, err)
	}
	return contents
}
func readMachine(path string) *vm.VirtualMachine {
	contents := readFile(path)

	machine := new(vm.VirtualMachine)

	err := xml.Unmarshal(contents, machine)
	if err != nil {
		goreland.LogFatal("Error unmarshaling file '%s': %v", path, err)
	}
	return machine
}
func generateCPUArguments(machine *vm.VirtualMachine) []string {
	if !machine.KVM {
		return []string{}
	}
	if machine.CpuCores != "" {
		return []string{"-cpu", "host", "-smp", machine.CpuCores}
	}
	return []string{"-cpu", "host"}
}
func generateKVMArguments(machine *vm.VirtualMachine) []string {
	if machine.KVM {
		return []string{"-enable-kvm"}
	}
	return []string{}
}
func generateBootArguments(machine *vm.VirtualMachine) []string {
	if machine.Boot == "menu" {
		return []string{"-boot", menuBoot}
	} else if machine.Boot == "iso" {
		return []string{"-boot", isoBoot}
	}
	return []string{}

}
func generateExternalDiskArguments(machine *vm.VirtualMachine) []string {
	if machine.ExternalDisk != "" {
		return []string{"-hdb", machine.ExternalDisk}

	}
	return []string{}
}
func generateUEFIArguments(machine *vm.VirtualMachine) []string {
	if machine.UEFI != "" {
		return []string{"-bios", machine.UEFI}
	}
	return []string{}
}
func generateISOArguments(machine *vm.VirtualMachine) []string {
	if machine.ISO != "" {
		option := []string{"-cdrom", machine.ISO}
		return option
	}
	return []string{}

}
func generateDriveArguments(machine *vm.VirtualMachine) []string {
	if machine.DiskPath != "" {
		option := []string{"-hda", machine.DiskPath}
		return option
	}
	return []string{}
}
func generateMemArguments(machine *vm.VirtualMachine) []string {
	if machine.MemSize != "" {
		return []string{"-m", machine.MemSize}
	}
	return []string{}
}
