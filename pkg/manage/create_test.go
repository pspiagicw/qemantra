package manage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pspiagicw/qemantra/pkg/image"
	"github.com/pspiagicw/qemantra/pkg/run"

	"github.com/stretchr/testify/assert"
)

type TestExecutor struct {
	errorExecute bool
	Command      []string
}

func (t *TestExecutor) Execute(command string, options []string) error {
	if t.errorExecute {
		return fmt.Errorf("fake error")
	}

	t.Command = []string{command}
	t.Command = append(t.Command, options...)
	return nil
}
func (t *TestExecutor) GetCommand() []string {
	return t.Command
}

type TestConfig struct {
	imagepath   string
	machinepath string
	configpath  string
}

func (t *TestConfig) GetImageDir() string {
	return t.imagepath
}
func (t *TestConfig) GetMachineDir() string {
	return t.machinepath
}
func (t *TestConfig) GetConfigDir() string {
	return t.configpath
}

func setupTestComplex(t *testing.T, files map[string]([]byte)) (string, func(t *testing.T)) {
	dirpath, err := os.MkdirTemp("", "listing")
	if err != nil {
		t.Fatalf("error setting up the test: %v", err)
	}
	for name, contents := range files {
		tmpfilepath := filepath.Join(dirpath, name)
		err := os.WriteFile(tmpfilepath, contents, 0644)
		if err != nil {
			t.Fatalf("error setting up the test: %v", err)
		}
	}

	return dirpath, func(t *testing.T) {
		os.RemoveAll(dirpath)
	}
}

func setupDir(t *testing.T) (string, func(t *testing.T)) {
	dirpath, err := os.MkdirTemp("", "listing")
	if err != nil {
		t.Fatalf("error setting up the test: %v", err)
	}
	return dirpath, func(t *testing.T) {
		os.RemoveAll(dirpath)
	}
}

func setupTestSimple(t *testing.T, files []string) (string, func(t *testing.T)) {
	dirpath, err := os.MkdirTemp("", "listing")
	if err != nil {
		t.Fatalf("error setting up the test: %v", err)
	}
	for _, name := range files {
		tmpfilepath := filepath.Join(dirpath, name)
		err := os.WriteFile(tmpfilepath, []byte(""), 0644)
		if err != nil {
			t.Fatalf("error setting up the test: %v", err)
		}
	}

	return dirpath, func(t *testing.T) {
		os.RemoveAll(dirpath)
	}

}

func TestIfMachineExists(t *testing.T) {
	path, tearDown := setupTestComplex(t, map[string]([]byte){
		"test.json":         []byte("{ \"Name\": \"Test\"}"),
		"test_machine.json": []byte("{ \"Name\": \"Test Machine\"}"),
	})
	defer tearDown(t)

	previousConfigProvider := ConfigProvider

	ConfigProvider = &TestConfig{
		machinepath: path,
		imagepath:   path,
		configpath:  path,
	}
	run.ConfigProvider = &TestConfig{
		machinepath: path,
		imagepath:   path,
		configpath:  path,
	}

	machines := []Machine{
		{
			Runner: run.Runner{
				Name: "Test",
			},
		},
		{
			Runner: run.Runner{
				Name: "Test Machine",
			},
		},
		{
			Runner: run.Runner{
				Name: "Does not exist machine",
			},
		},
	}

	wanted := []bool{
		true,
		true,
		false,
	}
	for i, tt := range machines {
		got := ifMachineExists(&tt)
		assert.Equal(t, got, wanted[i])
	}

	ConfigProvider = previousConfigProvider
	run.ConfigProvider = previousConfigProvider

}
func TestCreateMachine(t *testing.T) {
	// path, tearDown := setupTestComplex(
	// 	t,
	// 	files,
	// )
	path, tearDown := setupDir(t)
	defer tearDown(t)
	previousExecProvider := image.ExecProvider
	image.ExecProvider = &TestExecutor{
		errorExecute: false,
	}
	previousConfigProvider := image.ConfigProvider
	image.ConfigProvider = &TestConfig{
		machinepath: path,
		configpath:  path,
		imagepath:   path,
	}
	ConfigProvider = &TestConfig{
		machinepath: path,
		imagepath:   path,
		configpath:  path,
	}
	t.Run("create a machine", func(t *testing.T) {
		machines := []Machine{
			{
				NoDisk: true,
				Runner: run.Runner{
					Name:     "Example",
					CpuCores: "2",
					MemSize:  "4G",
				},
			},
			{
				NoDisk:     false,
				DiskName:   "test.img",
				DiskSize:   "10G",
				DiskFormat: "qcow2",
				Runner: run.Runner{
					CpuCores:  "2",
					MemSize:   "4G",
					Name:      "Example Machine",
					DrivePath: filepath.Join(path, "test.img"),
				},
			},
		}
		wanted := [][]string{
			{},
			{"qemu-img", "create", "-f", "qcow2", filepath.Join(path, "test.img"), "10G"},
		}
		for i, machine := range machines {
			CreateMachine(&machine)
			shortName := generateShortName(machine.Name)
			path := filepath.Join(path, shortName)
			assert.FileExists(t, path)
			assertJSONFileEQ(t, path, machine.Runner)
			assert.ElementsMatch(t, wanted[i], image.ExecProvider.GetCommand())
		}
	})
	ConfigProvider = previousConfigProvider
	image.ConfigProvider = previousConfigProvider
	image.ExecProvider = previousExecProvider
}
func TestShortName(t *testing.T) {
	tables := []string{
		"Example Machine",
		"Test Machine",
		"Steam Deck",
	}

	wanted := []string{"example_machine.json", "test_machine.json", "steam_deck.json"}

	for i, tt := range tables {
		got := generateShortName(tt)
		want := wanted[i]

		if got != want {
			t.Errorf("got %v , wanted %v", got, want)
		}
	}
}

func TestGetRunnerPath(t *testing.T) {
	tables := []string{
		"Example Machine",
		"Test Machine",
		"Steam Deck",
	}

	wanted := []string{"example_machine.json", "test_machine.json", "steam_deck.json"}

	examplePath := "machinedir"
	// previousExecProvider := ExecProvider
	previousConfigProvider := ConfigProvider
	// ExecProvider = &TestExecutor{}
	ConfigProvider = &TestConfig{
		machinepath: examplePath,
	}

	for i, tt := range tables {
		got := generateRunnerPath(tt)
		want := filepath.Join(examplePath, wanted[i])

		assert.Equal(t, got, want)
	}
	// ExecProvider = previousExecProvider
	ConfigProvider = previousConfigProvider
}

func assertJSONFileEQ(t testing.TB, filepath string, value interface{}) {
	t.Helper()
	contents, err := os.ReadFile(filepath)
	// log.Println(string(contents))
	if err != nil {
		t.Fatalf("Can't read file '%s'", filepath)
	}
	var copy run.Runner
	err = json.Unmarshal(contents, &copy)
	if err != nil {
		t.Fatalf("Error Unmarshaling value: %v", err)
	}
    fmt.Println(value , copy)
	assert.Equal(t, value, copy)
    // assert.ObjectsAreEqualj
}
