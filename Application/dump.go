package Application

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)

func DumpApplications(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return DumpApplication(c, c.Args())
	}
	apps, err := getApplications()
	if err != nil {
		return err
	}
	appsString, err := json.MarshalIndent(apps, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(appsString))
	return nil
}

func DumpApplication(c *cli.Context, args []string) error {
	appName := args[0]
	app, err := getApplication(appName)
	if err != nil {
		return err
	}
	appString, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(appString))
	return nil
}
