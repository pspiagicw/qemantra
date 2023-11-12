package argparse

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/manifoldco/promptui"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/machine"
	"github.com/pspiagicw/qemantra/pkg/manage"
	runner "github.com/pspiagicw/qemantra/pkg/run"
)

func check(args []string) {
	config.PerformCheck()
}

func run(args []string) {

	flag := flag.NewFlagSet("qemantra run", flag.ExitOnError)

	iso := flag.String("iso", "", "Path of the ISO to attach")
	boot := flag.String("boot", "iso", "Boot order")
	kvm := flag.Bool("kvm", true, "Enable KVM")
	externaldisk := flag.String("external", "", "External disk to attach")
	uefi := flag.String("uefi", "", "Path to OVMF(.fd) file")

	flag.Parse(args)

	machines := manage.ListMachines()

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	goreland.LogInfo("Machine %s Selected", name)

	m := manage.FindMachine(name)

	m.Iso = *iso
	m.Boot = *boot
	m.KVM = *kvm
	m.UEFI = *uefi
	m.ExternalDisk = *externaldisk

	runner.RunMachine(m)

	goreland.LogSuccess("Machine successfully ran!")
}

func list(args []string) {
	flag := flag.NewFlagSet("qemantra list", flag.ExitOnError)

	_ = flag.Bool("image", false, "List images instead of machines.")

	flag.Parse(args)

	machines := manage.ListMachines()

	if len(machines) == 0 {
		goreland.LogFatal("No virtual machines created!")
	}

	headers := []string{"Name", "Cores", "Memory", "Drive"}

	rows := [][]string{}
	for _, machine := range machines {

		// fmt.Printf("\t\"%s\"\t\tCores: %s, Memory: %s, Disk: %s\n", machine.Name, machine.CpuCores, machine.MemSize , machine.DrivePath)
		rows = append(rows, []string{machine.Name, machine.CpuCores, machine.MemSize, machine.DrivePath})
	}
	goreland.LogTable(headers, rows)
}
func generatePrompt(prompt string, template *machine.Machine, key string) string {

	if template == nil {
		return prompt
	}

	value := reflect.Indirect(reflect.ValueOf(template)).FieldByName(key).String()

	newPrompt := fmt.Sprintf("%s (%s)", prompt, value)

	return newPrompt

}

func buildMachine(template *machine.Machine) *machine.Machine {
	machine := &machine.Machine{}

	machine.CpuCores = userPrompt(generatePrompt("CPU Cores", template, "CpuCores"), coresValidator)
	machine.MemSize = userPrompt(generatePrompt("RAM Size", template, "MemSize"), ramValidator)
	machine.SystemCommand = userSelection(generatePrompt("System Command", template, "SystemCommand"), []string{
		"qemu-system-x86_64",
		"qemu-system-i386",
	})

	wantDisk := userPrompt(generatePrompt("Do you want to attach disk? (Y/N)", template, "NoDisk"), func(string) error { return nil })

	if wantDisk == "y" || wantDisk == "Y" {
		machine.NoDisk = false
		machine.DiskName = userPrompt(generatePrompt("Disk Name", template, "DiskName"), func(string) error { return nil })
		machine.DiskSize = userPrompt(generatePrompt("Disk Size", template, "DiskSize"), ramValidator)
		choices := []string{"raw", "vdi", "qcow2"}

		machine.DiskFormat = userSelection(generatePrompt("Disk Format", template, "DiskFormat"), choices)
	} else {
		machine.NoDisk = true
	}

	return machine

}
func create(options []string) {

	flag := flag.NewFlagSet("qemantra create", flag.ExitOnError)

	flag.Parse(options)

	goreland.LogInfo("Creating a new machine.")

	name := userPrompt("Name", func(string) error { return nil })

	machine := buildMachine(nil)

	machine.Name = name

	manage.CreateMachine(machine)

	goreland.LogSuccess("Created machine successfully!")
}

func rename(args []string) {

	flag := flag.NewFlagSet("qemantra rename", flag.ExitOnError)

	flag.Parse(args)

	machines := manage.ListMachines()

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	goreland.LogInfo("Selected Machine %s", name)

	newName := userPrompt("New Name", func(string) error { return nil })

	manage.RenameMachine(name, newName)

	goreland.LogSuccess(" Machine renamed")

}

func edit(args []string) {

	flag := flag.NewFlagSet("qemantra edit", flag.ExitOnError)

	flag.Parse(args)
	machines := manage.ListMachines()

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	goreland.LogInfo("Machine %s selected", name)

	newMachine := manage.FindMachine(name)

	if newMachine == nil {
		goreland.LogFatal("Machine %s not found", name)
	}

	m := buildMachine(newMachine)

	if m.DiskName != newMachine.DiskName && m.DiskName != "" {
		goreland.LogInfo("Disk change detected!")
		img := &image.Image{
			Type: m.DiskFormat,
			Name: m.DiskName,
			Size: m.DiskSize,
		}
		path, err := image.CreateImage(img)
		if err != nil {
			goreland.LogFatal("Could not create disk: %v", err)
		}
		goreland.LogInfo("Created disk at %s", path)
		m.DrivePath = path
	}

	m.Name = name

	manage.EditMachine(m)

	goreland.LogInfo("Edited machine successfully!")
}

func userPrompt(label string, validation func(string) error) string {
	prompt := promptui.Prompt{Label: label, Validate: validation}

	value, err := prompt.Run()

	if err != nil {
		goreland.LogFatal("Something went wrong: %q", err)
	}

	return value
}
func userSelection(label string, choices []string) string {

	prompt := promptui.Select{Label: label, Items: choices}

	_, value, err := prompt.Run()

	if err != nil {
		goreland.LogFatal("Something went wrong: %q", err)
	}

	return value
}
