package help

import (
	"github.com/pspiagicw/pelp"
)

func HelpCreate() {
	pelp.Print("Create virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra create"})
	pelp.HeaderWithDescription("description", []string{
		"This should start a interactive prompt.",
		"It expects the following information",
		"   - Name",
		"   - CPU Cores",
		"   - RAM Size (Should be in format, 2G, 4G)",
		"   - System Command (architecture of virtual machine i386, x86)",
		"   - Disk Name (Only if you attach disk)",
		"   - Disk Size (Only if you attach disk)",
		"   - Disk Format (Only if you attach disk)",
	})
}
func HelpRun() {
	pelp.Print("Run virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra run [flags]"})
	flags := []string{"iso", "boot", "kvm", "external", "uefi"}
	descriptions := []string{
		"Path of the ISO to attach",
		"Edit boot order ('menu' or 'iso', default 'iso')",
		"Enable/Disable KVM (Enabled by default)",
		"Path of external disk to attach",
		"Path of OVMF file for a UEFI firmware.",
	}

	pelp.Flags("flags", flags, descriptions)
	pelp.Examples("example", []string{
		"qemantra run -iso <path-of-iso> ", "qemantra run --boot menu --iso <path-of-iso> "})
}
func HelpList() {
	pelp.Print("List virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra list"})
}
func HelpRename() {
	pelp.Print("Rename virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra rename"})

}
func HelpEdit() {
	pelp.Print("Edit virtual machine.")
	pelp.HeaderWithDescription("usage", []string{"qemantra edit"})

}
