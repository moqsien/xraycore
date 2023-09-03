package tcp

import (
	"github.com/moqsien/xraycore/common"
	"github.com/moqsien/xraycore/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
