package main

import (
	// "fmt"

	"github.com/pspiagicw/lazyqemu/pkg/argparser"
	// "github.com/pspiagicw/lazyqemu/pkg/config"
	// "github.com/pspiagicw/lazyqemu/pkg/dirs"
	// "github.com/pspiagicw/lazyqemu/pkg/image"
	// "github.com/pspiagicw/lazyqemu/pkg/runner"
)

func main() {
	// config := config.GetConfig()
	// fmt.Println(config)
	// for _, file := range dirs.ListDirs(config.MachineDir) {
	// 	runner.RunMachine(config, file)
	// }
	// imageInstance := image.Image{
	// 	Name: "Arch Linux.img",
	// 	Type: image.IMAGE_RAW,
	// 	Size: "20G",
	// }
	// imagePath := image.CreateImage(&imageInstance)
	// template := &runner.RunnerTemplate{
	// 	Name:      "Arch Linux",
	// 	ImagePath: imagePath,
	// }
	// runner.CreateMachine(template)
	argparser.ParseArguments()
}
