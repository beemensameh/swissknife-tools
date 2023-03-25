package main

import (
	"log"
	"os"

	swisshashing "github.com/beemensameh/swissknife-tools/hashing"
	swissjson "github.com/beemensameh/swissknife-tools/json"
	swisstime "github.com/beemensameh/swissknife-tools/time"
	swissuuid "github.com/beemensameh/swissknife-tools/uuid"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "swissknife-tools Tools",
		Usage:       "A CLI application for many tools",
		Description: "A quick and amazing tools for speed up your work",
		Commands: []*cli.Command{
			swissjson.JSONMinifyCmd,
			swisstime.TimeNowCmd,
			swissuuid.GenerateUUIDCmd,
			swisshashing.HashFileCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Error in app.Run:\n", err)
	}
}
