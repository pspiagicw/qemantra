package image

import (
	"fmt"
	"github.com/pspiagicw/qemantra/pkg/dir"
	"path/filepath"
	"strings"
)

func getImageList() []string {
	paths := make([]string, 0)
	for _, file := range dir.ListDir(ConfigProvider.GetImageDir()) {
		filepath := getFilePath(file.Name())
		paths = append(paths, filepath)
	}
	return paths

}
func ListImages(verbose bool) {
	images := getImageList()
	for i, image := range images {
		fmt.Printf("%d) Path: %s\n", i, image)
	}
}
func getShortName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = name + ".json"
	return name
}
func getFilePath(name string) string {
	machineDir := ConfigProvider.GetMachineDir()
	path := filepath.Join(machineDir, name)
	return path
}
