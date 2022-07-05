package argparser

/*

This file is incharge of parsing the OPTIONS struct to execute the corresponding function.

The most important function here is the `ParseOptions()` function
It has a if statement that executes a helper function for each of the command.

This helper function takes care of logging anything , parsing and finally calling the function responsible for that action.
These function are named {{ .Command}}Execute()

Some {{ .Command}}Options need to be converted to respective structs for future processing.
These are done by helper functions. These are named as {{ .Command}}OptionsTo{{ .RequiredStruct }}().

Here {{ .RequiredStruct }} is the name of the struct in Camel Case.

*/
import (
	"log"

	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/creator"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/console"
	"github.com/pspiagicw/qemantra/pkg/runner"
)

func ParseOptions(global *Flags, version string) {
	switch {
	case createMachineCommand.Used:
		createMachineExecute(global.createMachineFlags)
	case createImgCommand.Used:
		createImgExecute(global.createImgFlags)
	case runCommand.Used:
		runExecute(global.runFlags)
	case listCommand.Used:
		listExecute(global.listFlags)
	case checkCommand.Used:
		config.PerformCheck()
	case renameCommand.Used:
		renameExecute(global.renameFlags)
	case editCommand.Used:
		editExecute(global.editFlags)
	default:
		console.ShowBanner(version)

	}
}
func RunOptionsToRunner(option *RunFlags, runner *runner.Runner) {
	if option.iso != "" {
		runner.Iso = option.iso
	}
	if option.externaldisk != "" {
		runner.ExternalDisk = option.externaldisk
	}
	if option.boot != "" {
		runner.Boot = option.boot
	}
	if option.uefi != false {
		config.EnsureUEFIReady()
		runner.UEFI = true
	}
	if option.no_kvm != false {
		runner.KVM = true
	}
}

// -- CREATE MACHINE
func createMachineExecute(options *CreateMachineFlags) {
	log.Println("Creating a new machine!")
	cr := (*creator.Machine)(options)
	creator.CreateMachine(cr)
}

// -- CREATE IMG
func createImgExecute(options *CreateImgFlags) {
	log.Println("Creating a new image!")
	im := (*image.Image)(options)
	_, err := image.CreateImage(im)
	if err != nil {
		log.Fatalf("Error creating image %v", err)
	}
}

// -- RUN
func runExecute(options *RunFlags) {
	log.Println("Finding the given machine!")
	machine := runner.FindMachine(options.name, true)

	if machine == nil {
		log.Fatalf("Machine %s not found", options.name)
	}
	RunOptionsToRunner(options, machine)
	runner.RunMachine(machine)
}

// -- LIST
func listExecute(options *ListFlags) {
	if options.Img {
		runner.ListImages(options.Verbose)
	} else {
		runner.ListMachines(options.Verbose)
	}
}

// -- RENAME
func renameExecute(options *RenameFlags) {
	oldName := options.OldName
	newName := options.NewName
	runner.RenameMachine(oldName, newName)
}

// -- EDIT
func editExecute(options *EditFlags) {
	log.Println("Finding the given Machine")

	runner := runner.FindMachine(options.Name, false)

	if runner == nil {
		log.Fatalf("Machine %s not found", options.Name)
	}

	machine := (*creator.Machine)(options)
	creator.EditMachine(machine, runner)
}
