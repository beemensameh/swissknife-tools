package main

import (
	"log"

	swisstime "github.com/beemensameh/swissknife-tools/time"
	"github.com/spf13/cobra"
)

var (
	name    string
	version string
)

func main() {
	rootCmd := &cobra.Command{
		Use:     name,
		Short:   "A CLI application for many tools",
		Long:    "A quick and amazing tools for speed up your work",
		Version: version,
	}
	rootCmd.AddCommand(swisstime.TimeNowCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
