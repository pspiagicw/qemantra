package runner

import (
	"bytes"
	"log"
	"os/exec"
)

type Runner struct {
	Name          string `json:"name"`
	DrivePath     string `json:"drivePath"`
	SystemCommand string `json:"systemCommand"`
	MemSize       string `json:"memSize"`
	CpuCores      string `json:"cpuCores"`
	Iso           string `json:"-"`
}

func RunMachine(runner *Runner) {
	startMachine(runner)
}
func startMachine(runner *Runner) {
	options := constructOptions(runner)
	cmd := exec.Command(runner.SystemCommand, options... )

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		log.Printf("Some error occured %v", err)
		log.Fatalf("The err %s", out.String())
	}
}
func constructOptions(runner *Runner) []string {
	options := []string{
	}
	options = append(options , getMemOptions(runner)...)
	options = append(options,getMiscOptions(runner)...)
	options = append(options,getIsoOptions(runner)...)
	options = append(options , getDriveOptions(runner)...)
	return options
}
func getIsoOptions(runner *Runner) []string {
	// fmt.Println("-cdrom " + runner.Iso)
	option := []string{"-cdrom" , runner.Iso}
	return option

}
func getDriveOptions(runner *Runner) []string {
	if runner.DrivePath != "" {
	    option := []string{ "-hda" , runner.DrivePath  , }
	    return option
	}
	return []string{}
}
func getMemOptions(runner *Runner) []string {
	option := []string{ "-m", runner.MemSize }
	return option
}
func getCpuOptions(runner *Runner) []string {
	return []string{""}
}
func getMiscOptions(runner *Runner) []string {
	return []string{"-enable-kvm"}
}
func getBootOptions(runner *Runner) string {
	return ""
}
