// MIT License
//
// Copyright (c) 2022 pspiagicw
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package argparse

import (
	"flag"
	"fmt"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/pkg/console"
)

func ParseOptions(version string) {

	flag.Usage = PrintUsage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		PrintUsage()
		goreland.LogFatal("No subcommands provided!")
	}

	cmd, args := flag.Args()[0], flag.Args()[1:]

	handlers := map[string]func([]string){
		"edit":   edit,
		"create": create,
		"list":   list,
		"run":    run,
		"rename": rename,
		"check":  check,
		"help":   help,
	}

	handleFunc, exists := handlers[cmd]

	if exists {
		handleFunc(args)
	} else {
		PrintUsage()
	}
}
func help(args []string) {
	PrintUsage()
}
func PrintUsage() {
	console.ShowBanner()
	fmt.Println("qemantra is a command line tool for creating, running and managing virtual machines using QEMU/KVM.")
	fmt.Println()
	fmt.Println("The available commands.")
	fmt.Println()
	fmt.Println("\thelp\t\tShow this message and exit.")
	fmt.Println("\tcheck\t\tCheck dependencies for qemantra.")
	fmt.Println("\tcreate\t\tCreate a virtual machine.")
	fmt.Println("\tlist\t\tList all virtual machines.")
	fmt.Println("\trun\t\tRun a virtual machine.")
	fmt.Println("\trename\t\tRename a virtual machine.")
	fmt.Println("\tedit\t\tEdit configuration of a virtual machine.")
}
