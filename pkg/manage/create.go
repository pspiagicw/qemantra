// This package can create and edit machines.
package manage

/*
- CreateMachine
- Edit Machine

Both require a special struct Machine which stores the information required
to create and edit a machine.
*/
import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/run"
    log "github.com/pspiagicw/colorlog"
)

const SYSTEM_COMMAND string = "qemu-system-x86_64"

var ConfigProvider = config.GetConfig()

// Main struct to create
type Machine struct {
	NoDisk     bool
	DiskName   string
	DiskFormat string
	DiskSize   string
	run.Runner
}

func CreateMachine(machine *Machine) {
	if ifMachineExists(machine) {
		log.LogFatal("Machine '%s' already exists!", machine.Name)
	}
	imagepath := createImage(machine)
	runner := constructRunner(imagepath, machine)
	err := SaveRunnerToDisk(runner)
	if err != nil {
		log.LogFatal("Could not create new machine %v", err)
	}
}
func ifMachineExists(machine *Machine) bool {
	runner := FindMachine(machine.Name, false)
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
	log.LogInfo("Imagepath: %s", imagepath)
	if err != nil {
		log.LogFatal("Can't create the disk: %v", err)
	}
	return imagepath
}

func constructRunner(imagepath string, machine *Machine) *run.Runner {
	machine.Runner.DrivePath = imagepath

	return &machine.Runner
}

func SaveRunnerToDisk(runner *run.Runner) error {
	contents, err := json.Marshal(runner)
	if err != nil {
		return err
	}
	filepath := generateRunnerPath(runner.Name)
	return writeFile(contents, filepath)
}
func writeFile(contents []byte, filepath string) error {
	err := ioutil.WriteFile(filepath, contents, 0644)
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
	name = strings.ReplaceAll(name, " ", "_")
	name = name + ".json"
	return name
}
