package main

import (
	"os"

	"github.com/ryomak/music-cli/mcli"
)

func main() {
  app := mcli.New("music-cli", "BGM", "1.0.0")
	app.Run(os.Args)
}
