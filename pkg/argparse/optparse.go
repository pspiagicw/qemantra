// MIT License
//
// Copyright (c) 2022 pspiagicw
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package argparse

/*
This file is incharge of parsing the OPTIONS struct to execute the corresponding function.

*/
import (
	"flag"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	log "github.com/pspiagicw/colorlog"
	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/console"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/machine"
	"github.com/pspiagicw/qemantra/pkg/manage"
	runner "github.com/pspiagicw/qemantra/pkg/run"
)

func ParseOptions(version string) {

	showVersion := flag.Bool("version", false, "Show version info")
	showVerbose := flag.Bool("verbose", false, "Shoe verbose info")

	flag.Parse()

	if *showVersion {
		console.ShowBanner(version)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		console.ShowBanner(version)
		console.ShowSubcommands()
		os.Exit(0)
	}

	args := flag.Args()

	cmd, args := flag.Args()[0], flag.Args()[1:]

	switch cmd {
	case "edit":
		edit(args, *showVerbose)
	case "create":
		create(args, *showVerbose)
	case "list":
		list(args, *showVerbose)
	case "run":
		run(args, *showVerbose)
	case "rename":
		rename(args, *showVerbose)
	case "check":
		config.PerformCheck()
	default:
		console.ShowBanner(version)
		console.ShowSubcommands()

	}

}
func run(args []string, verbose bool) {

	flag := flag.NewFlagSet("qemantra run", flag.ExitOnError)

	iso := flag.String("iso", "", "Path of the ISO to attach")
	boot := flag.String("boot", "iso", "Boot order")
	kvm := flag.Bool("kvm", true, "Enable KVM")
	externaldisk := flag.String("external", "", "External disk to attach")
	uefi := flag.String("uefi", "", "Path to OVMF(.fd) file")

	flag.Parse(args)

	machines := manage.ListMachines(verbose)

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	m := manage.FindMachine(name)

	m.Iso = *iso
	m.Boot = *boot
	m.KVM = *kvm
	m.UEFI = *uefi
	m.ExternalDisk = *externaldisk

	runner.RunMachine(m)
}

func list(args []string, verbose bool) {
	flag := flag.NewFlagSet("qemantra list", flag.ExitOnError)

	_ = flag.Bool("image", false, "List images instead of machines.")

	flag.Parse(args)

	machines := manage.ListMachines(verbose)

	for i, runner := range machines {
		fmt.Printf("%d) Name: %s\n", i+1, runner.Name)
		if verbose {
			fmt.Printf("    MemSize: %s\n", runner.MemSize)
			fmt.Printf("    CpuCores: %s\n", runner.CpuCores)
			fmt.Printf("    DrivePath: %s\n", runner.DrivePath)
		}
	}
}

func buildMachine() *machine.Machine {
	machine := &machine.Machine{}

	machine.CpuCores = userPrompt("CPU Cores", coresValidator)
	machine.MemSize = userPrompt("RAM Size", ramValidator)
	machine.SystemCommand = userSelection("System Command", []string{
		"qemu-system-x86_64",
		"qemu-system-i386",
	})
	wantDisk := userPrompt("Do you want to attach disk? (Y/N)", func(string) error { return nil })

	if wantDisk == "y" || wantDisk == "Y" {
		machine.NoDisk = false
		machine.DiskName = userPrompt("Disk Name", func(string) error { return nil })
		machine.DiskSize = userPrompt("Disk Size", ramValidator)
		choices := []string{"raw", "vdi", "qcow2"}

		machine.DiskFormat = userSelection("Disk Format", choices)
	} else {
		machine.NoDisk = true
	}

	return machine

}
func create(options []string, verbose bool) {

	flag := flag.NewFlagSet("qemantra create", flag.ExitOnError)

	flag.Parse(options)
	log.LogInfo("Creating a new machine!")
	name := userPrompt("Name", func(string) error { return nil })
	machine := buildMachine()
	machine.Name = name

	manage.CreateMachine(machine)
}

func rename(args []string, verbose bool) {

	flag := flag.NewFlagSet("qemantra rename", flag.ExitOnError)

	flag.Parse(args)

	machines := manage.ListMachines(verbose)

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	newName := userPrompt("New Name", func(string) error { return nil })

	manage.RenameMachine(name, newName)

}

func edit(args []string, verbose bool) {

	flag := flag.NewFlagSet("qemantra edit", flag.ExitOnError)

	flag.Parse(args)
	machines := manage.ListMachines(verbose)

	choices := []string{}

	for _, m := range machines {
		choices = append(choices, m.Name)
	}

	name := userSelection("Select Machine", choices)

	newMachine := manage.FindMachine(name)

	if newMachine == nil {
		log.LogFatal("Machine %s not found", name)
	}

	m := buildMachine()

	if m.DiskName != newMachine.DiskName && m.DiskName != "" {
		fmt.Println("Need to create a new disk")
		img := &image.Image{
			Type: m.DiskFormat,
			Name: m.DiskName,
			Size: m.DiskSize,
		}
		path, err := image.CreateImage(img)
		if err != nil {
			log.LogFatal("Could not create disk: %v", err)
		}
		m.DrivePath = path
	}

	m.Name = name

	manage.EditMachine(m)
}

func userPrompt(label string, validation func(string) error) string {
	prompt := promptui.Prompt{Label: label, Validate: validation}

	value, err := prompt.Run()

	if err != nil {
		log.LogFatal("Something went wrong: %q", err)
	}

	return value
}
func userSelection(label string, choices []string) string {

	prompt := promptui.Select{Label: label, Items: choices}

	_, value, err := prompt.Run()

	if err != nil {
		log.LogFatal("Something went wrong: %q", err)
	}

	return value
}
