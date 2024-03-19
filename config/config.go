package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/mitchellh/go-homedir"
	"github.com/pspiagicw/goreland"
)

type Conf struct {
	ImageDir   string `toml:"imageDir"`
	MachineDir string `toml:"machineDir"`
}

func ImageDir() string {
	conf := readConf()

	ensureExists(conf.ImageDir)

	return conf.ImageDir
}
func MachineDir() string {
	conf := readConf()

	ensureExists(conf.MachineDir)

	return conf.MachineDir
}
func ensureExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		goreland.LogInfo("Directory '%s' doesn't exist. Creating...", path)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			goreland.LogFatal("Error creating '%s': %v", path, err)
		}
	}
}
func getConfigPath() string {
	path, err := xdg.SearchConfigFile("qemantra/config.toml")

	if err != nil {
		goreland.LogFatal("Error getting config file: %v", err)
	}

	return path
}
func readConf() *Conf {
	path := getConfigPath()

	conf := parseConf(path)

	checkConf(conf)

	sanitizeConf(conf)

	return conf
}
func checkConf(conf *Conf) {
	if conf.ImageDir == "" {
		goreland.LogFatal("ImageDir is not set in config file")
	}
	if conf.MachineDir == "" {
		goreland.LogFatal("MachineDir is not set in config file")
	}
}
func sanitizeConf(conf *Conf) {
	conf.ImageDir = expandPath(conf.ImageDir)
	conf.MachineDir = expandPath(conf.MachineDir)
}
func expandPath(path string) string {
	p, err := homedir.Expand(path)
	if err != nil {
		goreland.LogFatal("Couldn't expand path: %v", err)
	}
	return p
}
func parseConf(path string) *Conf {

	conf := new(Conf)

	_, err := toml.DecodeFile(path, conf)

	if err != nil {
		goreland.LogFatal("Error parsing config: %v", err)

	}

	return conf
}
