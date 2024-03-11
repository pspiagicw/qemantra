package help

import "fmt"

func HelpCreate() {
	fmt.Println("Create virtual machine.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra create")
	fmt.Println()
	fmt.Println("This should start a interactive prompt.")
	fmt.Println("It expects the following information")
	fmt.Println("   - Name")
	fmt.Println("   - CPU Cores")
	fmt.Println("   - RAM Size (Should be in format, 2G, 4G)")
	fmt.Println("   - System Command (architecture of virtual machine i386, x86)")
	fmt.Println("   - Disk Name (Only if you attach disk)")
	fmt.Println("   - Disk Size (Only if you attach disk)")
	fmt.Println("   - Disk Format (Only if you attach disk)")
	fmt.Println()
}
func HelpRun() {
	fmt.Println("Run virtual machine.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra run [flags]")
	fmt.Println()
	flags := `
--iso
--boot
--kvm
--external
--uefi`
	descriptions := `
Path of the ISO to attach
Edit boot order ('menu' or 'iso', default 'iso')
Enable/Disable KVM (Enabled by default)
Path of external disk to attach
Path of OVMF file for a UEFI firmware.`

	fmt.Println("FLAGS")
	printAligned(flags, descriptions)
	fmt.Println("EXAMPLE")
	fmt.Println("   qemantra run -iso <path-of-iso> ")
	fmt.Println("   qemantra run --boot menu --iso <path-of-iso> ")
	fmt.Println()
}
func HelpList() {
	fmt.Println("List virtual machine.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra list")
	fmt.Println()
}
func HelpRename() {
	fmt.Println("Rename virtual machine.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra rename")
	fmt.Println()

}
func HelpEdit() {
	fmt.Println("Edit virtual machine.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   qemantra edit")
	fmt.Println()

}
