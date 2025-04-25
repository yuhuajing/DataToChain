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
	"sync"
)

func main() {
	sc, num := sendtx.BuildContract()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(sc common.Address, num uint64) {
		defer wg.Done() // 确保在协程完成时调用 Done
		SubStakingEvent(sc, num)
	}(sc, num)

	explorer.Explorer()

	wg.Wait()
}

var (
	logsChan = make(chan types.Log, 0)
)

func SubStakingEvent(contractAddress common.Address, number uint64) {
	//contractAddress := common.HexToAddress("0x28ddd076b87131356c1f70278f64eba3f8de307e")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: big.NewInt(int64(number)),
		Topics:    [][]common.Hash{{common.HexToHash("0x87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f703")}},
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
			ethclientevent.ParseEventLog(contractAddress, lplog)
		}
	}
}
