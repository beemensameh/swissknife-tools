package swissfiles

import "github.com/beemensameh/swissknife-tools/internal/validate"

type ServeStatic struct {
	Path string
}

func (s *ServeStatic) Validated() error {
	return validate.FileOrDirValidation(s.Path, true)
}
