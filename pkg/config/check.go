package config

import (
	"os"
	"os/exec"
	"runtime"

	// log "github.com/pspiagicw/colorlog"
	"github.com/pspiagicw/goreland"
)

const OVMF_PATH = "/usr/share/ovmf"
const GOOS_WINDOWS = "windows"

func EnsureSystemReady() {
	if !checkSystemType() {
		goreland.LogFatal("Sorry only Linux is tested and supported!")
	}
	if !checkQEMU() {
		goreland.LogInfo("Error detecting 'qemu-img' command , please ensure it exists")
	}

}

func checkSystemType() bool {
	if runtime.GOOS == GOOS_WINDOWS {
		return false
	}
	return true
}

func checkQEMU() bool {
	_, err := exec.LookPath("qemu-img")
	if err != nil {
		return false
	}
	return true

}

func EnsureUEFIReady() {
	if !CheckUEFI() {
		goreland.LogFatal("UEFI not supported! Install `OVMF` to enable support")
	}
}

func CheckUEFI() bool {
	_, err := os.Stat(OVMF_PATH)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func PerformCheck() {
	if CheckUEFI() {
		goreland.LogSuccess("✓ UEFI Support enabled!")
	} else {
		goreland.LogError("! UEFI Support disabled!")
	}

	if checkQEMU() {
		goreland.LogSuccess("✓ QEMU Installed!")
	} else {
		goreland.LogFatal("! QEMU Not Installed!")
	}
	if checkSystemType() {
		goreland.LogSuccess("✓ Your Platform is supported!")
	} else {
		goreland.LogError("! Your Platform is not supported!")
	}
}
