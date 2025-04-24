package ethclientevent

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"main/core/database"
)

func ParseEventLog(contract common.Address, log types.Log) {
	database.InsertLogs(contract, log)
}
