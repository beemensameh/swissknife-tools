package swisstime

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeNow(t *testing.T) {
	testCases := []struct {
		desc          string
		timeCLI       TimeCLI
		isFail        bool
		expectedError error
	}{
		{
			desc: "Should return time now successfully",
			timeCLI: TimeCLI{
				Format: newTimeFormat(0),
			},
		},
		{
			desc: "Should return validation error",
			timeCLI: TimeCLI{
				Format:   newTimeFormat(1),
				Interval: -5,
			},
			isFail:        true,
			expectedError: errors.New("Interval should be gte (1)"),
		},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			err := timeNow(&tC.timeCLI)

			if tC.isFail {
				require.EqualError(t, err, tC.expectedError.Error())
				return
			}

			require.Nil(t, err)
		})
	}
}

func TestDisplayTime(t *testing.T) {
	var testTime time.Time = time.Now()
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

	for i, v := range timeFromat {
		testCases[fmt.Sprintf("Should return time with %s format, when pass format number %d", v, i)] = testCase{
			timeCLI: TimeCLI{
				Format: newTimeFormat(i),
			},
			expectedResult: fmt.Sprintf(displayStyle, testTime.Format(timeFromat[i])),
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
