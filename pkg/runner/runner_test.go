package runner

import "testing"

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
	}
	wanted := [][]string{
		{"-enable-kvm", "-hda", "test.img", "-cpu" , "host"},
		{"-m", "4G", "-enable-kvm", "-hda", "test.img" , "-cpu" , "host"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-cpu" , "host" ,  "-smp", "2"},
		{"-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img",  "-cpu" , "host" ,"-smp", "2"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img" , "-cpu" , "host"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "menu=on", "-cpu" , "host" , "-hdb", "externaltest.img"},
		{"-m", "4G", "-enable-kvm", "-cdrom", "test.iso", "-hda", "test.img", "-boot", "d", "-cpu" , "host" ,  "-smp", "2", "-hdb", "externaltest.img" },
	}
	for i, tt := range table {
		want := wanted[i]
		got := constructOptions(&tt)
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
