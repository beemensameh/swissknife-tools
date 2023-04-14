package swissjson

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	input, output string

	JSONCmd = &cobra.Command{
		Use:   "json",
		Short: "Use json tool",
	}
)

func init() {
	minifyCmd := &cobra.Command{
		Use:   "minify",
		Short: "Minify json",
		RunE:  jsonMinifyAction,
	}
	minifyCmd.Flags().StringVarP(&input, "input", "i", "", "The path for the input file")
	minifyCmd.Flags().StringVarP(&output, "output", "o", "", "The path for the output file (if this flag doesn't add will print it in terminal)")
	if err := minifyCmd.MarkFlagRequired("input"); err != nil {
		log.Fatal(err)
	}
	if err := minifyCmd.MarkFlagFilename("input", "txt", "json"); err != nil {
		log.Fatal(err)
	}
	JSONCmd.AddCommand(minifyCmd)
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
