package argparser

import (
	"fmt"

	"github.com/integrii/flaggy"
)

func ParseArguments() {
	flaggy.SetName("qemantra")
	flaggy.SetDescription("Control QEMU like magic")
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.SetVersion("0.01")
	flaggy.DefaultParser.AdditionalHelpAppend = "https://github.com/pspiagicw/qemantra"
	run_options, create_options := addSubCommands()
	flaggy.Parse()
	fmt.Println(run_options)
	fmt.Println(create_options)

}
func addSubCommands() (*runCommandOptions, *createImgOptions) {
	runSubcommand, run_options := addRunCommand()
	createImgSubCommand, create_img_options := addCreateImgCommand()
	flaggy.AttachSubcommand(runSubcommand, 1)
	flaggy.AttachSubcommand(createImgSubCommand, 1)
	return run_options, create_img_options
}

type runCommandOptions struct {
	name         string
	iso          string
	diskname     string
	externaldisk string
}

func newRunCommandOptions() *runCommandOptions {
	return &runCommandOptions{
		name:         "",
		diskname:     "",
		iso:          "",
		externaldisk: "",
	}
}
func addRunCommand() (*flaggy.Subcommand, *runCommandOptions) {
	options := newRunCommandOptions()
	run := flaggy.NewSubcommand("run")
	run.String(&options.name, "n", "name", "Name of the Machine")
	run.String(&options.iso, "i", "iso", "Name of the ISO")
	run.String(&options.diskname, "d", "disk", "Add disk to boot order")
	run.String(&options.externaldisk, "e", "externaldisk", "Add external disk to boot order")
	return run, options
}

type createImgOptions struct {
	name   string
	format string
}

func newCreateImgOptions() *createImgOptions {
	return &createImgOptions{
		name:   "",
		format: "",
	}
}
func addCreateImgCommand() (*flaggy.Subcommand, *createImgOptions) {
	options := newCreateImgOptions()
	create_img := flaggy.NewSubcommand("create-img")
	create_img.String(&options.name, "n", "name", "Name of the disk")
	create_img.String(&options.format, "f", "format", "Format of the disk")
	return create_img, options
}

type listImgOptions struct {
	name bool
}
