package config

import (
	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/pspiagicw/goreland"
)

type Conf struct {
	ImageDir   string
	MachineDir string
}

func ImageDir() string {
	conf := readConf()

	return conf.ImageDir
}
func MachineDir() string {
	conf := readConf()

	return conf.MachineDir
}
func getConfigPath() string {
	path, err := xdg.ConfigFile("qemantra/config.toml")

	if err != nil {
		goreland.LogFatal("Error getting config path: %v", err)
	}

	return path
}
func readConf() *Conf {
	path := getConfigPath()

	conf := parseConf(path)

	return conf
}
func parseConf(path string) *Conf {

	conf := new(Conf)

	_, err := toml.DecodeFile(path, conf)

	if err != nil {
		goreland.LogFatal("Error parsing config: %v", err)

	}

	return conf
}
