package swisshashing

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var HashFileCmd = &cli.Command{
	Name:   "hash:file",
	Usage:  "Hash a file",
	Action: hashFileAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Usage:    "The path for the file (required)",
			Aliases:  []string{"p"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "algorithm",
			Usage:   "The algorithm for hashing",
			Aliases: []string{"algo"},
			Value:   string(SHA256),
		},
	},
}

func hashFileAction(cliContext *cli.Context) error {
	return hashFile(&HashFile{
		Path:      cliContext.String("path"),
		Algorithm: AlgorithmType(cliContext.String("algorithm")),
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
