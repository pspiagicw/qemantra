package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
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
	fmt.Println("Manage QEMU/KVM virtual machines.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra [command] [args]")
	fmt.Println()
}
func HelpUsage(version string) {
	PrintVersion(version)
	printHeader()
	printCommands()
	printFooter()
}
func printFooter() {
	fmt.Println("MORE HELP")
	fmt.Println("  Use 'qemantra help [command]' for more info about a command.")
}
func printCommands() {
	fmt.Println("COMMANDS")
	commands := `
check:
create:
list:
run:
rename:
edit:
version:
help:`
	messages := `
Check for dependencies
Create virtual machines
List virtual machines
Run virtual machines
Rename virtual machines
Edit virtual machines
Show version info
Show help info`
	printAligned(commands, messages)
	fmt.Println()
}
func printAligned(left, right string) {
	leftCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(left).MarginLeft(2).String()
	rightCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(right).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, leftCol, rightCol))

	fmt.Println()
}
func PrintVersion(version string) {
	fmt.Printf("qemantra version: %s\n", version)
}
