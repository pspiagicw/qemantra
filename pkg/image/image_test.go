package image

import (
	"testing"
)

type TestExecutor struct {
	Command []string
}

func (t *TestExecutor) Execute(command string, options []string) error {
	t.Command = []string{command}
	t.Command = append(t.Command, options...)
	return nil
}
func (t *TestExecutor) GetCommand() []string {
	return t.Command

}

type TestConfig struct {
}

func (t *TestConfig) GetImageDir() string {
	return "testdir/images"
}
func (t *TestConfig) GetMachineDir() string {
	return ""
}
func TestCreateImage(t *testing.T) {
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
	ConfigProvider = &TestConfig{}
	for i, tt := range tables {
		CreateImage(&tt)
		got := ExecProvider.GetCommand()
		want := wanted[i]
		assertEqual(t, got, want)

	}
	ExecProvider = previousExecProvider
	ConfigProvider = previousConfigProvider
}

func assertEqual(t testing.TB, got []string, want []string) {
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
