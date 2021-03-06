package main

import (
	"github.com/pspiagicw/qemantra/pkg/argparser"
	"github.com/pspiagicw/qemantra/pkg/config"
)

// VERSION variable which declares which version currently we are having
// var VERSION = "0.0.1"
var VERSION string

// This is the main function.
func main() {
	config.EnsureSystemReady()
	options := argparser.ParseFlags(VERSION)
	argparser.ParseOptions(options, VERSION)
}
