package swissuuid

import (
	"errors"
	"fmt"
	"os"
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
		&cli.IntFlag{
			Name:    "uuid-security-type",
			Usage:   "It is DCE security uuid types (should be one of [0, 1, 2] and should add value when uuid version is 2)",
			Value:   0,
			Aliases: []string{"uid-sec-type"},
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "It is used for uuid version 3 and 5 (maybe anything - no constrained)",
		},
		&cli.Int64Flag{
			Name:    "number",
			Usage:   "Number of UUID need to generate",
			Value:   1,
			Aliases: []string{"n"},
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
		Version:  cliContext.Int("version"),
		Domain:   uuid.Domain(cliContext.Int("uuid-security-type")),
		Name:     []byte(cliContext.String("name")),
		Number:   cliContext.Int64("number"),
		Separate: cliContext.String("separated"),
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
			uuidCLI.Domain,
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

func uuidVersion(version int, domain uuid.Domain, name []byte) (uuid.UUID, error) {
	switch version {
	case 0:
		return uuid.Nil, nil
	case 1:
		return uuid.NewUUID()
	case 2:
		return uuid.NewDCESecurity(domain, uint32(os.Getuid()))
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
