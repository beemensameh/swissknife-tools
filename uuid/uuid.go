package swissuuid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var GenerateUUIDCmd = &cli.Command{
	Name:    "uuid:generate",
	Usage:   "Generate UUID for any UUID version",
	Action:  GenerateUUIDAction,
	Aliases: []string{"uuid:gen"},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "version",
			Usage:   "The uuid version",
			Value:   4,
			Aliases: []string{"v"},
		},
		&cli.StringFlag{
			Name:    "uuid-type",
			Usage:   "It is DCE uuid types (should be one of [group, person, security] and should add value when uuid version is 2)",
			Value:   "group",
			Aliases: []string{"uid-type"},
		},
		&cli.StringFlag{
			Name:    "uuid-security-type",
			Usage:   "It is DCE security uuid types (should be one of [group, person, org] and should add value when uuid version is 2 and the uuid type is security)",
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
			Name:    "separated",
			Usage:   "The separated character that should separate UUIDs",
			Aliases: []string{"sep"},
		},
	},
}

func GenerateUUIDAction(cliContext *cli.Context) error {
	return generateUUID(&UUIDCLI{
		Version:          cliContext.Int("version"),
		UUIDType:         UUIDTypeEnum(cliContext.String("uuid-type")),
		UUIDSecurityType: UUIDDCESecurityEnum(cliContext.String("uuid-security-type")),
		Name:             []byte(cliContext.String("name")),
		Number:           cliContext.Int64("number"),
		Separate:         cliContext.String("separated"),
	})
}

func generateUUID(uuidCLI *UUIDCLI) error {
	var allUUIDs []string

	err := uuidCLI.validated()
	if err != nil {
		return err
	}

	for range make([]int, uuidCLI.Number) {
		uuid, err := uuidVersion(
			uuidCLI.Version,
			uuidCLI.UUIDType,
			uuidCLI.UUIDSecurityType,
			uuidCLI.Name,
		)
		if err != nil {
			return err
		}

		if uuidCLI.Separate == "" {
			fmt.Println(uuid)
		} else {
			allUUIDs = append(allUUIDs, uuid.String())
		}
	}

	if len(allUUIDs) > 0 {
		fmt.Println(strings.Join(allUUIDs, uuidCLI.Separate))
	}

	return nil
}

func uuidVersion(version int, uuidType UUIDTypeEnum, uuidDCESecurity UUIDDCESecurityEnum, name []byte) (uuid.UUID, error) {
	switch version {
	case 0:
		return uuid.Nil, nil
	case 1:
		return uuid.NewUUID()
	case 2:
		switch uuidType {
		case DCEGroup:
			return uuid.NewDCEGroup()
		case DCEPerson:
			return uuid.NewDCEPerson()
		case DCESecurity:
			return uuid.UUID{}, errors.New("DCESecurity doesn't support yet")
		default:
			return uuid.UUID{}, errors.New("UUID type doesn't exist")
		}
	case 3:
		return uuid.NewMD5(uuid.New(), name), nil
	case 4:
		return uuid.NewRandom()
	case 5:
		return uuid.NewSHA1(uuid.New(), name), nil
	default:
		return uuid.UUID{}, errors.New("version doesn't exist")
	}
}
