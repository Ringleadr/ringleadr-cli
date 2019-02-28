package main

import (
	"github.com/GodlikePenguin/agogos-cli/Application"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/General"
	"github.com/GodlikePenguin/agogos-cli/Host"
	"github.com/GodlikePenguin/agogos-cli/Init"
	"github.com/GodlikePenguin/agogos-cli/Networks"
	"github.com/GodlikePenguin/agogos-cli/Nodes"
	"github.com/GodlikePenguin/agogos-cli/Storage"
	"github.com/urfave/cli"
	"log"
	"os"
)

var buildTime string

func main() {
	Config.SetupConfig()
	app := cli.NewApp()

	app.Name = "agogos-cli"
	app.Usage = "Command line application to interact with an Agogos host"
	app.Version = buildTime

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
					Name:      "applications",
					Aliases:   []string{"app", "apps", "application"},
					Usage:     "List one or all Applications",
					Action:    Application.ListApplications,
					ArgsUsage: "The name of a specific application (optional)",
				},
				{
					Name:    "storage",
					Aliases: []string{"s"},
					Usage:   "List all storage volumes",
					Action:  Storage.ListStorage,
				},
				{
					Name:    "networks",
					Aliases: []string{"n", "network"},
					Usage:   "List all container networks",
					Action:  Networks.ListNetworks,
				},
				{
					Name:    "nodes",
					Aliases: []string{"node", "no"},
					Usage:   "List all nodes in the Agogos cluster",
					Action:  Nodes.ListNodes,
				},
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update a resource",
			Subcommands: []cli.Command{
				{
					Name:    "application",
					Aliases: []string{"app"},
					Usage:   "Update an existing application",
					Action:  Application.UpdateApplication,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Usage: "Load application config from `FILE` (required)",
						},
					},
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
					Aliases: []string{"app", "applications", "apps"},
					Usage:   "Delete an application with given name",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "all",
							Usage: "Delete all applications",
						},
					},
					Action: Application.DeleteApplication,
				},
				{
					Name:    "storage",
					Aliases: []string{"s"},
					Usage:   "Delete the storage volume with the given name",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "all",
							Usage: "Delete all storage",
						},
					},
					Action: Storage.DeleteStorage,
				},
				{
					Name:    "network",
					Aliases: []string{"n", "networks"},
					Usage:   "Delete the network with the given name",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "all",
							Usage: "Delete all networks",
						},
					},
					Action: Networks.DeleteNetwork,
				},
				{
					Name:    "node",
					Aliases: []string{"no"},
					Usage:   "Delete a node in the cluster (stop scheduling applications on the node)",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "noreschedule",
							Usage: "Don't reschedule Applications scheduled on this node",
						},
					},
					Action: Nodes.DeleteNode,
				},
				{
					Name:   "stats",
					Usage:  "Delete all resource statistics currently stored",
					Action: General.DeleteStats,
				},
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Set up an Agogos Host",
			Action:  Init.Init,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "connect, c",
					Usage: "The address of a primary host to connect to (if initialising a secondary node)",
				},
				cli.BoolFlag{
					Name:  "proxy, p",
					Usage: "Set to start the Agogos Host in proxy mode",
				},
			},
		},
		{
			Name:    "host",
			Aliases: []string{"h", "hosts"},
			Usage:   "Commands related to the host application",
			Subcommands: []cli.Command{
				{
					Name:   "update",
					Usage:  "Update the agogos-host binary on this machine",
					Action: Host.UpdateHost,
				},
				{
					Name:      "stats",
					Aliases:   []string{"s"},
					Usage:     "Display the stats for a given node",
					ArgsUsage: "(name: required)",
					Action:    Nodes.StatsForNode,
				},
			},
		},
		{
			Name:   "purge",
			Usage:  "Delete all resources in the system",
			Action: General.Purge,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
