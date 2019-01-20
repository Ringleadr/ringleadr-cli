package Nodes

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

func ListNodes(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/nodes", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}

	var nodes []Datatypes.Node

	//Bind response to struct
	if err := json.Unmarshal(resp, &nodes); err != nil {
		return Errors.UnexpectedReponse()
	}
	return Format.PrintNodes(&nodes)
}
