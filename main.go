package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	os.Chdir(".")

	app := cli.NewApp()
	app.Name = "Video Finder"
	app.Description = "The finder for action video"

	app.Run(os.Args)
}
