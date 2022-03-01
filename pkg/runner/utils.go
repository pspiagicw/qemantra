package runner

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/config"
)

func readFile(file string) ([]byte, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return contents, nil
}
func getFileName(file fs.FileInfo) string {
	machineDir := config.GetConfig().MachineDir
	path := filepath.Join(machineDir, file.Name())
	return path
}

func checkName(filepath string, name string) (*Runner, bool) {
	runner, err := decodeFileToRunner(filepath)
	if err != nil {
		log.Fatalf("Can't decode file %s , %v", filepath, err)
	}

	if runner.Name == name {
		return runner, true
	}
	return nil, false
}
func decodeByteToRunner(contents []byte) (*Runner, error) {
	var runner Runner
	err := json.Unmarshal(contents, &runner)
	if err != nil {
		return nil, err
	}
	return &runner, nil
}
func decodeFileToRunner(filepath string) (*Runner, error) {
	contents, err := readFile(filepath)
	if err != nil {
		return nil, err
	}
	runner, err := decodeByteToRunner(contents)
	if err != nil {
		return nil, err
	}
	return runner, nil
}
