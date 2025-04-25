package config

import (
	"main/common/dbconn"
	"main/common/ethconn"
)

var (
	EthServer = "ws://localhost:7545"
	Client    = ethconn.ConnBlockchain(EthServer)
	Dba       = dbconn.Buildconnect()
)
