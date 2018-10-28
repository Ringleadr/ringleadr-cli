package Application

import (
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
)

func ListApplications(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/applications", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}
	println(string(resp))
	return nil
}
