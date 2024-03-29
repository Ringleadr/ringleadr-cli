package Nodes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Errors"
	"github.com/Ringleadr/ringleadr-cli/internal/Format"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	Datatypes "github.com/Ringleadr/ringleadr-datatypes"
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

	url := fmt.Sprintf("%s/node/%s", Config.GetAgogosHostUrl(), name)

	r := c.Bool("noreschedule")
	if r {
		url += "?noreschedule=true"
	}
	_, err := Requests.DeleteRequest(url)
	if err != nil {
		return err
	}
	return nil
}

func StatsForNode(c *cli.Context) error {
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("node name not specified")
	}

	name := c.Args()[0]
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/node/%s/stats", Config.GetAgogosHostUrl(), name))
	if err != nil {
		return err
	}

	stats := &Datatypes.NodeStatsEntry{}
	err = json.Unmarshal(resp, stats)
	if err != nil {
		return errors.New("Error parsing response from server: " + err.Error())
	}
	return Format.PrintNodeStats(stats)
}
