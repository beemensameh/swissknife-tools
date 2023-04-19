package swisstime

import (
	"errors"
	"time"

	"github.com/beemensameh/swissknife-tools/internal/color"
)

var timeFormat = map[uint]string{
	1:  time.ANSIC,
	2:  time.UnixDate,
	3:  time.RubyDate,
	4:  time.RFC822,
	5:  time.RFC822Z,
	6:  time.RFC850,
	7:  time.RFC1123,
	8:  time.RFC1123Z,
	9:  time.RFC3339,
	10: time.RFC3339Nano,
	11: time.Kitchen,
	12: time.Stamp,
	13: time.StampMilli,
	14: time.StampMicro,
	15: time.StampNano,
}

func newTimeFormat(format uint) string {
	return timeFormat[format]
}

type TimeCLI struct {
	Format   string
	Update   bool
	Interval uint
	Zone     string
	Loc      *time.Location
}

func (t *TimeCLI) validated() error {
	if t.Interval <= 0 && t.Update {
		color.PrintlnColor("Interval should be large that 0. Change interval to 1.", color.Yellow)
		t.Interval = 1
	}
	loc := time.Local
	if t.Zone != "" {
		var err error
		loc, err = time.LoadLocation(t.Zone)
		if err != nil {
			return errors.New(color.SprintfColor(err.Error(), color.Red))
		}
	}
	t.Loc = loc
	return nil
}
