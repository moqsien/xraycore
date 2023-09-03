package conf_test

import (
	"testing"

	"github.com/moqsien/xraycore/common/net"
	"github.com/moqsien/xraycore/common/protocol"
	"github.com/moqsien/xraycore/common/serial"
	. "github.com/moqsien/xraycore/infra/conf"
	"github.com/moqsien/xraycore/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "xray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "xray-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
