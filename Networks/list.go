package Networks

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

func ListNetworks(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/networks", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}

	var networks []Datatypes.Network

	//Bind response to struct
	if err := json.Unmarshal(resp, &networks); err != nil {
		return Errors.UnexpectedReponse()
	}
	return Format.PrintNetworks(&networks)
}
