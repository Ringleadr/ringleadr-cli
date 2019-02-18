package Networks

import (
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
)

func DeleteNetwork(c *cli.Context) error {
	if c.Bool("all") {
		return DeleteAllNetworks(c)
	}
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("network name not specified")
	}

	name := c.Args()[0]
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/networks/%s", Config.GetAgogosHostUrl(), name))
	return err
}

func DeleteAllNetworks(c *cli.Context) error {
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/all/networks", Config.GetAgogosHostUrl()))
	return err
}
