package main

import (
	"fmt"
	"github.com/forktopot/ksubdomain/cmd/ksubdomainstart"
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
			ksubdomainstart.EnumCommand,
			ksubdomainstart.VerifyCommand,
			ksubdomainstart.TestCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
	fmt.Println("123")
}
