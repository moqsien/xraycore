package core_test

import (
	"testing"

	"github.com/moqsien/xraycore/app/dispatcher"
	"github.com/moqsien/xraycore/app/proxyman"
	"github.com/moqsien/xraycore/common"
	"github.com/moqsien/xraycore/common/net"
	"github.com/moqsien/xraycore/common/protocol"
	"github.com/moqsien/xraycore/common/serial"
	"github.com/moqsien/xraycore/common/uuid"
	. "github.com/moqsien/xraycore/core"
	"github.com/moqsien/xraycore/features/dns"
	"github.com/moqsien/xraycore/features/dns/localdns"
	_ "github.com/moqsien/xraycore/main/distro/all"
	"github.com/moqsien/xraycore/proxy/dokodemo"
	"github.com/moqsien/xraycore/proxy/vmess"
	"github.com/moqsien/xraycore/proxy/vmess/outbound"
	"github.com/moqsien/xraycore/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
