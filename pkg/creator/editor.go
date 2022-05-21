package creator

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/dirs"
	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/runner"
)

func EditMachine(editor *MachineCreator , currentRunner *runner.Runner) {
	editedRunner := mergeCreatorRunnner(currentRunner , editor)
	fmt.Println(editedRunner)
	err := encodeJsonToFile(editedRunner)

	if err != nil {
		log.Fatalf("Error writing to file , %v", err)
	}
}


func mergeCreatorRunnner(runner *runner.Runner , creator *MachineCreator) *runner.Runner {

	if creator.CpuCores != "" {
		runner.CpuCores = creator.CpuCores
	}

	if creator.MemSize != ""{
		runner.MemSize = creator.MemSize
	}

	if creator.NoDisk {
		runner.DrivePath = ""
	}

	if creator.DiskName != "" && creator.NoDisk != true {
		runner.DrivePath = imageHandler(creator)
	}
	
	return runner
}
func imageHandler(machine *MachineCreator) string {
	if machine.NoDisk {
		return ""
	}
	for _, image := range dirs.ListDirs(ConfigProvider.GetImageDir()) {
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
