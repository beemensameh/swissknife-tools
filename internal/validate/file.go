package validate

import (
	"errors"
	"os"

	"github.com/beemensameh/swissknife-tools/internal/color"
)

func FileOrDirValidation(path string, isDir bool) error {
	if f, err := os.Stat(path); err != nil {
		return errors.New(color.SprintfColor("Should pass a valid path file", color.Red))
	} else if f.IsDir() && !isDir {
		return errors.New(color.SprintfColor("Should pass a file not a directory", color.Red))
	}
	return nil
}

func IsFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
