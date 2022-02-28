// Package to parse arguments passed to qemantra
package argparser

import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/creator"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/prompt"
	"github.com/pspiagicw/qemantra/pkg/runner"
)

// Check the parsed arguments and run corresponding actions
// The input is Options which contains all the options passed.
func ParseAndRun(globalOptions *Options , version string) {
	if globalOptions.CreateMachineCommand.Used {
		log.Println("Creating a new machine!")
		cr := globalOptionsToMachineCreator(globalOptions)
		creator.CreateNewMachine(cr)
	} else if globalOptions.CreateImgCommand.Used {
		log.Println("Creating a new image!")
		im := globalOptionsToImage(globalOptions)
		image.CreateImage(im)
	} else if globalOptions.RunOptionCommand.Used {
		log.Println("Finding the given machine!")
		machine := runner.FindMachine(globalOptions.RunOptions.name)
		if machine == nil {
			log.Fatalf("Machine %s not found", globalOptions.RunOptions.name)
		}
		addRunnerOptions(globalOptions.RunOptions , machine)
		runner.RunMachine(machine)

	} else if globalOptions.ListCommand.Used {
		runner.ListMachines(globalOptions.ListOptions.Img)
	} else {
		prompt.ShowBanner(version)
	}
}
func addRunnerOptions(option *RunCommandOptions , runner *runner.Runner) {
	if option.iso != "" {
		runner.Iso = option.iso
	}
	if option.externaldisk != "" {
		runner.ExternalDisk = option.externaldisk
	}
}

// Convert given arguments(as options) to instance of MachineCreator
// Machine Creator can be used to create a machine
func globalOptionsToMachineCreator(globalOptions *Options) *creator.MachineCreator {
	cr := &creator.MachineCreator{
		Name:       globalOptions.CreateMachineOptions.Name,
		NoDisk:     globalOptions.CreateMachineOptions.NoDisk,
		MemSize:    globalOptions.CreateMachineOptions.MemSize,
		CpuCores:   globalOptions.CreateMachineOptions.CpuCores,
		DiskName:   globalOptions.CreateMachineOptions.DiskName,
		DiskFormat: globalOptions.CreateMachineOptions.DiskFormat,
		DiskSize:   globalOptions.CreateMachineOptions.DiskSize,
	}
	return cr

}

// Convert given arguments(as options) to instance of Image
// Image struct can be used to create a new image.
func globalOptionsToImage(globalOptions *Options) *image.Image {
	im := &image.Image{
		Type: globalOptions.CreateImgOptions.Format,
		Name: globalOptions.CreateImgOptions.Name,
		Size: globalOptions.CreateImgOptions.Size,
	}
	return im
}
