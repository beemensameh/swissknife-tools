package swisshashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
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
	content, err := ioutil.ReadFile(hashFile.Path) // the file is inside the local directory
	if err != nil {
		return "", err
	}

	switch hashFile.Algorithm {
	case CRC32:
		return fmt.Sprintf("%x", crc32.ChecksumIEEE(content)), nil
	case MD5:
		return fmt.Sprintf("%x", md5.Sum(content)), nil
	case SHA1:
		return fmt.Sprintf("%x", sha1.Sum(content)), nil
	case SHA256:
		return fmt.Sprintf("%x", sha256.Sum256(content)), nil
	case SHA384:
		return fmt.Sprintf("%x", sha512.Sum384(content)), nil
	case SHA512:
		return fmt.Sprintf("%x", sha512.Sum512(content)), nil
	case All:
		return fmt.Sprintf(
			"crc32:\t%x\nmd5:\t%x\nsha1:\t%x\nsha256:\t%x\nsha384:\t%x\nsha512:\t%x",
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
