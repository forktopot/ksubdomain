package runner

import (
	"context"
	"github.com/forktopot/ksubdomain/core"
	"github.com/forktopot/ksubdomain/core/gologger"
	"github.com/forktopot/ksubdomain/core/options"
	"github.com/forktopot/ksubdomain/runner/outputter"
	"github.com/forktopot/ksubdomain/runner/outputter/output"
	"github.com/forktopot/ksubdomain/runner/processbar"
	"testing"
)

func TestRunner(t *testing.T) {
	process := processbar.FakeScreenProcess{}
	screenPrinter, _ := output.NewScreenOutputNoWidth()
	domains := []string{"stu.baidu.com", "haokan.baidu.com"}
	domainChanel := make(chan string)
	go func() {
		for _, d := range domains {
			domainChanel <- d
		}
		close(domainChanel)
	}()
	opt := &options.Options{
		Rate:        options.Band2Rate("1m"),
		Domain:      domainChanel,
		DomainTotal: 2,
		Resolvers:   options.GetResolvers(""),
		Silent:      false,
		TimeOut:     5,
		Retry:       1,
		Method:      VerifyType,
		DnsType:     "a",
		Writer: []outputter.Output{
			screenPrinter,
		},
		ProcessBar: &process,
		EtherInfo:  options.GetDeviceConfig(),
	}
	opt.Check()
	r, err := New(opt)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
	ctx := context.Background()
	r.RunEnumeration(ctx)
	r.Close()
}

func TestRunnerEnum(t *testing.T) {
	process := processbar.ScreenProcess{}
	screenPrinter, _ := output.NewScreenOutputNoWidth()
	domains := core.GetDefaultSubdomainData()
	domainChanel := make(chan string)
	go func() {
		for _, d := range domains {
			domainChanel <- d + ".baidu.com"
		}
		close(domainChanel)
	}()
	opt := &options.Options{
		Rate:        options.Band2Rate("1m"),
		Domain:      domainChanel,
		DomainTotal: len(domains),
		Resolvers:   options.GetResolvers(""),
		Silent:      false,
		TimeOut:     5,
		Retry:       1,
		Method:      VerifyType,
		DnsType:     "a",
		Writer: []outputter.Output{
			screenPrinter,
		},
		ProcessBar: &process,
		EtherInfo:  options.GetDeviceConfig(),
	}
	opt.Check()
	r, err := New(opt)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
	ctx := context.Background()
	r.RunEnumeration(ctx)
	r.Close()
}

func TestManyRunner(t *testing.T) {
	for i := 0; i < 5; i++ {
		TestRunner(t)
	}
}
