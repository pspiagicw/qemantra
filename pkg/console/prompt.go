package console

import "fmt"

const BANNER string = `
                                 | |            
  __ _  ___ _ __ ___   __ _ _ __ | |_ _ __ __ _ 
 / _' |/ _ \ '_ ' _ \ / _' | '_ \| __| '__/ _' |
| (_| |  __/ | | | | | (_| | | | | |_| | | (_| |
 \__, |\___|_| |_| |_|\__,_|_| |_|\__|_|  \__,_|
    | |                                         
    |_|
`

func ShowBanner(version string) {
	fmt.Print(BANNER)
	fmt.Println()
	fmt.Printf("\tVersion %s\n", version)
	fmt.Println("\tControl QEMU like magic!")
	fmt.Println("\tWelcome to qemantra")
	fmt.Println("\tRun `qemantra -h` for help")
	fmt.Println()
}
func ShowSubcommands() {
	fmt.Println("Available subcommands are [list,run,create,rename,edit]")
}
