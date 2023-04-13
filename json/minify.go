package swissjson

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	input, output string

	JSONMinifyCmd = &cobra.Command{
		Use:     "json:minify",
		Short:   "Minify json",
		Aliases: []string{"json:min"},
		RunE:    jsonMinifyAction,
	}
)

func init() {
	JSONMinifyCmd.Flags().StringVarP(&input, "input", "i", "", "The path for the input file")
	JSONMinifyCmd.Flags().StringVarP(&output, "output", "o", "", "The path for the output file (if this flag doesn't add will print it in terminal)")
	if err := JSONMinifyCmd.MarkFlagRequired("input"); err != nil {
		log.Fatal(err)
	}
}

func jsonMinifyAction(cmd *cobra.Command, args []string) error {
	return jsonMinify(&JSONValidation{
		Input:  input,
		Output: output,
	})
}

func jsonMinify(j *JSONValidation) error {
	err := j.validated()
	if err != nil {
		return err
	}

	j.Minify()

	if j.Output != "" {
		if err := os.WriteFile(j.Output, []byte(j.JSON), 0o600); err != nil {
			return err
		}
	} else {
		fmt.Println(j.JSON)
	}

	return nil
}
