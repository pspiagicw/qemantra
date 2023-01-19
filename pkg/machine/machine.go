package machine

type Machine struct {
	NoDisk        bool   `json:"-"`
	DiskName      string `json:"-"`
	DiskFormat    string `json:"-"`
	DiskSize      string `json:"-"`
	Name          string `json:"name"`
	DrivePath     string `json:"drivePath"`
	SystemCommand string `json:"systemCommand"`
	MemSize       string `json:"memSize"`
	CpuCores      string `json:"cpuCores"`
	Iso           string `json:"-"`
	ExternalDisk  string `json:"-"`
	Boot          string `json:"-"`
	UEFI          string `json:"-"`
	KVM           bool   `json:"-"`
}
