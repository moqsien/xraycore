//go:build !linux && !freebsd && !darwin
// +build !linux,!freebsd,!darwin

package tcp

import (
	"github.com/moqsien/xraycore/common/net"
	"github.com/moqsien/xraycore/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
