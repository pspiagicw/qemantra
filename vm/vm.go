package vm

import (
	prompt "github.com/pspiagicw/qemantra/prompts"
	"github.com/pspiagicw/qemantra/validators"
)

var ARCHS []string = []string{
	"qemu-system-x86_64",
	"qemu-system-i386",
}

var DTYPES []string = []string{
	"raw",
	"vdi",
	"qcow2",
}

type VirtualMachine struct {
	Name string

	Architecture string

	DiskName   string
	DiskFormat string
	DiskSize   string
	DiskPath   string

	ExternalDisk string
	ISO          string
	KVM          bool
	UEFI         string
	Boot         string

	CpuCores string
	MemSize  string
}

func PromptMachine(template *VirtualMachine) *VirtualMachine {

	if template == nil {
		template = new(VirtualMachine)
	}

	vm := new(VirtualMachine)

	vm.Name = prompt.QuestionPrompt("Name", validators.NameValidator, "")
	vm.CpuCores = prompt.QuestionPrompt("CPU Cores", validators.CoreValidator, template.CpuCores)
	vm.MemSize = prompt.QuestionPrompt("RAM Size", validators.MemoryValidator, template.MemSize)

	vm.Architecture = prompt.SelectionPrompt("Architecture", ARCHS)

	if prompt.ConfirmPrompt("Do you want to attach a disk ?") {
		vm.DiskName = prompt.QuestionPrompt("Disk Name", validators.NameValidator, template.DiskName)
		vm.DiskSize = prompt.QuestionPrompt("Disk Size", validators.MemoryValidator, template.DiskSize)
		vm.DiskFormat = prompt.SelectionPrompt("Disk Format", DTYPES)
	}

	return vm

}
