package argparser

import (
	"fmt"
	"github.com/integrii/flaggy"
)

// Struct to store all the parsed options and subcommand
type Options struct {
	RunOptions           *RunCommandOptions
	CreateImgOptions     *CreateImgOptions
	CreateMachineOptions *CreateMachineOptions
	ListOptions          *ListOptions

	RunOptionCommand     *flaggy.Subcommand
	CreateImgCommand     *flaggy.Subcommand
	CreateMachineCommand *flaggy.Subcommand
	ListCommand          *flaggy.Subcommand
}

// Function to parse all the command line arguments.
// Set's qemantra's version and description.
func ParseArguments(version string) *Options {
	flaggy.SetName("qemantra")
	flaggy.SetDescription("Control QEMU like magic")
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.SetVersion(version)
	flaggy.DefaultParser.AdditionalHelpAppend = "https://github.com/pspiagicw/qemantra"
	globalOptions := addSubCommands()
	flaggy.Parse()
	return globalOptions
}

// Helper function to add subcommands and options
func addSubCommands() *Options {
	runSubcommand, run_options := addRunCommand()
	createImgSubCommand, create_img_options := addCreateImgCommand()
	createMachineSubCommand, create_machine_options := addCreateMachineCommand()
	listSubCommand, list_options := addListCommand()
	flaggy.AttachSubcommand(runSubcommand, 1)
	flaggy.AttachSubcommand(createImgSubCommand, 1)
	flaggy.AttachSubcommand(createMachineSubCommand, 1)
	flaggy.AttachSubcommand(listSubCommand, 1)
	global := &Options{
		RunOptions:           run_options,
		CreateImgOptions:     create_img_options,
		CreateMachineOptions: create_machine_options,
		RunOptionCommand:     runSubcommand,
		CreateImgCommand:     createImgSubCommand,
		CreateMachineCommand: createMachineSubCommand,
		ListOptions:          list_options,
		ListCommand:          listSubCommand,
	}
	return global
}

// Struct to store parsed information if `run` subcommand is used.
type RunCommandOptions struct {
	name         string
	iso          string
	diskname     string
	externaldisk string
	boot         string
}

// String function for RunCommandOptions
func (r *RunCommandOptions) String() string {
	return fmt.Sprintf("Name: %s , iso: %s , diskname: %s , externaldisk: %s", r.name, r.iso, r.diskname, r.externaldisk)
}

// Helper function to create a instance of RunCommandOptions
func newRunCommandOptions() *RunCommandOptions {
	return &RunCommandOptions{
		name:         "",
		diskname:     "",
		iso:          "",
		externaldisk: "",
		boot:         "",
	}
}

// Function to add logic for parsing the `run` command
func addRunCommand() (*flaggy.Subcommand, *RunCommandOptions) {
	options := newRunCommandOptions()
	run := flaggy.NewSubcommand("run")
	run.String(&options.name, "n", "name", "Name of the Machine")
	run.String(&options.iso, "i", "iso", "Name of the ISO")
	run.String(&options.diskname, "d", "disk", "Add disk to boot order")
	run.String(&options.externaldisk, "e", "externaldisk", "Add external disk to boot order")
	run.String(&options.boot, "b", "boot", "Boot options")
	return run, options
}

// Struct to store parsed information if `create-img` subcommand is used.
type CreateImgOptions struct {
	Name   string
	Format string
	Size   string
}

// String fucntion to create a instance of CreateImgOptions
func (c *CreateImgOptions) String() string {
	return fmt.Sprintf("Name: %s , format: %s , size: %s", c.Name, c.Format, c.Size)
}

// Helper function to create instance of CreateImgOptions
func newCreateImgOptions() *CreateImgOptions {
	return &CreateImgOptions{
		Name:   "",
		Format: "",
		Size:   "",
	}
}

// Function to add logic for `create-img` command
func addCreateImgCommand() (*flaggy.Subcommand, *CreateImgOptions) {
	options := newCreateImgOptions()
	create_img := flaggy.NewSubcommand("create-img")
	create_img.String(&options.Name, "n", "name", "Name of the disk")
	create_img.String(&options.Format, "f", "format", "Format of the disk")
	create_img.String(&options.Size, "s", "size", "Size of the disk")
	return create_img, options
}

type CreateMachineOptions struct {
	Name       string
	NoDisk     bool
	DiskName   string
	DiskFormat string
	DiskSize   string
	MemSize    string
	CpuCores   string
}

func (c *CreateMachineOptions) String() string {
	if c.NoDisk {
		return fmt.Sprintf("Name: %s , Disk: %v , MemSize: %s , CpuCores: %s", c.Name, c.NoDisk, c.MemSize, c.CpuCores)
	} else {
		return fmt.Sprintf("Name: %s , Diskname: %s , Diskformat: %s , Disksize: %s , MemSize: %s , CpuCores: %s", c.Name, c.DiskName, c.DiskFormat, c.DiskSize, c.MemSize, c.CpuCores)
	}
}

func newCreateMachineOptions() *CreateMachineOptions {
	return &CreateMachineOptions{
		Name:       "",
		NoDisk:     false,
		DiskName:   "",
		DiskFormat: "",
		DiskSize:   "",
		MemSize:    "",
		CpuCores:   "",
	}
}
func addCreateMachineCommand() (*flaggy.Subcommand, *CreateMachineOptions) {
	options := newCreateMachineOptions()

	create_machine := flaggy.NewSubcommand("create-machine")
	create_machine.String(&options.Name, "n", "name", "Name of the machine")
	create_machine.Bool(&options.NoDisk, "x", "no-disk", "Don't create disk")
	create_machine.String(&options.DiskName, "i", "disk-name", "Name of the disk")
	create_machine.String(&options.DiskFormat, "f", "disk-format", "Format of the disk")
	create_machine.String(&options.DiskSize, "s", "disk-size", "Size of the disk")
	create_machine.String(&options.MemSize, "m", "mem-size", "Ram to provide")
	create_machine.String(&options.CpuCores, "c", "cpu-cores", "Cores to provide")
	return create_machine, options
}

type ListOptions struct {
	Img bool
}

func newListOptions() *ListOptions {
	return &ListOptions{
		Img: false,
	}
}
func addListCommand() (*flaggy.Subcommand, *ListOptions) {
	options := newListOptions()

	list := flaggy.NewSubcommand("list")
	list.Bool(&options.Img, "i", "images", "List images")
	return list, options
}
