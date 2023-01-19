package argparse

import (
	"fmt"
	"strconv"
)

func coresValidator(input string) error {
	if len(input) == 0 {
		return fmt.Errorf("No input")

	}
	_, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return fmt.Errorf("%s is not a valid number.", input)
	}
	return nil

}
func ramValidator(input string) error {
	if len(input) == 0 {
		return fmt.Errorf("It is empty")
	}
	mtype := string(input[len(input)-1])
	input = string(input[:len(input)-1])

	if mtype != "M" && mtype != "G" {
		return fmt.Errorf("%s not a valid memory type", mtype)
	}

	_, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return fmt.Errorf("%s is not a valid number.", input)
	}
	return nil

}
func yesValidator(input string) error {
	if input != "y" && input != "Y" && input != "N" && input != "n" {
		return fmt.Errorf("Either y/n only.")
	}

	return nil
}
