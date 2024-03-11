package argparse

import (
	"flag"

	"github.com/pspiagicw/qemantra/help"
)

type Opts struct {
	Args    []string
	Version string
}

func ParseOptions(version string) *Opts {

	opts := new(Opts)
	flag.Usage = func() {
		help.HelpUsage(version)
	}

	opts.Version = version
	flag.Parse()

	opts.Args = flag.Args()

	return opts
}
