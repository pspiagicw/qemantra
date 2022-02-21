package main

import (
	"github.com/pspiagicw/qemantra/pkg/argparser"
)
const VERSION = "0.0.1"

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
	options := argparser.ParseArguments(VERSION)
	argparser.ParseAndRun(options , VERSION)
	// example := &image.Image{
	// 	Type: "raw",
	// 	Name: "gentoo.img",
	// 	Size: "10G",
	// }
	// fmt.Println(image.CreateImage(example))
	// cr := &creator.MachineCreator{
	// 	Name:       "Gentoo Linux",
	// 	NoDisk:     false,
	// 	MemSize:    "4G",
	// 	CpuCores:   "2",
	// 	Diskname:   "hello.img",
	// 	Disksize:   "10G",
	// 	Diskformat: "raw",
	// }
	// creator.CreateNewMachine(cr)
}
