package runner

import "fmt"

type RunnerTemplate struct {
	Name      string
	ImagePath string
}

func CreateMachine(template *RunnerTemplate) {
	fmt.Println(template)
}
