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
