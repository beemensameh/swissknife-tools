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

	"github.com/beemensameh/swissknife-tools/internal/color"
	"github.com/beemensameh/swissknife-tools/internal/validate"
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
	Path      string
	Algorithm AlgorithmType
	content   []byte
}

func (hf *HashFile) Validated() error {
	if err := validate.FileOrDirValidation(hf.Path, false); err != nil {
		return err
	}

	c, err := os.ReadFile(hf.Path)
	if err != nil {
		return err
	}

	hf.content = c

	if hf.Algorithm != "" {
		switch hf.Algorithm {
		case CRC32, MD5, SHA1, SHA256, SHA384, SHA512, All:
			return nil
		default:
			return errors.New(color.SprintfColor("Invalid algorithm", color.Red))
		}
	}
	return nil
}

func (hf *HashFile) Hash() (string, error) {
	switch hf.Algorithm {
	case CRC32:
		return fmt.Sprintf("CRC32: %x", crc32.ChecksumIEEE(hf.content)), nil
	case MD5:
		return fmt.Sprintf("MD5: %x", md5.Sum(hf.content)), nil
	case SHA1:
		return fmt.Sprintf("SHA1: %x", sha1.Sum(hf.content)), nil
	case SHA256:
		return fmt.Sprintf("SHA256: %x", sha256.Sum256(hf.content)), nil
	case SHA384:
		return fmt.Sprintf("SHA384: %x", sha512.Sum384(hf.content)), nil
	case SHA512:
		return fmt.Sprintf("SHA512: %x", sha512.Sum512(hf.content)), nil
	case All:
		return fmt.Sprintf(
			"CRC32:\t%x\nMD5:\t%x\nSHA1:\t%x\nSHA256:\t%x\nSHA384:\t%x\nSHA512:\t%x",
			crc32.ChecksumIEEE(hf.content),
			md5.Sum(hf.content),
			sha1.Sum(hf.content),
			sha256.Sum256(hf.content),
			sha512.Sum384(hf.content),
			sha512.Sum512(hf.content),
		), nil
	}
	return "", nil
}
