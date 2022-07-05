package image

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
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

func setupTest(t *testing.T, files []string) (string, func(t *testing.T)) {
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

func TestCreateImage(t *testing.T) {
	t.Run("image already exists", func(t *testing.T) {
		path, tearDown := setupTest(t, []string{"hello"})
		defer tearDown(t)
		tables := []Image{
			{
				Name: "hello",
				Type: "",
				Size: "10G",
			},
		}
		wanted := []string{
            "",
        }
		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		// Give error on purpose
		ExecProvider = &TestExecutor{
			errorExecute: true,
		}
		ConfigProvider = &TestConfig{
			imagepath: path,
		}
		for i, tt := range tables {
			got, err := CreateImage(&tt)
			want := wanted[i]
            assert.Equal(t , got , want)
            assert.Error(t , err)

		}
		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})

	t.Run("executor does not give error", func(t *testing.T) {
		const imageDir = "testdir/images"
		tables := []Image{
			{
				Name: "example",
				Type: "",
				Size: "10G",
			},
			{
				Name: "example",
				Type: "qcow2",
				Size: "5G",
			},
			{
				Name: "hello",
				Type: "",
				Size: "",
			},
			{
				Name: "hello",
				Type: "qcow2",
				Size: "",
			},
		}
		wanted := [][]string{
			{"qemu-img", "create", "-f", "raw", "testdir/images/example", "10G"},
			{"qemu-img", "create", "-f", "qcow2", "testdir/images/example", "5G"},
			{"qemu-img", "create", "-f", "raw", "testdir/images/hello", "10G"},
			{"qemu-img", "create", "-f", "qcow2", "testdir/images/hello", "10G"},
		}
		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		ExecProvider = &TestExecutor{}
		ConfigProvider = &TestConfig{
			imagepath: "testdir/images",
		}
		for i, tt := range tables {
            _ , err := CreateImage(&tt)
			got := ExecProvider.GetCommand()
			want := wanted[i]
            assert.ElementsMatch(t, got , want)
            assert.Nil(t , err)

		}
		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
	t.Run("executor gives an error", func(t *testing.T) {
		const imageDir = "testdir/images"
		tables := []Image{
			{
				Name: "example",
				Type: "",
				Size: "10G",
			},
		}
		wanted := []string{""}
		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		// Give error on purpose
		ExecProvider = &TestExecutor{
			errorExecute: true,
		}
		ConfigProvider = &TestConfig{
			imagepath: "testdir/images",
		}
		for i, tt := range tables {
			got, err := CreateImage(&tt)
			want := wanted[i]
            assert.Equal(t , got , want)
            assert.Error(t , err)

		}
		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
}

func TestFindImage(t *testing.T) {
	t.Run("Image exists", func(t *testing.T) {
		path, tearDown := setupTest(t, []string{"hello", "sello"})
		defer tearDown(t)

		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		ExecProvider = &TestExecutor{}
		ConfigProvider = &TestConfig{
			imagepath: path,
		}

		want := filepath.Join(path, "hello")
		got := FindImage("hello")

        assert.Equal(t , got , want)

		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider

	})
	t.Run("Image does not exist!", func(t *testing.T) {
		path, tearDown := setupTest(t, []string{"sello"})
		defer tearDown(t)

		previousExecProvider := ExecProvider
		previousConfigProvider := ConfigProvider
		ExecProvider = &TestExecutor{}
		ConfigProvider = &TestConfig{
			imagepath: path,
		}

		want := ""
		got := FindImage("hello")

		// if got != want {
		// 	t.Errorf("got %v , wanted %v", got, want)
		// }
        assert.Equal(t , got , want)

		ExecProvider = previousExecProvider
		ConfigProvider = previousConfigProvider
	})

}

func assertStringArray(t testing.TB, got []string, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Length of %v(%d) differs from %v(%d)", got, len(got), want, len(want))
	}
	for i, element := range got {
		if element != want[i] {
			t.Errorf("Element %d do not match , %v and %v", i, element, want[i])
		}
	}
}
