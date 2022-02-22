package config

import (
	"log"
	"os/exec"
	"runtime"
)

func EnsureSystemReady() {
	if runtime.GOOS == "windows" {
		log.Fatalf("Sorry only Linux is tested and supported!")
	}
	_ , err := exec.LookPath("qemu-img")
	if err != nil {
		log.Fatalf("Error detecting 'qemu-img' command %v , please ensure it exists" , err)
	}
	
}
