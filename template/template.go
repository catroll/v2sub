package template

import (
	"encoding/json"

	"github.com/arkrz/v2sub/types"
)

const (
	ListenOnLocalAddr = "127.0.0.1"
	ListenOnWanAddr   = "0.0.0.0"

	ListenOnSocksProtocol = "socks"
	ListenOnSocksPort     = 1081

	ListenOnHttpProtocol = "http"
	ListenOnHttpPort     = 1082
)

var domainStrategy = "ipondemand"
var loglevel = "debug"

var ConfigTemplate = &types.Config{
	SubUrl: "",
	Nodes:  types.Nodes{},
	V2rayConfig: types.V2ray{
		LogConfig: &types.LogConfig{
			LogLevel: loglevel,
		},
		RouterConfig: &types.RouterConfig{
			RuleList:       nil,
			DomainStrategy: domainStrategy,
		},
		OutboundConfigs: []types.OutboundConfig{},
		InboundConfigs: []types.InboundConfig{
			{
				Protocol: ListenOnSocksProtocol,
				Port:     ListenOnSocksPort,
				ListenOn: ListenOnLocalAddr,
				//PortRange: &conf.PortRange{ // [from, to]
				//	From: 1080,
				//	To:   1080,
				//}, // https://github.com/v2ray/v2ray-core/blob/v4.21.3/app/proxyman/inbound/always.go#L91
				//ListenOn: &conf.Address{Address: net.ParseAddress("127.0.0.1")},
			},
			{
				Protocol: ListenOnHttpProtocol,
				Port:     ListenOnHttpPort,
				ListenOn: ListenOnLocalAddr,
			},
		},
	},
}

// DefaultDNSConfigs 默认路由规则
// 参考 https://toutyrater.github.io/routing/configurate_rules.html
var DefaultDNSConfigs = &types.DNSConfig{Servers: []json.RawMessage{
	[]byte(`"114.114.114.114"`),
	[]byte(
		`{
			"address": "1.1.1.1",
			"port": 53,
			"domains": [
				"geosite:geolocation-!cn"
			]
		}`),
}}

var DefaultRouterConfigs = &types.RouterConfig{
	RuleList: []json.RawMessage{
		[]byte(
			`{
				"type": "field",
				"outboundTag": "direct",
				"domain": [
					"geosite:cn"
				]
			}`),
		[]byte(
			`{
                "type": "field",
                "outboundTag": "direct",
                "ip": [
                    "geoip:cn",
                    "geoip:private"
                ]
            }`),
		[]byte(
			`{
                "type": "field",
                "outboundTag": "proxy",
                "network": "udp,tcp"
            }`),
	},
	DomainStrategy: domainStrategy,
}

var DefaultOutboundConfigs = []types.OutboundConfig{
	{
		Protocol: "freedom",
		Tag:      "direct",
	},
	{
		Protocol: "blackhole",
		Tag:      "block",
	},
}
