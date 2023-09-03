package all

import (
	"github.com/moqsien/xraycore/main/commands/all/api"
	"github.com/moqsien/xraycore/main/commands/all/tls"
	"github.com/moqsien/xraycore/main/commands/base"
)

// go:generate go run github.com/moqsien/xraycore/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
