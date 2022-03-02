package runner

import (
	"fmt"
	"log"

	"github.com/pspiagicw/qemantra/pkg/dirs"
)

func FindMachine(name string) *Runner {
	for _, file := range dirs.ListDirs(ConfigProvider.GetMachineDir()) {
		filepath := getFileName(file)
		runner, ok := checkName(filepath, name)
		if ok {
			return runner
		}
	}
	return nil
}
func ListMachines(image bool) {
	if image {
		images := getImageList()
		for i, image := range images {
			fmt.Printf("%d) Path: %s\n", i, image)
		}
		return
	}
	machines := getMachineList()
	for i, runner := range machines {
		fmt.Printf("%d) Name: %s\n", i, runner.Name)
	}
}
func getMachineList() []Runner {
	runners := make([]Runner, 0)
	for _, file := range dirs.ListDirs(ConfigProvider.GetMachineDir()) {
		filepath := getFileName(file)
		runner, err := decodeFileToRunner(filepath)
		if err != nil {
			log.Fatalf("Can't parse %s , %v", filepath, err)
		}
		runners = append(runners, *runner)

	}
	return runners
}

func getImageList() []string {
	paths := make([]string, 0)
	for _, file := range dirs.ListDirs(ConfigProvider.GetImageDir()) {
		filepath := getFileName(file)
		paths = append(paths, filepath)
	}
	return paths

}
