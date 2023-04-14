package swissuuid

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	name, sep string
	v, ust    int
	num       int64

	GenerateUUIDCmd = &cobra.Command{
		Use:     "uuid:generate",
		Short:   "Generate UUID for any UUID version",
		Long:    "Generate UUID for any UUID version with many other configuration depend on the UUID type",
		Aliases: []string{"uuid:gen"},
		RunE:    generateUUIDAction,
	}
)

func init() {
	GenerateUUIDCmd.Flags().IntVarP(&v, "version", "v", 4, "The uuid version")
	GenerateUUIDCmd.Flags().IntVarP(&ust, "uuid-security-type", "u", 0, "It is DCE security uuid types (should be one of [0, 1, 2] and should add value when uuid version is 2)")
	GenerateUUIDCmd.Flags().StringVarP(&name, "name", "", "", "It is used for uuid version 3 and 5 (maybe anything - no constrained)")
	GenerateUUIDCmd.Flags().Int64VarP(&num, "number", "n", 1, "Number of UUID need to generate")
	GenerateUUIDCmd.Flags().StringVarP(&sep, "separated", "s", "", "The separated character that should separate UUIDs")
}

func generateUUIDAction(cmd *cobra.Command, args []string) error {
	return generateUUID(&UUIDCLI{
		Version:  v,
		Domain:   uuid.Domain(ust),
		Name:     []byte(name),
		Number:   num,
		Separate: sep,
	})
}

func generateUUID(uidCLI *UUIDCLI) error {
	var allUUIDs []string

	err := uidCLI.validated()
	if err != nil {
		return err
	}

	for range make([]int, uidCLI.Number) {
		uuid, err := uuidVersion(
			uidCLI.Version,
			uidCLI.Domain,
			uidCLI.Name,
		)
		if err != nil {
			return err
		}

		if uidCLI.Separate == "" {
			fmt.Println(uuid)
		} else {
			allUUIDs = append(allUUIDs, uuid.String())
		}
	}

	if len(allUUIDs) > 0 {
		fmt.Println(strings.Join(allUUIDs, uidCLI.Separate))
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
