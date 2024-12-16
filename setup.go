package tempo

import (
	"fmt"
	"os"
	"syscall"
)

type SetupCommand struct{}

const (
	storage = "/opt/ds"
)

func (c *SetupCommand) Help() string {
	return "Setup the tempo cli"
}

func (c *SetupCommand) Run(args []string) int {
	syscall.Umask(0)
	err := os.MkdirAll(storage, os.ModePerm)
	if err != nil {
		fmt.Printf("Unable to create the storage dir. Try running again with 'sudo'.\n")
		os.Exit(1)
	}
	return 0
}

func (c *SetupCommand) Synopsis() string {
	return "Create the necessary config files and directories to be able to use the Tempo cli."
}
