package config

import (
	"log"
	"os"
	"path/filepath"
)

const IMAGE_DIR_NAME = "images"
const MACHINE_DIR_NAME = "machines"

type Config struct {
	ImageDir   string
	MachineDir string
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

	configdir := appendPath(homedir, ".lazyqemu")
	return configdir
}

func getUserDirs() (string, string) {
	configdir := getConfigDir()

	imagedir := appendPath(configdir, IMAGE_DIR_NAME)

	machinedir := appendPath(configdir, MACHINE_DIR_NAME)

	return imagedir, machinedir

}
func GetConfig() *Config {
	imagedir, machinedir := getUserDirs()
	return &Config{
		ImageDir:   imagedir,
		MachineDir: machinedir,
	}
}
