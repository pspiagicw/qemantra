package main

import (
	"github.com/pspiagicw/qemantra/pkg/argparse"
	"github.com/pspiagicw/qemantra/pkg/config"
)

// VERSION variable which declares which version currently we are having
// var VERSION = "0.0.1"
var VERSION string

// This is the main function.
func main() {
	config.EnsureSystemReady()
	argparse.ParseOptions(VERSION)
}
