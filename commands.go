package tempo

import "github.com/mitchellh/cli"

var Commands = map[string]cli.CommandFactory{
	"cleanup": func() (cli.Command, error) {
		return &CleanupCommand{}, nil
	},
	"list": func() (cli.Command, error) {
		return &ListCommand{}, nil
	},
	"setup": func() (cli.Command, error) {
		return &SetupCommand{}, nil
	},
	"track": func() (cli.Command, error) {
		return &TrackCommand{}, nil
	},
}
