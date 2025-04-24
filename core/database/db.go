package database

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"main/common/config"
	"main/common/tabletypes"
	"main/trace"
	"strings"

	"github.com/jinzhu/gorm"
)

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

func GetAddressKey(user string, address common.Address) *tabletypes.AddressKey {
	dba := config.Dba
	addr := strings.ToLower(address.String())

	result := dba.Where("user = ? and address = ?", user, addr).Find(&tabletypes.AddressKey{}).Limit(1)

	var res = &tabletypes.AddressKey{}
	result.Scan(res)
	return res
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

	data, _ := filter.ParseOwnershipTransferred(log)
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

func GetContracts() (bool, string, uint64) {
	dba := config.Dba

	result := dba.Find(&tabletypes.Contracts{}).Limit(1)
	if result.RowsAffected == 0 {
		return false, "", 0
	}
	var res = &tabletypes.Contracts{}
	result.Scan(res)
	return true, res.Address, res.Block
}

func ModifyOwner(db *gorm.DB, address string, id int, owner string) {
	res := db.Model(&tabletypes.Owner{}).Where("address = ? AND tokenid = ?", address, id).First(&tabletypes.Owner{})
	if res.RowsAffected == 0 {
		db.Create(&tabletypes.Owner{
			Address: address,
			Owner:   owner,
			Tokenid: int64(id),
		})
	} else {
		db.Model(&tabletypes.Owner{}).Where("address = ? AND tokenid = ?", address, id).Update("owner", owner)
	}
}
