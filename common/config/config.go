package config

import (
	"main/common/dbconn"
	"main/common/ethconn"
)

var (
	EthServer = "http://localhost:7545"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0xE7C748D172f25EaCdefa12DAc5D9E43d2b433aeE"
	Dba       = dbconn.Buildconnect()
)
