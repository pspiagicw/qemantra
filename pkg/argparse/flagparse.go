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

/*
	Package parses the flags provided to qemantra.

# Introduction

Command parsing in qemantra is done using a external package flaggy.
The default flag package does not provide extensive support for subcommands.

qemantra supports multiple commands

 1. create-machine
 2. create-img
 3. run
 4. list
 5. check
 6. rename
 7. edit

The most efficient way of managing all these subcommands is to keep a struct for each indivisual command.
This struct stores the values provided through the command line.
*/
package argparse

import (
	"github.com/integrii/flaggy"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/manage"
)

// // Stores all the Flags passed.
// type Flags struct {
// 	runFlags           *RunFlags
// 	createImgFlags     *CreateImgFlags
// 	createMachineFlags *CreateMachineFlags
// 	listFlags          *ListFlags
// 	renameFlags        *RenameFlags
// 	editFlags          *EditFlags
// }

var (
	runFlags             *RunFlags
    createImgFlags *CreateImgFlags
    createMachineFlags *CreateMachineFlags
    listFlags *ListFlags
    renameFlags *RenameFlags
    editFlags *EditFlags

)
var (
	runCommand           *flaggy.Subcommand
	createImgCommand     *flaggy.Subcommand
	createMachineCommand *flaggy.Subcommand
	listCommand          *flaggy.Subcommand
	checkCommand         *flaggy.Subcommand
	renameCommand        *flaggy.Subcommand
	editCommand          *flaggy.Subcommand
)

// Function to parse all the command line arguments.
// Set's qemantra's version and description.
func initFlags(version string) {
	setFlaggyInfo(version)
	addSubCommands()
	flaggy.Parse()
}

// Helper function to set all the information about the program with flaggy.
func setFlaggyInfo(version string) {
	flaggy.SetName("qemantra")
	flaggy.SetDescription("Control QEMU like magic")
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.SetVersion(version)
	flaggy.DefaultParser.AdditionalHelpAppend = "https://github.com/pspiagicw/qemantra"
}

// Adds all the subCommands to the Options struct
func addSubCommands() {

	addCheckCommand()
	addRunCommand()
	addCreateImgCommand()
	addCreateMachineCommand()
	addListCommand()
	addRenameCommand()
	addEditCommand()

}

/*
   ALL COMMANDS START FROM HERE
*/

// --- RUN COMMAND
type RunFlags struct {
	name         string
	iso          string
	diskname     string
	externaldisk string
	boot         string
	uefi         bool
	no_kvm       bool
}

func addRunCommand() {
	flags := new(RunFlags)
	run := flaggy.NewSubcommand("run")

	run.String(&flags.name, "n", "name", "Name of the Machine")
	run.String(&flags.iso, "i", "iso", "Name of the ISO")
	run.String(&flags.diskname, "d", "disk", "Add disk to boot order")
	run.String(&flags.externaldisk, "e", "externaldisk", "Add external disk to boot order")
	run.String(&flags.boot, "b", "boot", "Boot options")
	run.Bool(&flags.uefi, "u", "uefi", "Enable UEFI support")
	run.Bool(&flags.no_kvm, "k", "no-kvm", "Disable KVM Support")

	flaggy.AttachSubcommand(run, 1)
    runCommand = run
    runFlags = flags
}

// -- CREATE IMG COMMAND

type CreateImgFlags image.Image

func addCreateImgCommand() {
	flags := new(CreateImgFlags)
	create_img := flaggy.NewSubcommand("create-img")
	create_img.String(&flags.Name, "n", "name", "Name of the disk")
	create_img.String(&flags.Type, "f", "format", "Type of the disk")
	create_img.String(&flags.Size, "s", "size", "Size of the disk")
	flaggy.AttachSubcommand(create_img, 1)

    createImgCommand = create_img
    createImgFlags = flags
}

type CreateMachineFlags manage.Machine

func addCreateMachineCommand() {
	flags := new(CreateMachineFlags)

	create_machine := flaggy.NewSubcommand("create-machine")
	create_machine.String(&flags.Name, "n", "name", "Name of the machine")
	create_machine.Bool(&flags.NoDisk, "x", "no-disk", "Don't create disk")
	create_machine.String(&flags.DiskName, "i", "disk-name", "Name of the disk")
	create_machine.String(&flags.DiskFormat, "f", "disk-format", "Format of the disk")
	create_machine.String(&flags.DiskSize, "s", "disk-size", "Size of the disk")
	create_machine.String(&flags.MemSize, "m", "mem-size", "Ram to provide")
	create_machine.String(&flags.CpuCores, "c", "cpu-cores", "Cores to provide")
	create_machine.String(&flags.Runner.SystemCommand, "S", "system-command", "System command to use")

	flaggy.AttachSubcommand(create_machine, 1)
    createMachineCommand = create_machine
    createMachineFlags = flags
}

type ListFlags struct {
	Img     bool
	Verbose bool
}

func addListCommand() {
	flags := new(ListFlags)

	list := flaggy.NewSubcommand("list")
	list.Bool(&flags.Img, "i", "images", "List images")
	list.Bool(&flags.Verbose, "v", "verbose", "All details")

	flaggy.AttachSubcommand(list, 1)
    listCommand = list
    listFlags = flags
}

type RenameFlags struct {
	OldName string
	NewName string
}

func addRenameCommand() {
	flags := new(RenameFlags)

	rename := flaggy.NewSubcommand("rename")
	rename.String(&flags.OldName, "o", "old-name", "Name of the macine currently")
	rename.String(&flags.NewName, "n", "new-name", "The new name to rename the machine to")
	flaggy.AttachSubcommand(rename, 1)

    renameCommand = rename
    renameFlags = flags

}

func addCheckCommand() {
	check := flaggy.NewSubcommand("check")
	flaggy.AttachSubcommand(check, 1)
    checkCommand = check
}

type EditFlags manage.Machine

func addEditCommand() {
	flags := new(EditFlags)

	edit_machine := flaggy.NewSubcommand("edit")
	edit_machine.String(&flags.Name, "n", "name", "Name of the machine")
	edit_machine.Bool(&flags.NoDisk, "x", "no-disk", "Don't create disk")
	edit_machine.String(&flags.DiskName, "i", "disk-name", "Name of the disk")
	edit_machine.String(&flags.DiskFormat, "f", "disk-format", "Format of the disk")
	edit_machine.String(&flags.DiskSize, "s", "disk-size", "Size of the disk")
	edit_machine.String(&flags.MemSize, "m", "mem-size", "Ram to provide")
	edit_machine.String(&flags.CpuCores, "c", "cpu-cores", "Cores to provide")

	flaggy.AttachSubcommand(edit_machine, 1)

    editCommand = edit_machine
    editFlags = flags
}
