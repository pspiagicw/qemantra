package argparser

/*

This file is incharge of parsing the OPTIONS struct to execute the corresponding function.

The most important function here is the `ParseOptions()` function
It has a if statement that executes a helper function for each of the command.

This helper function takes care of logging anything , parsing and finally calling the function responsible for that action.
These function are named {{ .Command}}Execute()

Some {{ .Command}}Options need to be converted to respective structs for future processing.
These are done by helper functions. These are named as {{ .Command}}OptionsTo{{ .RequiredStruct }}().
Here {{ .RequiredStruct }} is the name of the struct in Pascal Case.

*/
import (
    "log"

    "github.com/pspiagicw/qemantra/pkg/config"
    "github.com/pspiagicw/qemantra/pkg/creator"
    "github.com/pspiagicw/qemantra/pkg/image"
    "github.com/pspiagicw/qemantra/pkg/prompt"
    "github.com/pspiagicw/qemantra/pkg/runner"
)

func ParseOptions(global *Options, version string) {
    if global.CreateMachineCommand.Used {
        CreateMachineExecute(global.CreateMachineOptions)
   } else if global.CreateImgCommand.Used {
        CreateImgExecute(global.CreateImgOptions)
    } else if global.RunCommand.Used {
        RunExecute(global.RunOptions)
    } else if global.ListCommand.Used {
        ListExecute(global.ListOptions)
    } else if global.CheckCommand.Used {
        config.PerformCheck()
    } else if global.RenameCommand.Used {
        RenameExecute(global.RenameOptions)
    } else if global.EditCommand.Used {
        EditExecute(global.EditOptions)
    } else {
        prompt.ShowBanner(version)
    }
}
func RunOptionsToRunner(option *RunOptions, runner *runner.Runner) {
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
        runner.NO_KVM = true
    }
}

// -- CREATE MACHINE
func CreateMachineExecute(options *CreateMachineOptions) {
        log.Println("Creating a new machine!")
        cr := (*creator.MachineCreator)(options)
        creator.CreateNewMachine(cr)
}
// -- CREATE IMG
func CreateImgExecute(options *CreateImgOptions) {
        log.Println("Creating a new image!")
        im := (*image.Image)(options)
        image.CreateImage(im)
}

// -- RUN 
func RunExecute(options *RunOptions) {
        log.Println("Finding the given machine!")
        machine := runner.FindMachine(options.name)

        if machine == nil {
            log.Fatalf("Machine %s not found", options.name)
        }
        RunOptionsToRunner(options, machine)
        runner.RunMachine(machine)
}

// -- LIST
func ListExecute(options *ListOptions) {
    if options.Img { 
        runner.ListImages(options.Verbose)
    } else {
        runner.ListMachines(options.Verbose)
    }
}

// -- RENAME
func RenameExecute(options *RenameOptions) {
        oldName := options.OldName
        newName := options.NewName
        runner.RenameMachine(oldName, newName)
}

// -- EDIT
func EditExecute(options *EditOptions) {
        log.Println("Finding the given Machine")

        machine := runner.FindMachine(options.Name)

        if machine == nil {
            log.Fatalf("Machine %s not found", options.Name)
        }

        create_machine:= (*creator.MachineCreator)(options)
        creator.EditMachine(create_machine,  machine)
}
