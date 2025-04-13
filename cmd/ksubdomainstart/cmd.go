package ksubdomainstart

import (
	"os"

	"github.com/forktopot/ksubdomain/pkg/core/conf"
	"github.com/forktopot/ksubdomain/pkg/core/gologger"
	"github.com/urfave/cli/v2"
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
			DeviceCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
}
