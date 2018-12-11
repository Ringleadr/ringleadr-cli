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
	if len(c.Args()) > 0 {
		return ListApplication(c, c.Args())
	}
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

func ListApplication(c *cli.Context, args []string) error {
	appName := args[0]
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/application/%s", Config.GetAgogosHostUrl(), appName))
	if err != nil {
		return err
	}

	if string(resp) == "null" {
		return Format.PrintApplications(nil)
	}

	var app Datatypes.Application

	//Bind response to struct
	if err := json.Unmarshal(resp, &app); err != nil {
		return Errors.UnexpectedReponse()
	}

	var apps = []Datatypes.Application{app}
	return Format.PrintApplications(&apps)
}
