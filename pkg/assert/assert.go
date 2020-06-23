package assert

import (
	"github.com/getgauge-contrib/gauge-go/testsuit"
)

// NoError confirms the error returned is null
func NoError(err error) {
	if err != nil {
		testsuit.T.Fail(err)
	}
}
