package Storage

import (
	"encoding/json"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Errors"
	"github.com/Ringleadr/ringleadr-cli/internal/Format"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	Datatypes "github.com/Ringleadr/ringleadr-datatypes"
	"github.com/urfave/cli"
)

func ListStorage(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/storage", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}

	var storage []Datatypes.Storage

	//Bind response to struct
	if err := json.Unmarshal(resp, &storage); err != nil {
		return Errors.UnexpectedReponse()
	}
	return Format.PrintStorage(&storage)
}
