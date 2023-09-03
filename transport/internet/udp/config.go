package udp

import (
	"github.com/moqsien/xraycore/common"
	"github.com/moqsien/xraycore/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
