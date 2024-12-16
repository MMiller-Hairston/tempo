package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mmiller-hairston/tempo"
)

var (
	app     string = "tempo"
	version string = "0.0.1"
)

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	c.Commands = tempo.Commands

	es, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(es)
}
