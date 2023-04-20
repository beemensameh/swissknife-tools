package swisstime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidated(t *testing.T) {
	testCases := map[string]struct {
		timeCLI           TimeCLI
		isErr             bool
		isIntervalChanged bool
	}{
		"Should success when no args passed": {
			timeCLI: TimeCLI{},
		},
		"Should success when pass format": {
			timeCLI: TimeCLI{
				Format: newTimeFormat(1),
			},
		},
		"Should success when pass update with no interval": {
			timeCLI: TimeCLI{
				Update: true,
			},
			isIntervalChanged: true,
		},
		"Should success when pass update with interval": {
			timeCLI: TimeCLI{
				Update:   true,
				Interval: 1,
			},
		},
		"Should success when pass correct timezone": {
			timeCLI: TimeCLI{
				Zone: "Africa/Cairo",
			},
		},
		"Should fail when pass wrong timezone": {
			timeCLI: TimeCLI{
				Zone: "test",
			},
			isErr: true,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			t.Parallel()
			err := tc.timeCLI.validated()
			if tc.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if tc.isIntervalChanged {
				require.Equal(t, uint(1), tc.timeCLI.Interval)
			}
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
