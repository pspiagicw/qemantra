package prompt

import "fmt"

const BANNER string = `
░▄▀▄░█▀▀░█▄█░█▀█░█▀█░▀█▀░█▀▄░█▀█
░█\█░█▀▀░█░█░█▀█░█░█░░█░░█▀▄░█▀█
░░▀\░▀▀▀░▀░▀░▀░▀░▀░▀░░▀░░▀░▀░▀░▀
`

func ShowBanner(name string) {
	fmt.Println(BANNER)
	fmt.Println("Version 0.0.1")
	fmt.Println("Control QEMU like magic!")
	fmt.Println("Welcome to Qemantra")
	fmt.Println("Run `qemantra -h` for help")
}
