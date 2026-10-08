package validation

import (
	"errors"
	"strings"
)

func ValidateInputs(name string, email string, tickets int) error {
	if len(name) < 2 {
		err := errors.New("Invalid Name. Name should have atleast 2 characters")
		return err
	}
	if !strings.Contains(email, "@") {
		err := errors.New("Invalid Email. Email should have a @ sign")
		return err
	}
	if tickets <= 0 {
		err := errors.New("Invalid ticket count. Ticket count should be > 0")
		return err
	}
	return nil
}
