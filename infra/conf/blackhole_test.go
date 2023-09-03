package conf_test

import (
	"testing"

	"github.com/moqsien/xraycore/common/serial"
	. "github.com/moqsien/xraycore/infra/conf"
	"github.com/moqsien/xraycore/proxy/blackhole"
)

func TestHTTPResponseJSON(t *testing.T) {
	creator := func() Buildable {
		return new(BlackholeConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"response": {
					"type": "http"
				}
			}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			},
		},
		{
			Input:  `{}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{},
		},
	})
}
