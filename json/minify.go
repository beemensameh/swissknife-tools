package swissjson

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var JSONMinifyCmd = &cli.Command{
	Name:   "json:minify",
	Usage:  "Minify json",
	Action: jsonMinifyAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Usage:    "The path for the input file (required)",
			Aliases:  []string{"in"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Usage:   "The path for the output file (if this flag doesn't add will print it in terminal)",
			Aliases: []string{"o"},
		},
	},
}

func jsonMinifyAction(cliContext *cli.Context) error {
	return jsonMinify(&JSONValidation{
		InputPath:  cliContext.String("input"),
		OutputPath: cliContext.String("output"),
	})
}

func jsonMinify(j *JSONValidation) error {
	file, err := os.ReadFile(j.InputPath)
	if err != nil {
		return err
	}

	j.JSON = string(file)

	err = j.validated()
	if err != nil {
		return err
	}

	j.Minify()

	if j.OutputPath != "" {
		if err := os.WriteFile(j.OutputPath, []byte(j.JSON), 0o600); err != nil {
			return err
		}
	} else {
		fmt.Println(j.JSON)
	}

	return nil
}
