package config

import (
	"os"
	"os/exec"
	"runtime"

    log "github.com/pspiagicw/colorlog"
)

const OVMF_PATH = "/usr/share/ovmf"
const GOOS_WINDOWS = "windows"

// The function which checks if the system is ready.
// Checks include
// - If system is Unix/Linux
// - If QEMU is installed
func EnsureSystemReady() {
	if !checkSystemType() {
		log.LogFatal("Sorry only Linux is tested and supported!")
	}
	if !checkQEMU() {
		log.LogFatal("Error detecting 'qemu-img' command , please ensure it exists")
	}

}

// Uses `runtime.GOOS` to check if system is Unix-based.
// Returns `true` if system Unix based else false.
func checkSystemType() bool {
	if runtime.GOOS == GOOS_WINDOWS {
		return false
	}
	return true

}

// Checks if the `qemu-img` command available in `$PATH` variable.
func checkQEMU() bool {
	_, err := exec.LookPath("qemu-img")
	if err != nil {
		return false
	}
	return true

}

// Checks if UEFI files , using `OVMF` are available .
func EnsureUEFIReady() {
	if !CheckUEFI() {
		log.LogFatal("UEFI not supported! Install `OVMF` to enable support")
	}
}

// Checks for the required OVMF directory and returns false if not present.
func CheckUEFI() bool {
	_, err := os.Stat(OVMF_PATH)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func PerformCheck() {
	if CheckUEFI() {
		log.LogSuccess("✓ UEFI Support enabled!")
	} else {
		log.LogError("! UEFI Support disabled!")
	}

	if checkQEMU() {
		log.LogSuccess("✓ QEMU Installed!")
	} else {
		log.LogFatal("! QEMU Not Installed!")
	}
	if checkSystemType() {
		log.LogSuccess("✓ Your Platform is supported!")
	} else {
		log.LogError("! Your Platform is not supported!")
	}
}
