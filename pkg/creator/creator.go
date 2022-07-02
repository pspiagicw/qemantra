// This package can create and edit machines.
package creator

/*
This file has 2 important functions.
- CreateMachine 
- Edit Machine

Both require a special struct Machine which stores the information required
to create and edit a machine.
*/
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/runner"
)

const SYSTEM_COMMAND string = "qemu-system-x86_64"

var ConfigProvider = config.GetConfig()

// Main struct to create 
type Machine struct {
	Name       string
	NoDisk     bool
	DiskName   string
	DiskFormat string
	DiskSize   string
	MemSize    string
	CpuCores   string
}

func CreateMachine(machine *Machine) {
    if checkIfMachineExists(machine) {
        log.Fatalf("Machine '%s' already exists!" , machine.Name)
    }
	imagepath := createImage(machine)
	runner := constructRunner(imagepath, machine)
	err := encodeJsonToFile(runner)
	if err != nil {
		log.Fatalf("Could not create new machine %v", err)
	}
}
func checkIfMachineExists(machine *Machine) bool {
	runner := runner.FindMachine(machine.Name , false)
	if runner != nil {
		return true
	}
    return false
}
func createImage(machine *Machine) string {
	if machine.NoDisk {
		return ""
	}
	im := &image.Image{
		Type: machine.DiskFormat,
		Name: machine.DiskName,
		Size: machine.DiskSize,
	}
	imagepath, err := image.CreateImage(im)
	if err != nil {
        log.Fatalf("Can't create the disk: %v", err)
	}
	return imagepath
}
func constructRunner(im string, machine *Machine) *runner.Runner {
	runner := &runner.Runner{
		Name:          machine.Name,
		DrivePath:     im,
		SystemCommand: SYSTEM_COMMAND,
		MemSize:       machine.MemSize,
		CpuCores:      machine.CpuCores,
	}
	return runner
}

func encodeJsonToFile(runner *runner.Runner) error {
	contents, err := json.Marshal(runner)
	if err != nil {
		return err
	}
	filepath := getFileName(runner.Name)
	writeFile(contents, filepath)
	return nil
}
func writeFile(contents []byte, filepath string) error {
	err := ioutil.WriteFile(filepath, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
func getFileName(name string) string {
	machineDir := ConfigProvider.GetMachineDir()
	shortName := getShortName(name)
	return filepath.Join(machineDir, shortName)
}
func getShortName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = name + ".json"
	return name
}
