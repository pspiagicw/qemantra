package main

import (
	"github.com/pspiagicw/qemantra/argparse"
	"github.com/pspiagicw/qemantra/handle"
)

var VERSION string

func main() {
	opts := argparse.ParseOptions(VERSION)
	handle.HandleArgs(opts)
}
