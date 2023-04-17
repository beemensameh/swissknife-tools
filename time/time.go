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

	TimeCmd = &cobra.Command{
		Use:   "time",
		Short: "Use time tool",
	}
)

func init() {
	nowCmd := &cobra.Command{
		Use:   "now",
		Short: "Get time now",
		Long:  "Get time now and update every i second with f format",
		Run:   timeNowAction,
	}
	nowCmd.Flags().UintVarP(&format, "format", "f", 0, "time format (see ./docs/time.md)")
	nowCmd.Flags().BoolVarP(&update, "update", "u", false, "refresh the time depend on interval")
	nowCmd.Flags().UintVarP(&interval, "interval", "i", 1, "update the display time every i sec")
	TimeCmd.AddCommand(nowCmd)
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
