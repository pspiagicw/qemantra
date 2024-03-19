package help

import (
	"github.com/pspiagicw/pelp"
)

func HandleHelp(args []string, version string) {
	if len(args) == 0 {
		HelpUsage(version)
	} else {
		cmd := args[0]

		handler := map[string]func(){
			"create": HelpCreate,
			"run":    HelpRun,
			"list":   HelpList,
			"rename": HelpRename,
			"edit":   HelpEdit,
			"delete": HelpRemove,
		}

		handlerCmd, ok := handler[cmd]

		if !ok {
			HelpUsage(version)
		} else {
			handlerCmd()
		}

	}
}
func printHeader() {
	pelp.Print("Manage QEMU/KVM virtual machines.")
	pelp.HeaderWithDescription("usage", []string{"qemantra [command] [args]"})
}
func HelpUsage(version string) {
	PrintVersion(version)
	printHeader()
	printCommands()
	printFooter()
}
func printFooter() {
	pelp.HeaderWithDescription("more help", []string{"Use 'qemantra help [command]' for more info about a command."})
}
func printCommands() {
	commands := []string{"check:", "create:", "list:", "run:", "rename:", "edit:", "delete:", "version:", "help:"}
	messages := []string{"Check for dependencies", "Create virtual machines", "List virtual machines", "Run virtual machines", "Rename virtual machines", "Edit virtual machines", "Delete a virtual machine", "Show version info", "Show help info"}
	pelp.Aligned("commands", commands, messages)
}
func PrintVersion(version string) {
	pelp.Version("qemantra", version)
}
func HelpRemove() {
	pelp.Print("Delete virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra delete"})
}
