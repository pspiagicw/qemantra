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

The most important function here is the `ParseOptions()` function
It has a if statement that executes a helper function for each of the command.

This helper function takes care of logging anything , parsing and finally calling the function responsible for that action.
These function are named {{ .Command}}Execute()

Some {{ .Command}}Options need to be converted to respective structs for future processing.
These are done by helper functions. These are named as {{ .Command}}OptionsTo{{ .RequiredStruct }}().

Here {{ .RequiredStruct }} is the name of the struct in Camel Case.

*/
import (
	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/console"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/manage"
    log "github.com/pspiagicw/colorlog"
	runner "github.com/pspiagicw/qemantra/pkg/run"
)

func ParseOptions(version string) {
    initFlags(version)

	switch {
	case createMachineCommand.Used:
		createMachineExecute(createMachineFlags)
	case createImgCommand.Used:
		createImgExecute(createImgFlags)
	case runCommand.Used:
		runExecute(runFlags)
	case listCommand.Used:
		listExecute(listFlags)
	case checkCommand.Used:
		config.PerformCheck()
	case renameCommand.Used:
		renameExecute(renameFlags)
	case editCommand.Used:
		editExecute(editFlags)
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
	log.LogInfo("Creating a new machine!")
	cr := (*manage.Machine)(options)
	manage.CreateMachine(cr)
}

// -- CREATE IMG
func createImgExecute(options *CreateImgFlags) {
	log.LogInfo("Creating a new image!")
	im := (*image.Image)(options)
	_, err := image.CreateImage(im)
	if err != nil {
		log.LogFatal("Error creating image %v", err)
	}
}

// -- RUN
func runExecute(options *RunFlags) {
	log.LogInfo("Finding the given machine!")
	machine := manage.FindMachine(options.name, true)

	if machine == nil {
		log.LogFatal("Machine %s not found", options.name)
	}
	RunOptionsToRunner(options, machine)
	runner.RunMachine(machine)
}

// -- LIST
func listExecute(options *ListFlags) {
	if options.Img {
		image.ListImages(options.Verbose)
	} else {
		manage.ListMachines(options.Verbose)
	}
}

// -- RENAME
func renameExecute(options *RenameFlags) {
	oldName := options.OldName
	newName := options.NewName
	manage.RenameMachine(oldName, newName)
}

// -- EDIT
func editExecute(options *EditFlags) {
	log.LogInfo("Finding the given Machine")

	runner := manage.FindMachine(options.Name, false)

	if runner == nil {
		log.LogFatal("Machine %s not found", options.Name)
	}

	machine := (*manage.Machine)(options)
	manage.EditMachine(machine, runner)
}
