package Storage

import (
	"errors"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	"github.com/urfave/cli"
)

func DeleteStorage(c *cli.Context) error {
	if c.Bool("all") {
		return DeleteAllStorage(c)
	}
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("storage name not specified")
	}

	name := c.Args()[0]
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/storage/%s", Config.GetAgogosHostUrl(), name))
	return err
}

func DeleteAllStorage(c *cli.Context) error {
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/all/storage", Config.GetAgogosHostUrl()))
	return err
}
