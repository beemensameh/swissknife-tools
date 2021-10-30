package swisstime

import (
	"errors"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

var TimeNowCmd = &cli.Command{
	Name:    "time:now",
	Usage:   "Get time now and update very i second and with f format",
	Aliases: []string{"time:nw"},
	Action:  timeNowAction,
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

func timeNowAction(cliContext *cli.Context) error {
	if cliContext.Bool("update") {
		if cliContext.Int("interval") < 1 {
			return errors.New("interval should be positive number")
		} else {
			ticker := time.NewTicker(time.Duration(cliContext.Int("interval")) * time.Second)
			for range ticker.C {
				fmt.Print(displayTime(cliContext.Int64("format")))
			}
		}
	} else {
		fmt.Println(displayTime(cliContext.Int64("format")))
	}

	return nil
}

func displayTime(format int64) string {
	const displayStyle = "\r%s"

	switch format {
	case 1:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.ANSIC))
	case 2:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.UnixDate))
	case 3:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RubyDate))
	case 4:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC822))
	case 5:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC822Z))
	case 6:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC850))
	case 7:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC1123))
	case 8:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC1123Z))
	case 9:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC3339))
	case 10:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.RFC3339Nano))
	case 11:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.Kitchen))
	case 12:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.Stamp))
	case 13:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.StampMilli))
	case 14:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.StampMicro))
	case 15:
		return fmt.Sprintf(displayStyle, time.Now().Format(time.StampNano))
	default:
		return fmt.Sprintf(displayStyle, time.Now())
	}
}
