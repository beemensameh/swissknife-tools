package swisstime

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func newTimeFormat(format int) string {
	switch format {
	case 1:
		return time.ANSIC
	case 2:
		return time.UnixDate
	case 3:
		return time.RubyDate
	case 4:
		return time.RFC822
	case 5:
		return time.RFC822Z
	case 6:
		return time.RFC850
	case 7:
		return time.RFC1123
	case 8:
		return time.RFC1123Z
	case 9:
		return time.RFC3339
	case 10:
		return time.RFC3339Nano
	case 11:
		return time.Kitchen
	case 12:
		return time.Stamp
	case 13:
		return time.StampMilli
	case 14:
		return time.StampMicro
	case 15:
		return time.StampNano
	default:
		return ""
	}
}

type TimeCLI struct {
	Format   string
	Update   bool
	Interval int `validate:"omitempty,gte=1"`
}

func (uuidCLI *TimeCLI) validated() error {
	validate := validator.New()

	err := validate.Struct(uuidCLI)
	if err != nil {
		var errorMessages []string

		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("%s should be %s (%s)", err.Field(), err.Tag(), err.Param()))
		}

		return errors.New(strings.Join(errorMessages, "\n"))
	}

	return nil
}
