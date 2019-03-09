package main

import (
	"github.com/lyraeve/lyrae-go/command"
	"github.com/urfave/cli"
	"os"
)

func main() {
	os.Chdir(".")

	app := cli.NewApp()
	app.Name = "Video Finder"
	app.Description = "The finder for action video"
	app.Commands = []cli.Command{
		command.FindCommand,
	}

	app.Run(os.Args)
}
