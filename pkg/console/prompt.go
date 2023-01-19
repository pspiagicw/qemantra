package console

import "fmt"

const BANNER string = `
░▄▀▄░█▀▀░█▄█░█▀█░█▀█░▀█▀░█▀▄░█▀█
░█\█░█▀▀░█░█░█▀█░█░█░░█░░█▀▄░█▀█
░░▀\░▀▀▀░▀░▀░▀░▀░▀░▀░░▀░░▀░▀░▀░▀`

func ShowBanner(version string) {
	fmt.Println(BANNER)
	fmt.Println()
	fmt.Printf("Version %s\n", version)
	fmt.Println("Control QEMU like magic!")
	fmt.Println("Welcome to qemantra")
	fmt.Println("Run `qemantra -h` for help")
}
func ShowSubcommands() {
	fmt.Println("Available subcommands are [list,run,create,rename,edit]")
}
