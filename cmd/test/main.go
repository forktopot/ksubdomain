package main

import (
	"github.com/forktopot/ksubdomain/v2/cmd/ksubdomainstart"
	"github.com/forktopot/ksubdomain/v2/pkg/core/conf"
	"github.com/forktopot/ksubdomain/v2/pkg/core/gologger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:    conf.AppName,
		Version: conf.Version,
		Usage:   conf.Description,
		Commands: []*cli.Command{
			ksubdomainstart.EnumCommand,
			ksubdomainstart.VerifyCommand,
			ksubdomainstart.TestCommand,
			ksubdomainstart.DeviceCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
}
