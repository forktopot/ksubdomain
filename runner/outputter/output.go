package outputter

import (
	"github.com/forktopot/ksubdomain/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close()
}
