package Storage

import (
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
)

func CreateStorage(c *cli.Context) error {
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("storage name not specified")
	}

	name := c.Args()[0]
	_, err := Requests.PostRequest(fmt.Sprintf("%s/storage/%s", Config.GetAgogosHostUrl(), name), nil)
	if err != nil {
		return err
	}
	return nil
}
