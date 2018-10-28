package main

import (
	"github.com/GodlikePenguin/agogos-cli/Application"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	Config.SetupConfig()
	app := cli.NewApp()

	app.Name = "agogos-cli"
	app.Usage = "Command line application to interact with an Agogos host"

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a resource",
			Subcommands: []cli.Command{
				{
					Name:    "application",
					Aliases: []string{"app"},
					Usage:   "Create an application",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Usage: "Load application config from `FILE` (required)",
						},
					},
					Action: Application.CreateApplication,
				},
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List resources",
			Subcommands: []cli.Command{
				{
					Name:    "applications",
					Aliases: []string{"app", "application"},
					Usage:   "List all Applications",
					Action:  Application.ListApplications,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
