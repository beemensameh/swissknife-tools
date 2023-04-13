package swisshashing

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	path, algo string

	HashFileCmd = &cobra.Command{
		Use:     "hash:file",
		Short:   "Hash a file",
		Long:    "Hash file with many algorithms like (md5, SHA256, SHA512, etc...)",
		Aliases: []string{"hash:f"},
		RunE:    hashFileAction,
	}
)

func init() {
	HashFileCmd.Flags().StringVarP(&path, "path", "p", "", "The path for the file (required)")
	HashFileCmd.Flags().StringVarP(&algo, "algorithm", "a", "", "The algorithm for hashing")
}

func hashFileAction(cmd *cobra.Command, args []string) error {
	return hashFile(&HashFile{
		Path:      path,
		Algorithm: AlgorithmType(algo),
	})
}

func hashFile(hf *HashFile) error {
	err := hf.Validated()
	if err != nil {
		return err
	}

	hash, err := hf.Hash()
	if err != nil {
		return err
	}

	fmt.Println(hash)

	return nil
}
