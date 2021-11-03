package swisstime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// func TestTimeNow(t *testing.T) {
// 	testCases := []struct {
// 		desc    string
// 		timeCLI TimeCLI
// 		isFail  bool
// 	}{}
// }

func TestDisplayTime(t *testing.T) {
	var testTime time.Time = time.Now()

	testCases := []struct {
		desc         string
		timeCLI      TimeCLI
		expectedTime string
	}{
		{
			desc: "Should return time now without format, when pass wrong format number",
			timeCLI: TimeCLI{
				Format: newTimeFormat(0),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.String()),
		},
		{
			desc: "Should return time with ANSIC format, when pass format number 1",
			timeCLI: TimeCLI{
				Format: newTimeFormat(1),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.ANSIC)),
		},
		{
			desc: "Should return time with UnixDate format, when pass format number 2",
			timeCLI: TimeCLI{
				Format: newTimeFormat(2),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.UnixDate)),
		},
		{
			desc: "Should return time with RubyDate format, when pass format number 3",
			timeCLI: TimeCLI{
				Format: newTimeFormat(3),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RubyDate)),
		},
		{
			desc: "Should return time with RFC822 format, when pass format number 4",
			timeCLI: TimeCLI{
				Format: newTimeFormat(4),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC822)),
		},
		{
			desc: "Should return time with RFC822Z format, when pass format number 5",
			timeCLI: TimeCLI{
				Format: newTimeFormat(5),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC822Z)),
		},
		{
			desc: "Should return time with RFC850 format, when pass format number 6",
			timeCLI: TimeCLI{
				Format: newTimeFormat(6),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC850)),
		},
		{
			desc: "Should return time with RFC1123 format, when pass format number 7",
			timeCLI: TimeCLI{
				Format: newTimeFormat(7),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC1123)),
		},
		{
			desc: "Should return time with RFC1123Z format, when pass format number 8",
			timeCLI: TimeCLI{
				Format: newTimeFormat(8),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC1123Z)),
		},
		{
			desc: "Should return time with RFC3339 format, when pass format number 9",
			timeCLI: TimeCLI{
				Format: newTimeFormat(9),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC3339)),
		},
		{
			desc: "Should return time with RFC3339Nano format, when pass format number 10",
			timeCLI: TimeCLI{
				Format: newTimeFormat(10),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.RFC3339Nano)),
		},
		{
			desc: "Should return time with Kitchen format, when pass format number 11",
			timeCLI: TimeCLI{
				Format: newTimeFormat(11),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.Kitchen)),
		},
		{
			desc: "Should return time with Stamp format, when pass format number 12",
			timeCLI: TimeCLI{
				Format: newTimeFormat(12),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.Stamp)),
		},
		{
			desc: "Should return time with StampMilli format, when pass format number 13",
			timeCLI: TimeCLI{
				Format: newTimeFormat(13),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.StampMilli)),
		},
		{
			desc: "Should return time with StampMicro format, when pass format number 14",
			timeCLI: TimeCLI{
				Format: newTimeFormat(14),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.StampMicro)),
		},
		{
			desc: "Should return time with StampNano format, when pass format number 15",
			timeCLI: TimeCLI{
				Format: newTimeFormat(15),
			},
			expectedTime: fmt.Sprintf(displayStyle, testTime.Format(time.StampNano)),
		},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			actualTime := displayTime(testTime, tC.timeCLI.Format)

			require.Equal(t, tC.expectedTime, actualTime)
		})
	}
}
