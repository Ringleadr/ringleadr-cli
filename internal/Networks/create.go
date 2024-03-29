package Networks

import (
	"errors"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	"github.com/urfave/cli"
)

func CreateNetwork(c *cli.Context) error {
	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return errors.New("network name not specified")
	}

	name := c.Args()[0]
	_, err := Requests.PostRequest(fmt.Sprintf("%s/networks/%s", Config.GetAgogosHostUrl(), name), nil)
	if err != nil {
		return err
	}
	return nil
}
