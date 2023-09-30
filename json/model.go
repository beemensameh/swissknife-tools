package swissjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/beemensameh/swissknife-tools/internal/color"
	"github.com/beemensameh/swissknife-tools/internal/validate"
)

type JSONValidation struct {
	Input  string
	Output string
	JSON   string
}

func (j *JSONValidation) validated() error {
	if err := validate.FileOrDirValidation(j.Input, false); err != nil {
		return err
	}

	file, err := os.ReadFile(j.Input)
	if err != nil {
		return err
	}
	if !json.Valid(file) {
		return errors.New(color.SprintfColor("The file not have valid json", color.Red))
	}

	j.JSON = string(file)

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
