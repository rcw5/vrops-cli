package commands

import (
	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
	"github.com/urfave/cli"
)

func GetResourceKindsCommand(client clients.VRopsClientIntf) cli.Command {
	command := cli.Command{
		Name:  "resourcekinds",
		Usage: "get all resourcekinds for an adapter",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "output, o", Value: "table"},
		},
		Action: func(c *cli.Context) error {
			presenter := presenters.NewPresenter(c.String("output"))
			GetResourceKinds(c.Args().First(), client, presenter)
			return nil
		},
	}
	return command
}
func GetResourceKinds(adapterKind string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	resourceKinds, err := client.ResourceKinds(adapterKind)
	if err != nil {
		return err
	}

	presenter.PresentResourceKinds(resourceKinds)
	return nil
}
