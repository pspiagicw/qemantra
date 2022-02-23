package creator

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

type MachineCreator struct {
	Name       string
	NoDisk       bool
	DiskName   string
	DiskFormat string
	DiskSize   string
	MemSize    string
	CpuCores   string
}

func CreateNewMachine(machine *MachineCreator) {
	checkIfMachineExists(machine)
	imagepath := createImage(machine)
	runner := constructRunner(imagepath , machine)
	err := encodeJsonToFile(runner)
	if err != nil {
		log.Fatalf("Could not create new machine %v" , err)
	}
}
func checkIfMachineExists(machine *MachineCreator) {
	runner := runner.FindMachine(machine.Name)
	if runner != nil {
		log.Fatalf("Machine %s already exists!" , machine.Name)
	}
}
func createImage(machine *MachineCreator) string {
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
		log.Fatalf("Can't create the disk %v", err)
	}
	return imagepath
}
func constructRunner(im string, machine *MachineCreator) *runner.Runner {
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
	err := ioutil.WriteFile(filepath, contents, 0644 )
	if err != nil {
		return err
	}
	return nil
}
func getFileName(name string) string {
	machineDir := config.GetConfig().MachineDir
	shortName := getShortName(name)
	return filepath.Join(machineDir, shortName)
}
func getShortName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = name + ".json"
	return name
}
