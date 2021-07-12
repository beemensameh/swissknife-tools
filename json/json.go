package swissjson

import (
	"encoding/json"
	"errors"
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

func isJSON(s string) bool {
	var (
		jsFormat1 map[string]interface{}
		jsFormat2 []map[string]interface{}
	)

	return json.Unmarshal([]byte(s), &jsFormat1) == nil || json.Unmarshal([]byte(s), &jsFormat2) == nil
}

func jsonMinifyAction(cliContext *cli.Context) error {
	var (
		inputPath  = cliContext.String("input")
		outputPath = cliContext.String("output")
		reg        = regexp.MustCompile(`(?mi)[\n ]`)
	)

	file, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	if !isJSON(string(file)) {
		return errors.New("The file hasn't json string")
	}

	res := reg.ReplaceAllString(string(file), "")
	if outputPath != "" {
		if err := ioutil.WriteFile(outputPath, []byte(res), 0644); err != nil {
			return err
		}
	} else {
		fmt.Println(res)
	}

	return nil
}
