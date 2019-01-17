package main

import (
	"github.com/GodlikePenguin/agogos-cli/Application"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Init"
	"github.com/GodlikePenguin/agogos-cli/Networks"
	"github.com/GodlikePenguin/agogos-cli/Storage"
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
				{
					Name:    "storage",
					Aliases: []string{"s"},
					Usage:   "Create a storage volume with the given name",
					Action:  Storage.CreateStorage,
				},
				{
					Name:    "network",
					Aliases: []string{"n"},
					Usage:   "Create a network with the given name",
					Action:  Networks.CreateNetwork,
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
					Aliases: []string{"app", "apps", "application"},
					Usage:   "List all Applications",
					Action:  Application.ListApplications,
				},
				{
					Name:    "storage",
					Aliases: []string{"s"},
					Usage:   "List all storage volumes",
					Action:  Storage.ListStorage,
				},
				{
					Name:    "networks",
					Aliases: []string{"n"},
					Usage:   "List all container networks",
					Action:  Networks.ListNetworks,
				},
			},
		},
		{
			Name:  "dump",
			Usage: "Dump raw resource data",
			Subcommands: []cli.Command{
				{
					Name:    "applications",
					Aliases: []string{"app", "applications"},
					Usage:   "Dump all information about one or all application(s)",
					Action:  Application.DumpApplications,
				},
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a resource",
			Subcommands: []cli.Command{
				{
					Name:    "application",
					Aliases: []string{"app"},
					Usage:   "Delete an application with given name",
					Action:  Application.DeleteApplication,
				},
				{
					Name:    "storage",
					Aliases: []string{"s"},
					Usage:   "Delete the storage volume with the given name",
					Action:  Storage.DeleteStorage,
				},
				{
					Name:    "network",
					Aliases: []string{"n"},
					Usage:   "Delete the network with the given name",
					Action:  Networks.DeleteNetwork,
				},
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Set up an Agogos Host",
			Action:  Init.Init,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
