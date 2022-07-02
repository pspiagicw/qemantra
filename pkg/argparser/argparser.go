// This package parses the flags passed to the binary.
package argparser

/*
This file parses the command line arguments and converts them into a internal struct known as `Options`.
It uses `flaggy` for this parsing.

The OPTIONS struct stores the command itself and it's options.
The actual parsing of arguments happen in the `ParseArguments` function.

The `Options` struct has 2 types of field for each command.Here {{ .Command }} is the name of the command in Pascal case
- {{ .Command }}Options : Stores the flags/options for that command
- {{ .Command }}Command : Stores the command's instance itself as it is needed later.

Example `run` command has 2 fields in the Options struct
- RunOptions 
- RunCommand

Every command needs 1 functions to be initiated with the name add{{ .Command }}Command(). Thus the `run` command has 1 function
- addRunCommand: 
This function is incharge of creating the function , adding arguments to the function and registering with flaggy.

*/
import (
	"github.com/integrii/flaggy"
	"github.com/pspiagicw/qemantra/pkg/creator"
    "github.com/pspiagicw/qemantra/pkg/image"
)

// Master struct to store all the commands and options
type Options struct {
	RunOptions           *RunOptions
	CreateImgOptions     *CreateImgOptions
	CreateMachineOptions *CreateMachineOptions
	ListOptions          *ListOptions
	RenameOptions        *RenameOptions
	EditOptions          *EditOptions

	RunCommand           *flaggy.Subcommand
	CreateImgCommand     *flaggy.Subcommand
	CreateMachineCommand *flaggy.Subcommand
	ListCommand          *flaggy.Subcommand
	CheckCommand         *flaggy.Subcommand
	RenameCommand        *flaggy.Subcommand
	EditCommand          *flaggy.Subcommand
}

// Function to parse all the command line arguments.
// Set's qemantra's version and description.
func ParseArguments(version string) *Options {
	setFlaggyInfo(version)
	globalOptions := addSubCommands()
	flaggy.Parse()
	return globalOptions
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
func addSubCommands() *Options {
	global := new(Options)

	global.RunCommand, global.RunOptions = addRunCommand()
	global.CreateImgCommand, global.CreateImgOptions = addCreateImgCommand()
	global.CreateMachineCommand, global.CreateMachineOptions = addCreateMachineCommand()
	global.ListCommand, global.ListOptions = addListCommand()
	global.CheckCommand = addCheckCommand()
	global.RenameCommand, global.RenameOptions = addRenameCommand()
	global.EditCommand, global.EditOptions = addEditCommand()

	return global
}

/*
    ALL COMMANDS START FROM HERE
*/

// --- RUN COMMAND
type RunOptions struct {
	name         string
	iso          string
	diskname     string
	externaldisk string
	boot         string
	uefi         bool
	no_kvm       bool
}

func addRunCommand() (*flaggy.Subcommand, *RunOptions) {
	options := new(RunOptions)
	run := flaggy.NewSubcommand("run")

	run.String(&options.name, "n", "name", "Name of the Machine")
	run.String(&options.iso, "i", "iso", "Name of the ISO")
	run.String(&options.diskname, "d", "disk", "Add disk to boot order")
	run.String(&options.externaldisk, "e", "externaldisk", "Add external disk to boot order")
	run.String(&options.boot, "b", "boot", "Boot options")
	run.Bool(&options.uefi, "u", "uefi", "Enable UEFI support")
	run.Bool(&options.no_kvm, "k", "no-kvm", "Disable KVM Support")

	flaggy.AttachSubcommand(run, 1)
	return run, options
}

// -- CREATE IMG COMMAND

type CreateImgOptions  image.Image

func addCreateImgCommand() (*flaggy.Subcommand, *CreateImgOptions) {
	options := new(CreateImgOptions)
	create_img := flaggy.NewSubcommand("create-img")
	create_img.String(&options.Name, "n", "name", "Name of the disk")
	create_img.String(&options.Type, "f", "format", "Type of the disk")
	create_img.String(&options.Size, "s", "size", "Size of the disk")
	flaggy.AttachSubcommand(create_img, 1)
	return create_img, options
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

type CreateMachineOptions creator.Machine

func addCreateMachineCommand() (*flaggy.Subcommand, *CreateMachineOptions) {
	options := new(CreateMachineOptions)

	create_machine := flaggy.NewSubcommand("create-machine")
	create_machine.String(&options.Name, "n", "name", "Name of the machine")
	create_machine.Bool(&options.NoDisk, "x", "no-disk", "Don't create disk")
	create_machine.String(&options.DiskName, "i", "disk-name", "Name of the disk")
	create_machine.String(&options.DiskFormat, "f", "disk-format", "Format of the disk")
	create_machine.String(&options.DiskSize, "s", "disk-size", "Size of the disk")
	create_machine.String(&options.MemSize, "m", "mem-size", "Ram to provide")
	create_machine.String(&options.CpuCores, "c", "cpu-cores", "Cores to provide")

	flaggy.AttachSubcommand(create_machine, 1)
	return create_machine, options
}

// -- LIST COMMAND
type ListOptions struct {
	Img     bool
	Verbose bool
}

func addListCommand() (*flaggy.Subcommand, *ListOptions) {
	options := new(ListOptions)

	list := flaggy.NewSubcommand("list")
	list.Bool(&options.Img, "i", "images", "List images")
	list.Bool(&options.Verbose, "v", "verbose", "All details")

	flaggy.AttachSubcommand(list, 1)
	return list, options
}

// -- RENAME COMMAND
type RenameOptions struct {
	OldName string
	NewName string
}

func addRenameCommand() (*flaggy.Subcommand, *RenameOptions) {
	options := new(RenameOptions)

	rename := flaggy.NewSubcommand("rename")
	rename.String(&options.OldName, "o", "old-name", "Name of the macine currently")
	rename.String(&options.NewName, "n", "new-name", "The new name to rename the machine to")
	flaggy.AttachSubcommand(rename, 1)
	return rename, options

}


// -- CHECK COMMAND
func addCheckCommand() *flaggy.Subcommand {
	check := flaggy.NewSubcommand("check")
	flaggy.AttachSubcommand(check, 1)
	return check
}



// -- EDIT COMMAND
type EditOptions creator.Machine

func addEditCommand() (*flaggy.Subcommand, *EditOptions) {
	options := new(EditOptions)

	edit_machine := flaggy.NewSubcommand("edit")
	edit_machine.String(&options.Name, "n", "name", "Name of the machine")
	edit_machine.Bool(&options.NoDisk, "x", "no-disk", "Don't create disk")
	edit_machine.String(&options.DiskName, "i", "disk-name", "Name of the disk")
	edit_machine.String(&options.DiskFormat, "f", "disk-format", "Format of the disk")
	edit_machine.String(&options.DiskSize, "s", "disk-size", "Size of the disk")
	edit_machine.String(&options.MemSize, "m", "mem-size", "Ram to provide")
	edit_machine.String(&options.CpuCores, "c", "cpu-cores", "Cores to provide")

	flaggy.AttachSubcommand(edit_machine, 1)
	return edit_machine, options
}
