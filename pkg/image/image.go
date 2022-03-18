package image

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/config"
	"github.com/pspiagicw/qemantra/pkg/dirs"
	"github.com/pspiagicw/qemantra/pkg/executor"
)

const QEMU_IMAGE_CREATE_COMMMAND = "qemu-img"
const QEMU_IMAGE_CREATE_OPTIONS = "create"
const QEMU_IMAGE_CREATE_FORMAT_OPTION = "-f"

var ConfigProvider = config.GetConfig()
var ExecProvider = executor.GetExecutor()

type Image struct {
	Type string
	Name string
	Size string
}
func FindImage(name string) string {
	files := dirs.ListDirs(ConfigProvider.GetImageDir())
	for _ , i := range files {
		if i.Name() == name {
			return appendPath(ConfigProvider.GetImageDir() , name)
		}
	}
	return ""
}

func CreateImage(image *Image) (string, error) {
	imagepath := getImagePath(image)
	confirmImagePath(imagepath)
	options := getOptions(image)
	err := ExecProvider.Execute(QEMU_IMAGE_CREATE_COMMMAND , options)
	if err != nil {
		return "", err
	}
	return imagepath, nil
}

func getOptions(image *Image) []string{
	options := make([]string , 0)
	options = append(options , QEMU_IMAGE_CREATE_OPTIONS)
	options = append(options , getImageType(image)...)
	options = append(options , getImagePath(image))
	options = append(options , getImageSize(image )...)
	return options
	
}
func getImageSize(image *Image) []string {
	if image.Size == "" {
		return []string{"10G"}
	}
	return []string { image.Size }
}
func getImagePath(image *Image) string {
	imagesdir := ConfigProvider.GetImageDir()
	imagepath := appendPath(imagesdir, image.Name)
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

func getImageType(image *Image) []string {
	if image.Type == "" {
		return []string {"-f" , "raw"}
	}
	return []string{"-f" , image.Type}
}
