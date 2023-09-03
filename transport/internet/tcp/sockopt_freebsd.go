//go:build freebsd
// +build freebsd

package tcp

import (
	"github.com/moqsien/xraycore/common/net"
	"github.com/moqsien/xraycore/transport/internet"
	"github.com/moqsien/xraycore/transport/internet/stat"
)

// GetOriginalDestination from tcp conn
func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	la := conn.LocalAddr()
	ra := conn.RemoteAddr()
	ip, port, err := internet.OriginalDst(la, ra)
	if err != nil {
		return net.Destination{}, newError("failed to get destination").Base(err)
	}
	dest := net.TCPDestination(net.IPAddress(ip), net.Port(port))
	if !dest.IsValid() {
		return net.Destination{}, newError("failed to parse destination.")
	}
	return dest, nil
}
