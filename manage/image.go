package manage

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gosimple/slug"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/config"
	"github.com/pspiagicw/qemantra/vm"
)

const QEMU_IMAGE_CREATE_COMMMAND = "qemu-img"
const QEMU_IMAGE_CREATE_OPTIONS = "create"
const QEMU_IMAGE_CREATE_FORMAT_OPTION = "-f"

func createImage(m *vm.VirtualMachine) string {

	path := getImagePath(m)

	checkImageExists(path)

	m.DiskPath = path

	cmd := getImageCommand(m)

	executeCommand(cmd)

	return path
}
func executeCommand(cmd *exec.Cmd) {
	err := cmd.Run()

	if err != nil {
		goreland.LogFatal("Error running command '%s': %v", cmd.String(), err)
	}

}
func checkImageExists(path string) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrExist) {
		goreland.LogFatal("Disk '%s' exists!")
	}
	if err == nil {
		goreland.LogFatal("Error checking disk '%s': %v", path, err)
	}
}

func getImagePath(m *vm.VirtualMachine) string {
	name := slug.Make(m.DiskName)

	return filepath.Join(config.ImageDir(), name)
}
func getImageCommand(m *vm.VirtualMachine) *exec.Cmd {
	opts := make([]string, 0)
	opts = append(opts, QEMU_IMAGE_CREATE_OPTIONS)
	opts = append(opts, "-f")
	opts = append(opts, m.DiskFormat)
	opts = append(opts, m.DiskPath)
	opts = append(opts, m.DiskSize)

	cmd := exec.Command(QEMU_IMAGE_CREATE_COMMMAND, opts...)

	return cmd
}
