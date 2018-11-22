package Application

import (
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
)

func DeleteApplication(c *cli.Context) error {
	//TODO handle multiple application names
	appName := c.Args().Get(0)
	if appName == "" {
		cli.ShowSubcommandHelp(c)
		return errors.New("delete requires an application name to delete")
	}
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/applications/%s", Config.GetAgogosHostUrl(), appName))
	if err != nil {
		return err
	}
	return nil
}
