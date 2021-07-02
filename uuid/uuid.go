package swissuuid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

type UUIDType string

const (
	DCEGroup    UUIDType = "group"
	DCEPersion  UUIDType = "persion"
	DCESecurity UUIDType = "security"
)

type UUIDDCESecurity string

const (
	Person UUIDDCESecurity = "persion"
	Group  UUIDDCESecurity = "group"
	Org    UUIDDCESecurity = "org"
)

var GenerateUUIDCmd = &cli.Command{
	Name:   "uuid:generate",
	Usage:  "Generate UUID for any UUID version",
	Action: GenerateUUIDAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "version",
			Usage:   "The uuid version",
			Value:   4,
			Aliases: []string{"v"},
		},
		&cli.StringFlag{
			Name:    "uuid-type",
			Usage:   "It is DCE uuid types (should be one of [group, persion, security] and should add value when uuid version is 2)",
			Value:   "group",
			Aliases: []string{"uid-type"},
		},
		&cli.StringFlag{
			Name:    "uuid-security-type",
			Usage:   "It is DCE security uuid types (should be one of [group, persion, org] and should add value when uuid version is 2 and the uuid type is security)",
			Value:   "group",
			Aliases: []string{"uid-sec-type"},
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "It is used for uuid version 3 and 5 (maybe anything - no constrained)",
			Aliases: []string{"n"},
		},
		&cli.Int64Flag{
			Name:    "number",
			Usage:   "Number of UUID need to generate",
			Value:   1,
			Aliases: []string{"num"},
		},
		&cli.StringFlag{
			Name:    "seperated",
			Usage:   "The seperated character that should seperat UUIDs",
			Aliases: []string{"sep"},
		},
	},
}

func GenerateUUIDAction(cliContext *cli.Context) error {
	var allUUIDs []string

	for range make([]int, cliContext.Int64("number")) {
		uuid, err := uuidVersion(
			cliContext.Int64("version"),
			UUIDType(cliContext.String("uuid-type")),
			UUIDDCESecurity(cliContext.String("uuid-security-type")),
			[]byte(cliContext.String("name")),
		)
		if err != nil {
			return err
		}

		if cliContext.String("seperated") == "" {
			fmt.Println(uuid)
		} else {
			allUUIDs = append(allUUIDs, uuid.String())
		}
	}

	if len(allUUIDs) > 0 {
		fmt.Println(strings.Join(allUUIDs, ","))
	}

	return nil
}

func uuidVersion(version int64, uuidType UUIDType, uuidDCESecurity UUIDDCESecurity, name []byte) (uuid.UUID, error) {
	switch version {
	case 1:
		return uuid.NewUUID()
	case 2:
		switch uuidType {
		case DCEGroup:
			return uuid.NewDCEGroup()
		case DCEPersion:
			return uuid.NewDCEPerson()
		default:
			return uuid.UUID{}, errors.New("DCESecurity doesn't support yet")
		}
	case 3:
		return uuid.NewMD5(uuid.New(), name), nil
	case 4:
		return uuid.NewRandom()
	case 5:
		return uuid.NewSHA1(uuid.New(), name), nil
	default:
		return uuid.UUID{}, errors.New("Version doesn't found")
	}
}
