package Storage

import (
	"encoding/json"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-cli/Format"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/GodlikePenguin/agogos-datatypes"
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
