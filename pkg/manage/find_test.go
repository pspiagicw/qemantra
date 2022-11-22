package manage

import (
	"encoding/json"
	// "os"
	// "path/filepath"
	"github.com/pspiagicw/qemantra/pkg/run"
	"testing"
)

//	func setupTestComplex(t *testing.T, files map[string]([]byte)) (string, func(t *testing.T)) {
//		dirpath, err := os.MkdirTemp("", "listing")
//		if err != nil {
//			t.Fatalf("error setting up the test: %v", err)
//		}
//		for name, contents := range files {
//			tmpfilepath := filepath.Join(dirpath, name)
//			err := os.WriteFile(tmpfilepath, contents, 0644)
//			if err != nil {
//				t.Fatalf("error setting up the test: %v", err)
//			}
//		}
//
//		return dirpath, func(t *testing.T) {
//			os.RemoveAll(dirpath)
//		}
//	}
func TestFindMachine(t *testing.T) {
	t.Run("machine exists", func(t *testing.T) {
		runners := []run.Runner{
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "KVM Example",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "Something random",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
		}
		files := make(map[string]([]byte))
		for _, runnner := range runners {
			content, err := json.Marshal(runnner)
			if err != nil {
				t.Fatalf("Failed to marshal runner: %v", err)
			}
			files[runnner.Name] = content
		}

		path, tearDown := setupTestComplex(
			t,
			files,
		)
		defer tearDown(t)

		// previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider

		// ExecProvider = &TestExecutor{
		// 	errorExecute: false,
		// }
		ConfigProvider = &TestConfig{
			machinepath: path,
		}
		for _, want := range runners {
			got := FindMachine(want.Name, false)
			if *got != want {
				t.Errorf("got %v , wanted %v", got, want)
			}

		}
		// ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})

	t.Run("machine does not exist", func(t *testing.T) {
		runners := []run.Runner{
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "KVM Example",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "Something random",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
		}
		files := make(map[string]([]byte))
		for _, runnner := range runners {
			content, err := json.Marshal(runnner)
			if err != nil {
				t.Fatalf("Failed to marshal runner: %v", err)
			}
			files[runnner.Name+"123"] = content
		}

		path, tearDown := setupTestComplex(
			t,
			files,
		)
		defer tearDown(t)

		// previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider

		// ExecProvider = &TestExecutor{
		// 	errorExecute: false,
		// }
		ConfigProvider = &TestConfig{
			machinepath: path,
		}
		for range runners {
			got := FindMachine("nobody", false)
			if got != nil {
				t.Errorf("got %v , wanted %v", got, nil)
			}

		}
		// ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
}
