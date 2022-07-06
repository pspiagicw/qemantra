package run

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pspiagicw/qemantra/pkg/config"
)

var ConfigProvider = config.GetConfig()

func readFile(file string) ([]byte, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return contents, nil
}
func getFileName(file fs.FileInfo) string {
	machineDir := ConfigProvider.GetMachineDir()
	path := filepath.Join(machineDir, file.Name())
	return path
}

func checkName(filepath string, name string) (*Runner, bool) {
	runner, err := decodeFileToRunner(filepath)
	fmt.Println(runner)
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
