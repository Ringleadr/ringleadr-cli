package Application

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

func ListApplications(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/applications", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}
	var apps []Datatypes.Application

	//Bind response to struct
	if err := json.Unmarshal(resp, &apps); err != nil {
		return Errors.UnexpectedReponse()
	}
	return Format.PrintApplications(&apps)
}
