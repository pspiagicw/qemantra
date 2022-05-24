package config

import (
	"log"
	"os"
	"path/filepath"
)

const IMAGE_DIR_NAME = "images"
const MACHINE_DIR_NAME = "machines"
const QEMANTRA_DIR = ".qemantra"

type Config interface {
	GetImageDir() string
	GetMachineDir() string
	GetConfigDir() string
}

type UserConfig struct {
	ImageDir   string
	MachineDir string
	ConfigDir  string
}

func (u *UserConfig) GetImageDir() string {
	return u.ImageDir
}
func (u *UserConfig) GetMachineDir() string {
	return u.MachineDir
}
func (u *UserConfig) GetConfigDir() string {
	return u.ConfigDir
}

func getHomeDir() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Error while finding User Home Directory %v", home)
	}

	return home
}
func appendPath(path1 string, path2 string) string {
	path := filepath.Join(path1, path2)
	return path
}
func getConfigDir() string {
	homedir := getHomeDir()
	configdir := appendPath(homedir, QEMANTRA_DIR)
	ensureExists(configdir)
	return configdir
}

func getUserDirs() (string, string) {
	configdir := getConfigDir()

	imagedir := appendPath(configdir, IMAGE_DIR_NAME)
	ensureExists(imagedir)

	machinedir := appendPath(configdir, MACHINE_DIR_NAME)
	ensureExists(machinedir)

	return imagedir, machinedir

}
func GetConfig() Config {
	imagedir, machinedir := getUserDirs()
	configdir := getConfigDir()
	return &UserConfig{
		ImageDir:   imagedir,
		MachineDir: machinedir,
		ConfigDir:  configdir,
	}
}
func ensureExists(dir string) {
	if !dirExists(dir) {
		log.Printf("Creating %s as it does not exists!", dir)
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Error creating %s,%v", dir, err)

		}
	}
}

func dirExists(dir string) bool {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}
	return true

}
