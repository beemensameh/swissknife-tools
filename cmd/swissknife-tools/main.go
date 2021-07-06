package main

import (
	"log"
	"os"

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
			swisstime.TimeNowCmd,
			swissuuid.GenerateUUIDCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Error in app.Run:\n", err)
	}
}
