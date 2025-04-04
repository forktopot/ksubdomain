package ksubdomainstart

import (
	"github.com/forktopot/ksubdomain/core/conf"
	"github.com/forktopot/ksubdomain/core/gologger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:    conf.AppName,
		Version: conf.Version,
		Usage:   conf.Description,
		Commands: []*cli.Command{
			EnumCommand,
			VerifyCommand,
			TestCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
}
