package swisstime

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var timeFromat = map[int]string{
	1:  time.ANSIC,
	2:  time.UnixDate,
	3:  time.RubyDate,
	4:  time.RFC822,
	5:  time.RFC822Z,
	6:  time.RFC850,
	7:  time.RFC1123,
	8:  time.RFC1123Z,
	9:  time.RFC3339,
	10: time.RFC3339Nano,
	11: time.Kitchen,
	12: time.Stamp,
	13: time.StampMilli,
	14: time.StampMicro,
	15: time.StampNano,
}

func newTimeFormat(format int) string {
	return timeFromat[format]
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
