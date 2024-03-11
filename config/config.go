package config

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
func readConf() *Conf {
	return &Conf{
		ImageDir:   "/home/pspiagicw/.local/share/qemantra/images",
		MachineDir: "/home/pspiagicw/.local/share/qemantra",
	}
}

// func getConfPath() string {
// 	path, err := xdg.ConfigFile("qemantra/config.toml")
// 	if err != nil {
// 		goreland.LogFatal("Error getting config path.")
// 	}
// }
