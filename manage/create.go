package manage

import (
	"flag"
	"os"
	"path/filepath"

	"encoding/xml"

	"github.com/gosimple/slug"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/config"
	"github.com/pspiagicw/qemantra/help"
	"github.com/pspiagicw/qemantra/vm"
)

func parseCreateArgs(args []string) {
	flag := flag.NewFlagSet("qemantra create", flag.ExitOnError)

	flag.Usage = help.HelpCreate

	flag.Parse(args)
}
func CreateVM(args []string) {
	parseCreateArgs(args)

	m := vm.PromptMachine(nil)

	create(m)
}

func create(m *vm.VirtualMachine) {

	mpath := checkMachineExists(m)

	checkImage(m)

	saveToDisk(m, mpath)
}
func checkImage(m *vm.VirtualMachine) {

	if m.DiskName == "" {
		return
	}

	imagepath := createImage(m)

	m.DiskPath = imagepath
}
func saveToDisk(m *vm.VirtualMachine, path string) {
	contents := serialize(m)

	writeFile(contents, path)
}
func writeFile(contents []byte, path string) {
	err := os.WriteFile(path, contents, 0644)
	if err != nil {
		goreland.LogFatal("Error writing file '%s': %v", path, err)
	}
}
func serialize(m *vm.VirtualMachine) []byte {
	contents, err := xml.Marshal(m)
	if err != nil {
		goreland.LogFatal("Error encoding machine into XML.")
	}

	return contents
}
func getMachinePath(m *vm.VirtualMachine) string {
	name := slug.Make(m.Name)

	machinePath := filepath.Join(config.MachineDir(), name)

	return machinePath

}
func checkMachineExists(m *vm.VirtualMachine) string {
	machinePath := getMachinePath(m)

	_, err := os.Stat(machinePath)
	if err == nil {
		goreland.LogFatal("Machine '%s' already exists", m.Name)
	}

	return machinePath
}
