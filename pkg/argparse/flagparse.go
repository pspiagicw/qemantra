// Package parses the flags passed to `qemantra`
package argparse

/*
This file parses the command line arguments and converts them into a internal struct known as `Options`.
It uses `flaggy` for this parsing.

The OPTIONS struct stores the command itself and it's options.
The actual parsing of arguments happen in the `ParseArguments` function.

The `Options` struct has 2 types of field for each command.Here {{ .Command }} is the name of the command in Camel case
- {{ .Command }}Options : Stores the flags/options for that command

Every subcommand has a global variable named:
- {{ .Command }}Command : Stores the command's instance itself as it is needed later.

Example `run` command has 2 important elements
- RunOptions: In the `Options` struct
- RunCommand: Global Variable

Every command needs 1 functions to be initiated with the name add{{ .Command }}Command(). Thus the `run` command has 1 function
- addRunCommand:
This function is incharge of creating the function , adding arguments to the function and registering with flaggy.

*/
import (
	"github.com/integrii/flaggy"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/manage"
)

// Stores all the flags passed.
type Flags struct {
	runFlags           *RunFlags
	createImgFlags     *CreateImgFlags
	createMachineFlags *CreateMachineFlags
	listFlags          *ListFlags
	renameFlags        *RenameFlags
	editFlags          *EditFlags
}

var runCommand *flaggy.Subcommand
var createImgCommand *flaggy.Subcommand
var createMachineCommand *flaggy.Subcommand
var listCommand *flaggy.Subcommand
var checkCommand *flaggy.Subcommand
var renameCommand *flaggy.Subcommand
var editCommand *flaggy.Subcommand

// Function to parse all the command line arguments.
// Set's qemantra's version and description.
func ParseFlags(version string) *Flags {
	setFlaggyInfo(version)
	global := addSubCommands()
	flaggy.Parse()
	return global
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
func addSubCommands() *Flags {
	global := new(Flags)

	checkCommand = addCheckCommand()

	runCommand, global.runFlags = addRunCommand()
	createImgCommand, global.createImgFlags = addCreateImgCommand()
	createMachineCommand, global.createMachineFlags = addCreateMachineCommand()
	listCommand, global.listFlags = addListCommand()
	renameCommand, global.renameFlags = addRenameCommand()
	editCommand, global.editFlags = addEditCommand()

	return global
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

func addRunCommand() (*flaggy.Subcommand, *RunFlags) {
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
	return run, flags
}

// -- CREATE IMG COMMAND

type CreateImgFlags image.Image

func addCreateImgCommand() (*flaggy.Subcommand, *CreateImgFlags) {
	flags := new(CreateImgFlags)
	create_img := flaggy.NewSubcommand("create-img")
	create_img.String(&flags.Name, "n", "name", "Name of the disk")
	create_img.String(&flags.Type, "f", "format", "Type of the disk")
	create_img.String(&flags.Size, "s", "size", "Size of the disk")
	flaggy.AttachSubcommand(create_img, 1)
	return create_img, flags
}

// -- CREATE MACHINE COMMAND
// type CreateMachineOptions struct {
// 	Name       string
// 	NoDisk     bool
// 	DiskName   string
// 	DiskFormat string
// 	DiskSize   string
// 	MemSize    string
// 	CpuCores   string
// }

type CreateMachineFlags manage.Machine

func addCreateMachineCommand() (*flaggy.Subcommand, *CreateMachineFlags) {
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
	return create_machine, flags
}

// -- LIST COMMAND
type ListFlags struct {
	Img     bool
	Verbose bool
}

func addListCommand() (*flaggy.Subcommand, *ListFlags) {
	flags := new(ListFlags)

	list := flaggy.NewSubcommand("list")
	list.Bool(&flags.Img, "i", "images", "List images")
	list.Bool(&flags.Verbose, "v", "verbose", "All details")

	flaggy.AttachSubcommand(list, 1)
	return list, flags
}

// -- RENAME COMMAND
type RenameFlags struct {
	OldName string
	NewName string
}

func addRenameCommand() (*flaggy.Subcommand, *RenameFlags) {
	flags := new(RenameFlags)

	rename := flaggy.NewSubcommand("rename")
	rename.String(&flags.OldName, "o", "old-name", "Name of the macine currently")
	rename.String(&flags.NewName, "n", "new-name", "The new name to rename the machine to")
	flaggy.AttachSubcommand(rename, 1)
	return rename, flags

}

// -- CHECK COMMAND
func addCheckCommand() *flaggy.Subcommand {
	check := flaggy.NewSubcommand("check")
	flaggy.AttachSubcommand(check, 1)
	return check
}

// -- EDIT COMMAND
type EditFlags manage.Machine

func addEditCommand() (*flaggy.Subcommand, *EditFlags) {
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
	return edit_machine, flags
}
