package swissjson

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/urfave/cli/v2"
)

var JsonMinifyCmd = &cli.Command{
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
	return jsonMinify(&JsonValidation{
		InputPath:  cliContext.String("input"),
		OutputPath: cliContext.String("output"),
	})
}

func jsonMinify(jsonValidation *JsonValidation) error {
	reg := regexp.MustCompile(`(?mi)^\s+|\n`)

	file, err := ioutil.ReadFile(jsonValidation.InputPath)
	if err != nil {
		return err
	}

	jsonValidation.JSON = string(file)

	err = jsonValidation.validated()
	if err != nil {
		return err
	}

	res := reg.ReplaceAllString(jsonValidation.JSON, "")
	if jsonValidation.OutputPath != "" {
		if err := ioutil.WriteFile(jsonValidation.OutputPath, []byte(res), 0644); err != nil {
			return err
		}
	} else {
		fmt.Println(res)
	}

	return nil
}
