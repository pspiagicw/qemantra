package image

import (
	"bytes"
	"log"
	"os"
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
	confirmImagePath(imagepath)
	command := exec.Command(QEMU_IMAGE_CREATE_COMMMAND, QEMU_IMAGE_CREATE_OPTIONS, imagepath, image.Size)
	log.Printf("Executing '%s' on your operating system", command.String())
	var out bytes.Buffer
	command.Stderr = &out
	err := command.Run()
	if err != nil {
		return "", err
	}
	return imagepath, nil
}

func getImagePath(name string) string {
	imagesdir := config.GetConfig().GetImageDir()
	imagepath := appendPath(imagesdir, name)
	return imagepath
}
func appendPath(dir string, name string) string {
	return filepath.Join(dir, name)
}
func confirmImagePath(imagepath string) {
	_, err := os.Stat(imagepath)
	if os.IsNotExist(err) == false {
		log.Fatalf("Disk %s already exists", imagepath)
	}
}
