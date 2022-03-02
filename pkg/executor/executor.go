package executor

import (
	"log"
	"os/exec"
)

type Executor interface {
	Execute(cmd string , options []string) error
	GetCommand() []string
}
type SystemExecutor struct {
}
func (s *SystemExecutor) Execute(cmd string , options []string) error {
	command := exec.Command(cmd , options...)
	log.Printf("Executing '%s' on your operating system", command.String())
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
