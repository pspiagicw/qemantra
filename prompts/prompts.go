package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pspiagicw/goreland"
)

func QuestionPrompt(label string, validation func(string) error, defaultValue string) string {

	prompt := survey.Input{
		Message: label,
		Default: defaultValue,
	}

	var value string

	validator := func(ans interface{}) error {
		strValue, ok := ans.(string)
		if !ok {
			return fmt.Errorf("Can't convert to string")
		}

		return validation(strValue)
	}

	survey.AskOne(&prompt, &value, survey.WithValidator(validator), survey.WithValidator(survey.Required))
	if value == "" {
		goreland.LogFatal("Nothing was provided")
	}

	return value
}

func SelectionPrompt(label string, choices []string) string {
	prompt := &survey.Select{
		Message: label,
		Options: choices,
	}

	var value string

	survey.AskOne(prompt, &value)
	if value == "" {
		goreland.LogFatal("Nothing was selected")
	}

	return value
}
func ConfirmPrompt(label string) bool {

	var confirm bool

	prompt := &survey.Confirm{
		Message: label,
	}

	survey.AskOne(prompt, &confirm)

	return confirm

}
