package swissjson

import (
	"fmt"
	"os"
	"regexp"

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

func jsonMinify(jsonValidation *JSONValidation) error {
	reg := regexp.MustCompile(`(?mi)(^\s+|\n)`)
	reg2 := regexp.MustCompile(`(?mi)(\s+:|:\s+)`)

	file, err := os.ReadFile(jsonValidation.InputPath)
	if err != nil {
		return err
	}

	jsonValidation.JSON = string(file)

	err = jsonValidation.validated()
	if err != nil {
		return err
	}

	res := reg.ReplaceAllString(jsonValidation.JSON, "")
	res2 := reg2.ReplaceAllString(res, ":")
	if jsonValidation.OutputPath != "" {
		if err := os.WriteFile(jsonValidation.OutputPath, []byte(res2), 0o600); err != nil {
			return err
		}
	} else {
		fmt.Println(res2)
	}

	return nil
}
