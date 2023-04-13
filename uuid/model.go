package swissuuid

import (
	"errors"

	"github.com/beemensameh/swissknife-tools/internal/color"
	"github.com/google/uuid"
)

type UUIDCLI struct {
	Version  int         `validate:"required,gte=0,lte=5"`
	Domain   uuid.Domain `validate:"required_if=version 2,oneof=0 1 2"`
	Name     []byte      `validate:"omitempty"`
	Number   int64       `validate:"required,number"`
	Separate string      `validate:"omitempty"`
}

func (u *UUIDCLI) validated() error {
	if u.Version < 0 || u.Version > 5 {
		return errors.New(color.SprintfColor("The version should be between 0 to 5", color.Red))
	}
	if u.Version == 2 && u.Domain > 2 {
		return errors.New(color.SprintfColor("The uuid security type should be between 0 to 2", color.Red))
	}
	if u.Number < 1 {
		color.PrintlnColor("The number of generated uuids should be large than 0 (changed to 1)", color.Yellow)
	}

	return nil
}
