package main

import (
	"fmt"
	"os"
	"strings"

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
		client = clients.NewVROpsClient(c.String("url"), c.String("username"), c.String("password"), c.Bool("verbose"))
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
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose requests",
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
						err := commands.GetAdapterKinds(client, presenter)
						if err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
				cli.Command{
					Name:  "resourcekinds",
					Usage: "get all resourcekinds for an adapterkind",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "output, o", Value: "table"},
					},
					Action: func(c *cli.Context) error {
						if len(c.Args()) != 1 {
							fmt.Println("Error with arguments:", c.Args(), "\n")
							tmpl := cli.CommandHelpTemplate
							cli.CommandHelpTemplate = strings.Replace(cli.CommandHelpTemplate, "[arguments...]", "<adapterkind>", -1)
							cli.ShowCommandHelp(c, "resourcekinds")
							cli.CommandHelpTemplate = tmpl
							return nil
						}
						presenter := presenters.NewPresenter(c.String("output"))
						err := commands.GetResourceKinds(c.Args().First(), client, presenter)
						if err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
				cli.Command{
					Name:  "resources",
					Usage: "get all resources for an adapterkind",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "output, o", Value: "table"},
					},
					Action: func(c *cli.Context) error {
						if len(c.Args()) != 1 {
							fmt.Println("Error with arguments:", c.Args(), "\n")
							tmpl := cli.CommandHelpTemplate
							cli.CommandHelpTemplate = strings.Replace(cli.CommandHelpTemplate, "[arguments...]", "<adapterkind>", -1)
							cli.ShowCommandHelp(c, "resources")
							cli.CommandHelpTemplate = tmpl
							return nil
						}
						presenter := presenters.NewPresenter(c.String("output"))
						err := commands.GetResources(c.Args().First(), client, presenter)
						if err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)

}
