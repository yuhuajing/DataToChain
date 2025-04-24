package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"main/common/config"
	"main/core/ethclientevent"
	"main/core/sendtx"
	"main/explorer"
	"math/big"

	"time"
)

func main() {
	sc, num := sendtx.BuildContract()
	SubStakingEvent(sc, num)
	go explorer.Explorer()
	time.Sleep(10 * time.Second)
}

var (
	logsChan = make(chan types.Log, 0)
)

func SubStakingEvent(contract common.Address, number uint64) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contract},
		FromBlock: big.NewInt(int64(number)),
		Topics:    [][]common.Hash{{common.HexToHash("")}},
	}

	subevents, err := config.Client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		fmt.Println(fmt.Errorf("Subscribe Event error: %v", err))
		log.Fatal(err)
	}
	for {
		select {
		case err := <-subevents.Err():
			fmt.Println(fmt.Errorf("Parse Event error: %v", err))

		case lplog := <-logsChan:
			ethclientevent.ParseEventLog(contract, lplog)
		}
	}
}
