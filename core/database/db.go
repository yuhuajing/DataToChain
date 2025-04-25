package database

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"main/common/config"
	"main/common/tabletypes"
	"main/trace"
	"strings"
)

func GetAddressKey(user string, address common.Address) string {
	dba := config.Dba
	addr := strings.ToLower(address.String())

	result := dba.Where("user = ? and address = ?", user, addr).Find(&tabletypes.AddressKey{}).Limit(1)
	var res = &tabletypes.AddressKey{}
	result.Scan(res)
	return res.Key
}

// func Insert(dba *gorm.DB, logdata *types.Log) {
func InsertAddressKey(user string, address common.Address, key string) {
	dba := config.Dba
	addr := strings.ToLower(address.String())

	result := dba.Where("address = ?", addr).Find(&tabletypes.AddressKey{})
	if result.RowsAffected == 0 {
		dba.Create(&tabletypes.AddressKey{
			User:    user,
			Address: addr,
			Key:     key,
		})
	}
}

func GetContracts() (bool, string, uint64) {
	dba := config.Dba

	result := dba.First(&tabletypes.Contracts{})
	if result.RowsAffected == 0 {
		return false, "", 0
	}
	var res = &tabletypes.Contracts{}
	result.Scan(res)
	return true, res.Address, res.Block
}

func InsertContracts(address common.Address, number uint64) {
	dba := config.Dba
	addr := strings.ToLower(address.String())
	result := dba.Where("address = ?", addr).Find(&tabletypes.Contracts{})
	if result.RowsAffected == 0 {
		dba.Create(&tabletypes.Contracts{
			Address: addr,
			Block:   number,
		})
	}
}

func InsertLogs(contract common.Address, log types.Log) {
	filter, _ := trace.NewTraceFilterer(contract, config.Client)
	data, _ := filter.ParseAssert(log)
	blocknum := data.Blocknum
	ts := data.Timestamp
	sender := strings.ToLower(data.From.Hex())
	nonce := data.Nonce
	to := strings.ToLower(log.Address.Hex())
	amount := data.Amount
	var op = "deposit"
	if data.Op == 1 {
		op = "withdrawl"
	}
	index := log.Index
	txhash := log.TxHash.Hex()

	dba := config.Dba

	result := dba.Where("txhash = ?", txhash).Find(&tabletypes.Assert{})
	if result.RowsAffected == 0 {
		dba.Create(&tabletypes.Assert{
			Blocknumber: blocknum.Uint64(),
			Timestamp:   ts.Uint64(),
			From:        sender,
			Nonce:       nonce,
			To:          to,
			Amount:      amount.Uint64(),
			Types:       op,
			Logindex:    index,
			Txhash:      txhash,
		})
	}
}
