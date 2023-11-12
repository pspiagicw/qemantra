package execute

import (
	"os"
	"os/exec"

	"github.com/pspiagicw/goreland"
)

type Executor interface {
	Execute(cmd string, options []string) error
	GetCommand() []string
}
type SystemExecutor struct {
}

func (s *SystemExecutor) Execute(cmd string, options []string) error {
	command := exec.Command(cmd, options...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	// goreland.Printf("Executing '%s' on your operating system", command.String())
	goreland.LogExecSimple(command.String())
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
func (s *SystemExecutor) GetCommand() []string {
	return []string{}
}

func GetExecutor() Executor {
	return &SystemExecutor{}
}
