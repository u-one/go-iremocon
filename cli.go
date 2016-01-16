package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bash0C7/go-iremocon/iremocon"
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
			Name:  "au",
			Usage: "au command",
			Action: func(c *cli.Context) {
				conn, _ := net.Dial("tcp", "10.0.1.200:51013")
				defer conn.Close()
				ret, err := iremocon.Au(conn)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					println(ret, c.Args())
				}
			},
		},
		{
			Name:  "vr",
			Usage: "vr command",
			Action: func(c *cli.Context) {
				conn, _ := net.Dial("tcp", "10.0.1.200:51013")
				defer conn.Close()
				ret, err := iremocon.Vr(conn)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					println(ret, c.Args())
				}
			},
		},
	}
	return app
}
