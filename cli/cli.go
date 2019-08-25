package cli

import (
	"github.com/urfave/cli"
)

func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = getCommands()
	return app
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "play",
			Usage:  "play music",
			Action: play,
		},
	}
}
