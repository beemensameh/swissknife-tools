package main

import (
	"log"

	swisshashing "github.com/beemensameh/swissknife-tools/hashing"
	swissjson "github.com/beemensameh/swissknife-tools/json"
	swisstime "github.com/beemensameh/swissknife-tools/time"
	swissuuid "github.com/beemensameh/swissknife-tools/uuid"
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
		Long:    "A quick and amazing CLI application with many tools for speed up your work",
		Version: version,
	}
	rootCmd.AddCommand(swisstime.TimeCmd)
	rootCmd.AddCommand(swissjson.JSONCmd)
	rootCmd.AddCommand(swisshashing.HashCmd)
	rootCmd.AddCommand(swissuuid.UUIDCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
