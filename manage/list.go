package manage

import (
	"flag"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/help"
)

func ListVM(args []string) {
	flag := flag.NewFlagSet("qemantra list", flag.ExitOnError)

	flag.Usage = help.HelpList

	flag.Parse(args)

	machines := getMachines()

	if len(machines) == 0 {
		goreland.LogFatal("No machines found")
	}

	headers := []string{"Name", "CPU", "RAM", "Disk"}

	rows := [][]string{}

	for _, machine := range machines {
		row := []string{machine.Name, machine.CpuCores, machine.MemSize, machine.DiskPath}
		rows = append(rows, row)
	}

	goreland.LogTable(headers, rows)
}
