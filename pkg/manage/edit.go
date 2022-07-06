package manage

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/dir"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/run"
)

func EditMachine(m *Machine, r *run.Runner) {
	machine := RunnerToMachine(r, m)
	fmt.Println(machine)
	err := SaveRunnerToDisk(machine)

	if err != nil {
		log.Fatalf("Error writing to file , %v", err)
	}
}

func RunnerToMachine(runner *run.Runner, creator *Machine) *run.Runner {

	if creator.NoDisk {
		runner.DrivePath = ""
	}

	if creator.DiskName != "" && creator.NoDisk != true {
		runner.DrivePath = imageHandler(creator)
	}

	return runner
}
func imageHandler(machine *Machine) string {
	if machine.NoDisk {
		return ""
	}
	for _, image := range dir.ListDir(ConfigProvider.GetImageDir()) {
		if image.Name() == machine.DiskName {
			fmt.Println("Reached here")
			imagepath := filepath.Join(ConfigProvider.GetImageDir(), image.Name())
			fmt.Println(imagepath)
			return imagepath
		}
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
