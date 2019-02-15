package Nodes

import (
	"encoding/json"
	"errors"
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

func DeleteNode(c *cli.Context) error {
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("node name not specified")
	}

	name := c.Args()[0]
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/node/%s", Config.GetAgogosHostUrl(), name))
	if err != nil {
		return err
	}
	return nil
}
