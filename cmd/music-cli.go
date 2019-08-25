package main

import (
	"os"

	"github.com/ryomak/music-cli/cli"
)

func main() {
	app := cli.New("music-cli", "BGM", "1.0.0")
	app.Run(os.Args)
}
