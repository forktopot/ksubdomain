package outputter

import (
	"github.com/forktopot/ksubdomain/v2/pkg/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close() error
}
