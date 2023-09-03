package udp

import (
	"github.com/moqsien/xraycore/common/buf"
	"github.com/moqsien/xraycore/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
