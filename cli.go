package main

import (
	"os"
  "github.com/bash0C7/go-iremocon/lib"
	"github.com/codegangsta/cli"
)

var Version string = "0.0.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go-iremocon"
	app.Usage = "Iremocon is a golang library for managing iRemocon through telnet, inspired by r7kamura/iremocon"
	app.Version = Version
	app.Author = "bash0C7"
	app.Email = "ksb.4038.nullpointer+github@gmail.com"
	app.Commands = []cli.Command{
		{
			Name:    "au",
			Usage:   "au command",
			Action: func(c *cli.Context) {
        iremocon.au
				println("added task: ", c.Args())
			},
		},
		{
			Name:    "complete",
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args())
			},
		}
	}
	return app
}
