package scenarios

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
	"time"

	"github.com/sagernet/sing-shadowsocks/shadowaead_2022"
	"github.com/moqsien/xraycore/app/log"
	"github.com/moqsien/xraycore/app/proxyman"
	"github.com/moqsien/xraycore/common"
	clog "github.com/moqsien/xraycore/common/log"
	"github.com/moqsien/xraycore/common/net"
	"github.com/moqsien/xraycore/common/serial"
	"github.com/moqsien/xraycore/core"
	"github.com/moqsien/xraycore/proxy/dokodemo"
	"github.com/moqsien/xraycore/proxy/freedom"
	"github.com/moqsien/xraycore/proxy/shadowsocks_2022"
	"github.com/moqsien/xraycore/testing/servers/tcp"
	"github.com/moqsien/xraycore/testing/servers/udp"
	"golang.org/x/sync/errgroup"
)

func TestShadowsocks2022Tcp(t *testing.T) {
	for _, method := range shadowaead_2022.List {
		password := make([]byte, 32)
		rand.Read(password)
		t.Run(method, func(t *testing.T) {
			testShadowsocks2022Tcp(t, method, base64.StdEncoding.EncodeToString(password))
		})
	}
}

func TestShadowsocks2022Udp(t *testing.T) {
	for _, method := range shadowaead_2022.List {
		password := make([]byte, 32)
		rand.Read(password)
		t.Run(method, func(t *testing.T) {
			testShadowsocks2022Udp(t, method, base64.StdEncoding.EncodeToString(password))
		})
	}
}

func testShadowsocks2022Tcp(t *testing.T, method string, password string) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(serverPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&shadowsocks_2022.ServerConfig{
					Method:  method,
					Key:     password,
					Network: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(clientPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(dest.Address),
					Port:     uint32(dest.Port),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&shadowsocks_2022.ClientConfig{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(serverPort),
					Method:  method,
					Key:     password,
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	var errGroup errgroup.Group
	for i := 0; i < 10; i++ {
		errGroup.Go(testTCPConn(clientPort, 10240*1024, time.Second*20))
	}

	if err := errGroup.Wait(); err != nil {
		t.Error(err)
	}
}

func testShadowsocks2022Udp(t *testing.T, method string, password string) {
	udpServer := udp.Server{
		MsgProcessor: xor,
	}
	udpDest, err := udpServer.Start()
	common.Must(err)
	defer udpServer.Close()

	serverPort := udp.PickPort()
	serverConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(serverPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&shadowsocks_2022.ServerConfig{
					Method:  method,
					Key:     password,
					Network: []net.Network{net.Network_UDP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	udpClientPort := udp.PickPort()
	clientConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(udpClientPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(udpDest.Address),
					Port:     uint32(udpDest.Port),
					Networks: []net.Network{net.Network_UDP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&shadowsocks_2022.ClientConfig{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(serverPort),
					Method:  method,
					Key:     password,
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	var errGroup errgroup.Group
	for i := 0; i < 10; i++ {
		errGroup.Go(testUDPConn(udpClientPort, 1024, time.Second*5))
	}

	if err := errGroup.Wait(); err != nil {
		t.Error(err)
	}
}