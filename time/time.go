package swisstime

import (
	"errors"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

const displayStyle = "\r%s"

var TimeNowCmd = &cli.Command{
	Name:    "time:now",
	Usage:   "Get time now and update very i second and with f format",
	Aliases: []string{"time:nw"},
	Action:  TimeNowAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "format",
			Usage:   "Time format (the value should be int)",
			Aliases: []string{"f"},
		},
		&cli.BoolFlag{
			Name:    "update",
			Usage:   "Is time should continuously update",
			Value:   false,
			Aliases: []string{"up"},
		},
		&cli.IntFlag{
			Name:    "interval",
			Usage:   "Interval of update the display time (Work with `update` tag and should be positive number)",
			Value:   1,
			Aliases: []string{"i"},
		},
	},
}

func TimeNowAction(cliContext *cli.Context) error {
	return timeNow(&TimeCLI{
		Format:   newTimeFormat(cliContext.Int("format")),
		Update:   cliContext.Bool("update"),
		Interval: cliContext.Int("interval"),
	})
}

func timeNow(timeCLI *TimeCLI) error {
	if err := timeCLI.validated(); err != nil {
		return err
	}

	if timeCLI.Update {
		if timeCLI.Interval < 1 {
			return errors.New("interval should be positive number")
		} else {
			ticker := time.NewTicker(time.Duration(timeCLI.Interval) * time.Second)
			for range ticker.C {
				fmt.Print(displayTime(time.Now(), timeCLI.Format))
			}
		}
	} else {
		fmt.Println(displayTime(time.Now(), timeCLI.Format))
	}

	return nil
}

func displayTime(timeNow time.Time, format string) string {
	if format == "" {
		return fmt.Sprintf(displayStyle, timeNow.String())
	}

	return fmt.Sprintf(displayStyle, timeNow.Format(format))
}
