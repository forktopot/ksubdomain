package ksubdomainstart

import (
	"github.com/boy-hack/ksubdomain/core/options"
	"github.com/boy-hack/ksubdomain/runner"
	"github.com/urfave/cli/v2"
)

var TestCommand = &cli.Command{
	Name:  runner.TestType,
	Usage: "测试本地网卡的最大发送速度",
	Action: func(c *cli.Context) error {
		ether := options.GetDeviceConfig()
		runner.TestSpeed(ether)
		return nil
	},
}
