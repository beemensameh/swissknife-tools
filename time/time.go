package swisstime

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const displayStyle = "\r%s"

var (
	format, interval uint
	update           bool

	TimeNowCmd = &cobra.Command{
		Use:     "time:now",
		Short:   "Get time now",
		Long:    "Get time now and update every i second with f format",
		Aliases: []string{"time:nw"},
		Run:     timeNowAction,
	}
)

func init() {
	TimeNowCmd.Flags().UintVarP(&format, "format", "f", 0, "time format (see ./docs/time.md)")
	TimeNowCmd.Flags().BoolVarP(&update, "update", "u", false, "refresh the time depend on interval")
	TimeNowCmd.Flags().UintVarP(&interval, "interval", "i", 1, "update the display time every i sec")
}

func timeNowAction(cmd *cobra.Command, args []string) {
	timeNow(&TimeCLI{
		Format:   newTimeFormat(format),
		Update:   update,
		Interval: interval,
	})
}

func timeNow(t *TimeCLI) {
	t.validated()
	if t.Update {
		ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
		for range ticker.C {
			fmt.Print(displayTime(time.Now(), t.Format))
		}
	} else {
		fmt.Println(displayTime(time.Now(), t.Format))
	}
}

func displayTime(tNow time.Time, format string) string {
	if format == "" {
		return fmt.Sprintf(displayStyle, tNow.String())
	}

	return fmt.Sprintf(displayStyle, tNow.Format(format))
}
