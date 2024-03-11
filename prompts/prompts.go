package prompt

import (
	"github.com/manifoldco/promptui"
	"github.com/pspiagicw/goreland"
)

func QuestionPrompt(label string, validation func(string) error, defaultValue string) string {

	if defaultValue != "" {
		label = label + "(" + defaultValue + ") "
	}

	prompt := promptui.Prompt{Label: label, Validate: validation}

	value, err := prompt.Run()

	if err != nil {
		goreland.LogFatal("Something went wrong: %q", err)
	}

	return value
}

func SelectionPrompt(label string, choices []string) string {

	prompt := promptui.Select{Label: label, Items: choices}

	_, value, err := prompt.Run()

	if err != nil {
		goreland.LogFatal("Something went wrong: %q", err)
	}

	return value
}
func ConfirmPrompt(label string) bool {
	answer := QuestionPrompt(label, func(string) error { return nil }, "")

	if answer == "y" || answer == "Y" {
		return true
	}
	return false
}
