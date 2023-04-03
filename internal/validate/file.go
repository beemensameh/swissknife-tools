package validate

import "os"

func IsFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
