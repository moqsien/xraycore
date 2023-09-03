package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/moqsien/xraycore/app/dispatcher"
	_ "github.com/moqsien/xraycore/app/proxyman/inbound"
	_ "github.com/moqsien/xraycore/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/moqsien/xraycore/app/commander"
	_ "github.com/moqsien/xraycore/app/log/command"
	_ "github.com/moqsien/xraycore/app/proxyman/command"
	_ "github.com/moqsien/xraycore/app/stats/command"

	// Developer preview services
	_ "github.com/moqsien/xraycore/app/observatory/command"

	// Other optional features.
	_ "github.com/moqsien/xraycore/app/dns"
	_ "github.com/moqsien/xraycore/app/dns/fakedns"
	_ "github.com/moqsien/xraycore/app/log"
	_ "github.com/moqsien/xraycore/app/metrics"
	_ "github.com/moqsien/xraycore/app/policy"
	_ "github.com/moqsien/xraycore/app/reverse"
	_ "github.com/moqsien/xraycore/app/router"
	_ "github.com/moqsien/xraycore/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/moqsien/xraycore/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/moqsien/xraycore/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/moqsien/xraycore/proxy/blackhole"
	_ "github.com/moqsien/xraycore/proxy/dns"
	_ "github.com/moqsien/xraycore/proxy/dokodemo"
	_ "github.com/moqsien/xraycore/proxy/freedom"
	_ "github.com/moqsien/xraycore/proxy/http"
	_ "github.com/moqsien/xraycore/proxy/loopback"
	_ "github.com/moqsien/xraycore/proxy/shadowsocks"
	_ "github.com/moqsien/xraycore/proxy/socks"
	_ "github.com/moqsien/xraycore/proxy/trojan"
	_ "github.com/moqsien/xraycore/proxy/vless/inbound"
	_ "github.com/moqsien/xraycore/proxy/vless/outbound"
	_ "github.com/moqsien/xraycore/proxy/vmess/inbound"
	_ "github.com/moqsien/xraycore/proxy/vmess/outbound"
	_ "github.com/moqsien/xraycore/proxy/wireguard"

	// Transports
	_ "github.com/moqsien/xraycore/transport/internet/domainsocket"
	_ "github.com/moqsien/xraycore/transport/internet/grpc"
	_ "github.com/moqsien/xraycore/transport/internet/http"
	_ "github.com/moqsien/xraycore/transport/internet/kcp"
	_ "github.com/moqsien/xraycore/transport/internet/quic"
	_ "github.com/moqsien/xraycore/transport/internet/reality"
	_ "github.com/moqsien/xraycore/transport/internet/tcp"
	_ "github.com/moqsien/xraycore/transport/internet/tls"
	_ "github.com/moqsien/xraycore/transport/internet/udp"
	_ "github.com/moqsien/xraycore/transport/internet/websocket"

	// Transport headers
	_ "github.com/moqsien/xraycore/transport/internet/headers/http"
	_ "github.com/moqsien/xraycore/transport/internet/headers/noop"
	_ "github.com/moqsien/xraycore/transport/internet/headers/srtp"
	_ "github.com/moqsien/xraycore/transport/internet/headers/tls"
	_ "github.com/moqsien/xraycore/transport/internet/headers/utp"
	_ "github.com/moqsien/xraycore/transport/internet/headers/wechat"
	_ "github.com/moqsien/xraycore/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/moqsien/xraycore/main/json"
	_ "github.com/moqsien/xraycore/main/toml"
	_ "github.com/moqsien/xraycore/main/yaml"

	// Load config from file or http(s)
	_ "github.com/moqsien/xraycore/main/confloader/external"

	// Commands
	_ "github.com/moqsien/xraycore/main/commands/all"
)
