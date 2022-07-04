package creator

import (
    "fmt"
    "testing"
    "os"
    "path/filepath"
    )

type TestExecutor struct {
    errorExecute bool
	Command []string
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

// func TestCreateMachine(t *testing.T) {
//     t.Run("machine already exists" , func(t *testing.T) {
// 		runners := []runner.Runner{
// 			{
// 				Name:          "test",
// 				DrivePath:     "test.img",
// 				SystemCommand: "qemu-system-x86_64",
// 				MemSize:       "",
// 				CpuCores:      "",
// 				Iso:           "",
// 			},
// 			{
// 				Name:          "KVM Example",
// 				DrivePath:     "test.img",
// 				SystemCommand: "qemu-system-x86_64",
// 				MemSize:       "",
// 				CpuCores:      "",
// 				Iso:           "",
// 			},
// 			{
// 				Name:          "Something random",
// 				DrivePath:     "test.img",
// 				SystemCommand: "qemu-system-x86_64",
// 				MemSize:       "",
// 				CpuCores:      "",
// 				Iso:           "",
// 			},
// 		}
//         files := make( map[string]([]byte))
// 		for _, runnner := range runners {
// 			content, err := json.Marshal(runnner)
// 			if err != nil {
// 				t.Fatalf("Failed to marshal runner: %v", err)
// 			}
// 			files[runnner.Name] = content
// 		}
//
// 		path, tearDown := setupTestComplex(
// 			t,
// 			files,
// 		)
// 		defer tearDown(t)
//
// 		image.ExecProvider = &TestExecutor{
// 			errorExecute: false,
// 		}
//         image.ConfigProvider = &TestConfig {
//             machinepath: path,
//         }
// 		ConfigProvider = &TestConfig{
// 			machinepath: path,
//             imagepath: path,
// 		}
//
//     })
// }
func TestShortName(t *testing.T) {
    tables := []string{
        "Example Machine",
        "Test Machine",
        "Steam Deck",
    }

    wanted := []string{"example_machine.json" , "test_machine.json" , "steam_deck.json"}


    for i , tt := range tables {
        got := getShortName(tt)
        want := wanted[i]

        if got != want {
            t.Errorf("got %v , wanted %v" , got , want)
        }
    }
}

func TestGetFileName(t *testing.T) {
    tables := []string{
        "Example Machine",
        "Test Machine",
        "Steam Deck",
    }

    wanted := []string{"example_machine.json" , "test_machine.json" , "steam_deck.json"}

    examplePath := "machinedir"
    // previousExecProvider := ExecProvider
    previousConfigProvider := ConfigProvider
    // ExecProvider = &TestExecutor{}
    ConfigProvider = &TestConfig{
        machinepath: examplePath,
    }

    for i , tt := range tables {
        got := getFileName(tt)
        want := filepath.Join(examplePath, wanted[i])

        if got != want {
            t.Errorf("got %v , wanted %v" , got , want)
        }
    }
    // ExecProvider = previousExecProvider
    ConfigProvider = previousConfigProvider
}

