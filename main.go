package main

import (
	"github.com/pspiagicw/qemantra/pkg/argparse"
	"github.com/pspiagicw/qemantra/pkg/config"
)

var VERSION string

func main() {
	config.EnsureSystemReady()
	argparse.ParseOptions(VERSION)
}
