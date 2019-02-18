package Application

import (
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
)

func DeleteApplication(c *cli.Context) error {
	if c.Bool("all") {
		return DeleteAllApplications(c)
	}
	//TODO handle multiple application names
	appName := c.Args().Get(0)
	if appName == "" {
		cli.ShowSubcommandHelp(c)
		return errors.New("delete requires an application name to delete")
	}
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/applications/%s", Config.GetAgogosHostUrl(), appName))
	return err
}

func DeleteAllApplications(c *cli.Context) error {
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/all/applications", Config.GetAgogosHostUrl()))
	return err
}
