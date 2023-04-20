package swisshashing

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	path, algo string

	HashCmd = &cobra.Command{
		Use:   "hash",
		Short: "Use hast tool",
	}
)

func init() {
	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "Hash a file",
		Long:  "Hash file with many algorithms like (md5, SHA256, SHA512, etc...)",
		RunE:  hashFileAction,
	}
	fileCmd.Flags().StringVarP(&path, "path", "p", "", "The path for the file")
	fileCmd.Flags().StringVarP(&algo, "algorithm", "a", "", "The algorithm for hashing")
	if err := fileCmd.MarkFlagRequired("path"); err != nil {
		log.Fatal(err)
	}
	HashCmd.AddCommand(fileCmd)
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
