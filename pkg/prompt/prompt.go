package prompt

import "fmt"

const BANNER string = `
░▄▀▄░█▀▀░█▄█░█▀█░█▀█░▀█▀░█▀▄░█▀█
░█\█░█▀▀░█░█░█▀█░█░█░░█░░█▀▄░█▀█
░░▀\░▀▀▀░▀░▀░▀░▀░▀░▀░░▀░░▀░▀░▀░▀`

func ShowBanner(version string) {
	fmt.Println(BANNER)
	fmt.Printf("Version %s\n", version)
	fmt.Println("Control QEMU like magic!")
	fmt.Println("Welcome to Qemantra")
	fmt.Println("Run `qemantra -h` for help")
}
