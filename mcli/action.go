package mcli

import (
	"fmt"
	"path/filepath"

	"github.com/ryomak/musicbox/sound"
	"github.com/urfave/cli"
)

func play(c *cli.Context) {
	for _, v := range c.Args() {
		_, filename := filepath.Split(v)
		fmt.Printf("start %s... \n", filename)
		err := sound.Sound(v)
		if err != nil {
			fmt.Printf("cannot start skipping %s \n", filename)
		}
	}
}
