package image

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pspiagicw/lazyqemu/pkg/config"
)

const QEMU_IMAGE_CREATE_COMMMAND = "qemu-img"
const QEMU_IMAGE_CREATE_OPTIONS = "create -f raw"

type Image struct {
	Type string
	Name string
	Size string
}

func CreateImage(image *Image) string {
	imagepath := getImagePath(image.Name)
	command := exec.Command(QEMU_IMAGE_CREATE_COMMMAND, imagepath, QEMU_IMAGE_CREATE_OPTIONS, image.Size)
	fmt.Println(command)
	return imagepath
}

func getImagePath(name string) string {
	imagesdir := config.GetConfig().ImageDir
	imagepath := appendPath(imagesdir, name)
	return imagepath

}
func appendPath(dir string, name string) string {
	return filepath.Join(dir, name)
}
