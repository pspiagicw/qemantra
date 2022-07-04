package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
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

func TestRunMachine(t *testing.T) {
	t.Run("executor works", func(t *testing.T) {
		table := []Runner{
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-aarch64",
				MemSize:       "4G",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-aarch64",
				MemSize:       "4G",
				CpuCores:      "2",
				Iso:           "test.iso",
			},
		}
		wanted := [][]string{
			{"qemu-system-x86_64", "-enable-kvm", "-hda", "test.img", "-cpu", "host"},
			{"qemu-aarch64", "-m", "4G", "-enable-kvm", "-hda", "test.img", "-cpu", "host"},
			{"qemu-aarch64", "-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-cpu", "host", "-smp", "2"},
		}

		path, tearDown := setupTestSimple(t, []string{})
		defer tearDown(t)

		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		// Give error on purpose
		ExecProvider = &TestExecutor{
			errorExecute: false,
		}
		ConfigProvider = &TestConfig{
			imagepath: path,
		}

		for i, tt := range table {
			RunMachine(&tt)
			want := wanted[i]
			got := ExecProvider.GetCommand()
			assertStringArray(t, got, want)
		}

		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
	t.Run("executor does not work", func(t *testing.T) {
		table := []Runner{
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-system-x86_64",
				MemSize:       "",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-aarch64",
				MemSize:       "4G",
				CpuCores:      "",
				Iso:           "",
			},
			{
				Name:          "test",
				DrivePath:     "test.img",
				SystemCommand: "qemu-aarch64",
				MemSize:       "4G",
				CpuCores:      "2",
				Iso:           "test.iso",
			},
		}
		wanted := [][]string{
			{},
			{},
			{},
		}

		path, tearDown := setupTestSimple(t, []string{})
		defer tearDown(t)

		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		// Give error on purpose
		ExecProvider = &TestExecutor{
			errorExecute: true,
		}
		ConfigProvider = &TestConfig{
			imagepath: path,
		}

		for i, tt := range table {
			RunMachine(&tt)
			want := wanted[i]
			got := ExecProvider.GetCommand()
			assertStringArray(t, got, want)
		}

		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
	// t.Run("external disk exists", func(t *testing.T) {
	// 	table := []Runner{
	// 		{
	// 			Name:          "test",
	// 			DrivePath:     "test.img",
	// 			SystemCommand: "qemu-system-x86_64",
	// 			MemSize:       "",
	// 			CpuCores:      "",
	// 			Iso:           "",
	// 			ExternalDisk:  "hello",
	// 		},
	// 	}
	// 	wanted := [][]string{
	// 		{"qemu-system-x86_64", "-enable-kvm", "-hda", "test.img", "-cpu", "host"},
	// 	}
	//
	// 	path, tearDown := setupTest(t, []string{"hello"})
	// 	defer tearDown(t)
	//
	// 	previousExecProvider := ExecProvider
	// 	previousConfigProvider := ConfigProvider
	//
	// 	ExecProvider = &TestExecutor{
	// 		errorExecute: false,
	// 	}
	// 	ConfigProvider = &TestConfig{
	// 		imagepath: path,
	// 	}
	//
	// 	for i, tt := range table {
	// 		RunMachine(&tt)
	// 		want := wanted[i]
	// 		got := ExecProvider.GetCommand()
	// 		assertStringArray(t, got, want)
	// 	}
	//
	// 	ExecProvider = previousExecProvider
	// 	ConfigProvider = previousConfigProvider
	//
	// })
}

func TestConstructOptions(t *testing.T) {
	table := []Runner{
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-system-x86_64",
			MemSize:       "",
			CpuCores:      "",
			Iso:           "",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "",
			Iso:           "",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "2",
			Iso:           "test.iso",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "",
			CpuCores:      "2",
			Iso:           "test.iso",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "",
			Iso:           "test.iso",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "",
			Iso:           "test.iso",
			ExternalDisk:  "externaltest.img",
			Boot:          "menu",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "2",
			Iso:           "test.iso",
			ExternalDisk:  "externaltest.img",
			Boot:          "iso",
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "2",
			Iso:           "test.iso",
			ExternalDisk:  "externaltest.img",
			Boot:          "iso",
			NO_KVM:        true,
		},
		{
			Name:          "test",
			DrivePath:     "test.img",
			SystemCommand: "qemu-aarch64",
			MemSize:       "4G",
			CpuCores:      "2",
			Iso:           "test.iso",
			ExternalDisk:  "externaltest.img",
			Boot:          "iso",
			UEFI:          true,
		},
		{
			Name:          "test",
			DrivePath:     "",
			SystemCommand: "qemu-system-x86_64",
			MemSize:       "",
			CpuCores:      "",
			Iso:           "",
		},
	}
	wanted := [][]string{
		{"-enable-kvm", "-hda", "test.img", "-cpu", "host"},
		{"-m", "4G", "-enable-kvm", "-hda", "test.img", "-cpu", "host"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-cpu", "host", "-smp", "2"},
		{"-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-cpu", "host", "-smp", "2"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-cpu", "host"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "menu=on", "-cpu", "host", "-hdb", "externaltest.img"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "d", "-cpu", "host", "-smp", "2", "-hdb", "externaltest.img"},
		{"-m", "4G", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "d", "-cpu", "host", "-smp", "2", "-hdb", "externaltest.img"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "d", "-cpu", "host", "-smp", "2", "-bios", OVMF_PATH, "-hdb", "externaltest.img"},
		{"-enable-kvm", "-cpu", "host"},
	}
	for i, tt := range table {
		want := wanted[i]
		got := constructArguments(&tt)
		assertStringArray(t, got, want)

	}
}

func assertStringArray(t testing.TB, got []string, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("Length of %v(%d) and %v(%d) do not match", got, len(got), want, len(want))
	}
	for i, element := range got {
		if element != want[i] {
			t.Fatalf("%v and %v don't match , element %d(%v) differ", got, want, i, element)
		}

	}

}
