package swissjson

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type JSONValidation struct {
	InputPath  string `validate:"required,file"`
	OutputPath string `validate:"omitempty"`
	JSON       string `validate:"required,json"`
}

func (j *JSONValidation) validated() error {
	validate := validator.New()
	if err := validate.Struct(j); err != nil {
		var errorMessages []string

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New(err.Error())
		}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("%s should be %s (%s)", err.Field(), err.Tag(), err.Param()))
		}

		return errors.New(strings.Join(errorMessages, "\n"))
	}

	return nil
}

func (j *JSONValidation) Minify() {
	for _, ch := range []string{",", ":", "[", "]", "{", "}"} {
		skipped := ch
		if ch == "[" || ch == "]" || ch == "{" || ch == "}" {
			skipped = fmt.Sprintf("\\%s", ch)
		}
		reg := regexp.MustCompile(fmt.Sprintf(`(?mi)(\s+%s|%s\s+)`, skipped, skipped))
		j.JSON = reg.ReplaceAllString(j.JSON, ch)
	}
}
