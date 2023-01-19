// This package can create and edit machines.
package manage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	log "github.com/pspiagicw/colorlog"
	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/machine"
)

const SYSTEM_COMMAND string = "qemu-system-x86_64"

var ConfigProvider = config.GetConfig()

func CreateMachine(machine *machine.Machine) {
	if ifMachineExists(machine) {
		log.LogFatal("Machine '%s' already exists!", machine.Name)
	}
	machine.DrivePath = createImage(machine)
	err := saveToDisk(machine)
	if err != nil {
		log.LogFatal("Could not create new machine %v", err)
	}
}
func ifMachineExists(machine *machine.Machine) bool {
	runner := FindMachine(machine.Name)
	if runner != nil {
		return true
	}
	return false
}

func createImage(machine *machine.Machine) string {
	if machine.NoDisk {
		return ""
	}
	im := &image.Image{
		Type: machine.DiskFormat,
		Name: machine.DiskName,
		Size: machine.DiskSize,
	}
	imagepath, err := image.CreateImage(im)
	log.LogInfo("Imagepath: %s", imagepath)
	if err != nil {
		log.LogFatal("Can't create the disk: %v", err)
	}
	return imagepath
}

func saveToDisk(machine *machine.Machine) error {
	filepath := generateRunnerPath(machine.Name)
	contents, err := json.Marshal(machine)
	if err != nil {
		return err
	}
	return writeFile(contents, filepath)
}
func writeFile(contents []byte, filepath string) error {
	err := os.WriteFile(filepath, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
func getRunnerPath(name string) string {
	machineDir := ConfigProvider.GetMachineDir()
	path := filepath.Join(machineDir, name)
	return path
}
func generateRunnerPath(name string) string {
	machineDir := ConfigProvider.GetMachineDir()
	shortName := generateShortName(name)
	return filepath.Join(machineDir, shortName)
}

func generateShortName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = name + ".json"
	return name
}
