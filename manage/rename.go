package manage

import (
	"flag"
	"os"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/qemantra/help"
	prompt "github.com/pspiagicw/qemantra/prompts"
	"github.com/pspiagicw/qemantra/validators"
)

func parseRenameArgs(args []string) {
	flag := flag.NewFlagSet("qemantra rename", flag.ExitOnError)

	flag.Usage = help.HelpRename

	flag.Parse(args)

}

func RenameVM(args []string) {

	parseRenameArgs(args)

	name, selected := selectMachine()

	newName := prompt.QuestionPrompt("New Name", validators.NameValidator, "")

	if newName == name {
		goreland.LogFatal("New name is same as old.")
	}

	deleteFile(getMachinePath(selected))

	selected.Name = newName

	saveToDisk(selected, getMachinePath(selected))
}
func deleteFile(path string) {
	err := os.Remove(path)

	if err != nil {
		goreland.LogFatal("Error removing '%s': %v", path, err)
	}
}
