package main

import (
	"fmt"
  "os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "play music"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:   "play",
			Usage:  "play music",
			Action: play,
		},
	}

	app.Run(os.Args)
}

func play(c *cli.Context) {
	for _, v := range c.Args() {
		fmt.Println(v)
	}
}
