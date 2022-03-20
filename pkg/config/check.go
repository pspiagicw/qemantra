package config

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

const OVMF_PATH = "/usr/share/ovmf"
const GOOS_WINDOWS = "windows"
func EnsureSystemReady() {
	if !CheckSystemType() {
		log.Fatalf("Sorry only Linux is tested and supported!")
	}
	if !CheckQEMU() {
		log.Fatalln("Error detecting 'qemu-img' command %v , please ensure it exists")
	}

}
func CheckSystemType() bool {
	if runtime.GOOS == GOOS_WINDOWS {
		return false
	}
	return true
	
}
func CheckQEMU() bool {
	_, err := exec.LookPath("qemu-img")
	if err != nil {
		return false
	}
	return true
	
}
func EnsureUEFIReady() {
	if !CheckUEFI() {
		log.Fatalf("UEFI not supported! Install `OVMF` to enable support")
	}
}
func CheckUEFI() bool {
	_ , err := os.Stat(OVMF_PATH)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
