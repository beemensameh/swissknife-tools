package swisstime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeNow(t *testing.T) {
	testCases := map[string]struct {
		timeCLI TimeCLI
	}{
		"Should return time now successfully": {
			timeCLI: TimeCLI{
				Format: newTimeFormat(0),
			},
		},
		"Should return time now even the interval is negative or zero": {
			timeCLI: TimeCLI{
				Format:   newTimeFormat(1),
				Interval: 0,
			},
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			err := timeNow(&tc.timeCLI)
			require.Nil(t, err)
		})
	}
}

func TestDisplayTime(t *testing.T) {
	testTime := time.Now()
	type testCase struct {
		timeCLI        TimeCLI
		expectedResult string
	}

	testCases := map[string]testCase{
		"Should return time now without format, when pass wrong format number": {
			timeCLI: TimeCLI{
				Format: newTimeFormat(0),
			},
			expectedResult: fmt.Sprintf(displayStyle, testTime.String()),
		},
	}

	for i, v := range timeFormat {
		testCases[fmt.Sprintf("Should return time with %s format, when pass format number %d", v, i)] = testCase{
			timeCLI: TimeCLI{
				Format: newTimeFormat(i),
			},
			expectedResult: fmt.Sprintf(displayStyle, testTime.Format(timeFormat[i])),
		}
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			actualResult := displayTime(testTime, tc.timeCLI.Format)

			require.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
