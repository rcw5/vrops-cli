package main

import (
	"os"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/commands"
	"github.com/rcw5/vrops-cli/presenters"
	"github.com/urfave/cli"
)

func main() {
	var client clients.VRopsClient

	app := cli.NewApp()
	app.Name = "vrops-cli"
	app.Usage = "CLI to interact with VMware vRealize Operations Manager"
	app.Version = "0.0.1"
	app.Before = func(c *cli.Context) error {
		client = clients.NewVROpsClient(c.String("url"), c.String("username"), c.String("password"))
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "vRops Username",
			EnvVar: "VROPS_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "vRops Password",
			EnvVar: "VROPS_PASSWORD",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "vRops URL",
			EnvVar: "VROPS_URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Retrieve data from vROps",
			Subcommands: []cli.Command{
				cli.Command{
					Name:  "adapterkinds",
					Usage: "get all adapterkinds",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "output, o", Value: "table"},
					},
					Action: func(c *cli.Context) error {
						presenter := presenters.NewPresenter(c.String("output"))
						commands.GetAdapterKinds(client, presenter)
						return nil
					},
				},
				cli.Command{
					Name:  "resourcekinds",
					Usage: "get all resourcekinds for an adapter",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "output, o", Value: "table"},
					},
					Action: func(c *cli.Context) error {
						presenter := presenters.NewPresenter(c.String("output"))
						commands.GetResourceKinds(c.Args().First(), client, presenter)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)

}
