package command_test

import (
	"context"
	"testing"

	"github.com/moqsien/xraycore/app/dispatcher"
	"github.com/moqsien/xraycore/app/log"
	. "github.com/moqsien/xraycore/app/log/command"
	"github.com/moqsien/xraycore/app/proxyman"
	_ "github.com/moqsien/xraycore/app/proxyman/inbound"
	_ "github.com/moqsien/xraycore/app/proxyman/outbound"
	"github.com/moqsien/xraycore/common"
	"github.com/moqsien/xraycore/common/serial"
	"github.com/moqsien/xraycore/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
