package ethconn

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ConnBlockchain(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	if err != nil {
		log.Fatalf("区块链连接失败，请确认已启动Ganache，error:%v", err)
	}
	return nclient
}
