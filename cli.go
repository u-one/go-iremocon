package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bash0C7/go-iremocon/iremocon"
	"github.com/codegangsta/cli"
)

var Version string = "0.0.0"

func action(c *cli.Context, iremoconFunc func(net.Conn, []string) (string, error)) {
	network := "tcp"
	address := "10.0.1.200:51013"

	conn, _ := net.Dial(network, address)
	defer conn.Close()
	ret, err := iremoconFunc(conn, c.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		println(ret, c.Args())
	}
}

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
				f := func(conn net.Conn, args []string) (string, error) {
					//a := args[0]
					return iremocon.Au(conn) //引数は？
				}
				action(c, f)
			},
		},
		{
			Name:  "vr",
			Usage: "vr command",
			Action: func(c *cli.Context) {
				f := func(conn net.Conn, args []string) (string, error) {
					//a := args[0]
					return iremocon.Vr(conn) //引数は？
				}
				action(c, f)
			},
		},
	}
	return app
}
