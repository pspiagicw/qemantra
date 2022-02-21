package image

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/config"
)

const QEMU_IMAGE_CREATE_COMMMAND = "qemu-img"
const QEMU_IMAGE_CREATE_OPTIONS = "create"

type Image struct {
	Type string
	Name string
	Size string
}

func CreateImage(image *Image) (string, error) {
	imagepath := getImagePath(image.Name)
	command := exec.Command(QEMU_IMAGE_CREATE_COMMMAND, QEMU_IMAGE_CREATE_OPTIONS, imagepath, image.Size)
	var out bytes.Buffer
	fmt.Println(command)
	command.Stderr = &out
	err := command.Run()
	if err != nil {
		return "", err
	}
	return imagepath, nil
}

func getImagePath(name string) string {
	imagesdir := config.GetConfig().ImageDir
	imagepath := appendPath(imagesdir, name)
	return imagepath
}
func appendPath(dir string, name string) string {
	return filepath.Join(dir, name)
}
