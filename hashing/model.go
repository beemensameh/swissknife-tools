package swisshashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AlgorithmType string

const (
	CRC32  AlgorithmType = "crc32"
	MD5    AlgorithmType = "md5"
	SHA1   AlgorithmType = "sha1"
	SHA256 AlgorithmType = "sha256"
	SHA384 AlgorithmType = "sha384"
	SHA512 AlgorithmType = "sha512"
	All    AlgorithmType = "all"
)

type HashFile struct {
	Path      string        `validate:"required,file"`
	Algorithm AlgorithmType `validate:"omitempty"`
}

func (hashFile *HashFile) Validated() error {
	validate := validator.New()
	err := validate.Struct(hashFile)
	if err != nil {
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

func (hashFile *HashFile) Hash() (string, error) {
	content, err := os.ReadFile(hashFile.Path) // the file is inside the local directory
	if err != nil {
		return "", err
	}

	switch hashFile.Algorithm {
	case CRC32:
		return fmt.Sprintf("CRC32: %x", crc32.ChecksumIEEE(content)), nil
	case MD5:
		return fmt.Sprintf("MD5: %x", md5.Sum(content)), nil
	case SHA1:
		return fmt.Sprintf("SHA1: %x", sha1.Sum(content)), nil
	case SHA256:
		return fmt.Sprintf("SHA256: %x", sha256.Sum256(content)), nil
	case SHA384:
		return fmt.Sprintf("SHA384: %x", sha512.Sum384(content)), nil
	case SHA512:
		return fmt.Sprintf("SHA512: %x", sha512.Sum512(content)), nil
	case All:
		return fmt.Sprintf(
			"CRC32:\t%x\nMD5:\t%x\nSHA1:\t%x\nSHA256:\t%x\nSHA384:\t%x\nSHA512:\t%x",
			crc32.ChecksumIEEE(content),
			md5.Sum(content),
			sha1.Sum(content),
			sha256.Sum256(content),
			sha512.Sum384(content),
			sha512.Sum512(content),
		), nil
	default:
		return "", errors.New("invalid algorithm")
	}
}
