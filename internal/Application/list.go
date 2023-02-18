package Application

import (
	"encoding/json"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Errors"
	"github.com/Ringleadr/ringleadr-cli/internal/Format"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	Datatypes "github.com/Ringleadr/ringleadr-datatypes"
	"github.com/urfave/cli"
)

func ListApplications(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return ListApplication(c, c.Args())
	}
	apps, err := getApplications()
	if err != nil {
		return err
	}
	return Format.PrintApplications(&apps)
}

func ListApplication(c *cli.Context, args []string) error {
	appName := args[0]
	app, err := getApplication(appName)
	if err != nil {
		return err
	}

	var apps []Datatypes.Application
	if app != nil {
		apps = append(apps, *app)
	}
	return Format.PrintApplications(&apps)
}

func getApplications() ([]Datatypes.Application, error) {
	bytes, err := getApplicationsBytes()
	if err != nil {
		return nil, err
	}
	var apps []Datatypes.Application

	//Bind response to struct
	if err := json.Unmarshal(bytes, &apps); err != nil {
		return nil, Errors.UnexpectedReponse()
	}
	return apps, nil
}

func getApplication(name string) (*Datatypes.Application, error) {
	bytes, err := getApplicationBytes(name)
	if err != nil {
		return nil, err
	}

	if string(bytes) == "null" {
		return nil, nil
	}

	var app Datatypes.Application

	//Bind response to struct
	if err := json.Unmarshal(bytes, &app); err != nil {
		return nil, Errors.UnexpectedReponse()
	}
	return &app, nil
}

func getApplicationsBytes() ([]byte, error) {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/applications", Config.GetAgogosHostUrl()))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getApplicationBytes(name string) ([]byte, error) {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/application/%s", Config.GetAgogosHostUrl(), name))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
