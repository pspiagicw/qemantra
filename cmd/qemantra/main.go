package main

import (
	"github.com/pspiagicw/qemantra/pkg/argparser"
	"github.com/pspiagicw/qemantra/pkg/config"
)

const VERSION = "0.0.1"

func main() {
	config.EnsureSystemReady()
	options := argparser.ParseArguments(VERSION)
	argparser.ParseAndRun(options, VERSION)
}
