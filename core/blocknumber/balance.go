package blocknumber

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"main/common/config"
)

func GetAddressBalance(address string) string {
	balance, err := config.Client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("err in GetBlockTime: %v", err)
	}

	return balance.String()
}
