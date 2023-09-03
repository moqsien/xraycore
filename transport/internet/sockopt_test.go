package internet_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/moqsien/xraycore/common"
	"github.com/moqsien/xraycore/common/buf"
	"github.com/moqsien/xraycore/testing/servers/tcp"
	. "github.com/moqsien/xraycore/transport/internet"
)

func TestTCPFastOpen(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: func(b []byte) []byte {
			return b
		},
	}
	dest, err := tcpServer.StartContext(context.Background(), &SocketConfig{Tfo: 256})
	common.Must(err)
	defer tcpServer.Close()

	ctx := context.Background()
	dialer := DefaultSystemDialer{}
	conn, err := dialer.Dial(ctx, nil, dest, &SocketConfig{
		Tfo: 1,
	})
	common.Must(err)
	defer conn.Close()

	_, err = conn.Write([]byte("abcd"))
	common.Must(err)

	b := buf.New()
	common.Must2(b.ReadFrom(conn))
	if r := cmp.Diff(b.Bytes(), []byte("abcd")); r != "" {
		t.Fatal(r)
	}
}
