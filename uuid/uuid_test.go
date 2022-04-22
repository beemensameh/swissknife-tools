package swissuuid

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUuidVersion(t *testing.T) {
	type testCase struct {
		uuidCli       UUIDCLI
		isFail        bool
		expectedError error
	}

	testCases := map[string]testCase{
		"Should return uuid v2 with persion domain successfully": {
			uuidCli: UUIDCLI{
				Version: 2,
				Domain:  uuid.Person,
				Name:    []byte("name"),
			},
		},
		"Should return uuid v2 with org domain successfully": {
			uuidCli: UUIDCLI{
				Version: 2,
				Domain:  uuid.Org,
				Name:    []byte("name"),
			},
		},
		"Should return uuid v2 with group domain successfully": {
			uuidCli: UUIDCLI{
				Version: 2,
				Domain:  uuid.Group,
				Name:    []byte("name"),
			},
		},
		"Should fail when uuid v6": {
			uuidCli: UUIDCLI{
				Version: 6,
			},
			isFail:        true,
			expectedError: errors.New("version doesn't exist"),
		},
	}

	for i := 0; i < 6; i++ {
		if i != 2 {
			testCases[fmt.Sprintf("Should return uuid v%d", i+1)] = testCase{
				uuidCli: UUIDCLI{
					Version: i,
				},
			}
		}
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			actual, err := uuidVersion(tc.uuidCli.Version, tc.uuidCli.Domain, tc.uuidCli.Name)
			if tc.isFail {
				require.EqualError(t, err, tc.expectedError.Error())
				return
			}

			if tc.uuidCli.Version == 2 {
				require.Equal(t, actual.Domain(), tc.uuidCli.Domain)
			}

			require.Nil(t, err)
			require.IsType(t, uuid.UUID{}, actual)
			require.True(t, actual.String() != "")
			require.True(t, int(actual.Version()) == tc.uuidCli.Version)
		})
	}
}
