package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/core/blocknumber"
	"main/core/database"
	"main/trace"
	"math/big"
	"strings"
)

func gentx(key string, value int64, gaslimit uint64) *bind.TransactOpts {

	if strings.HasPrefix(key, "0x") {
		key = strings.Trim(key, "0x")
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("解析账户私钥错误", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {

		log.Fatalf("连接区块链失败，请重启区块链后端. err =  %v", err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value) // in wei
	auth.GasLimit = gaslimit       // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrUpdateUserData(number string, workLocation string, workContent string, person string, advice string, workTime string, remarks string, imagesUrl string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx("", 0, 300000)
	if auth == nil {
		return ""
	}
	tx, err := instance.AddRecord(auth, number, workLocation, workContent, person, advice, workTime, remarks, imagesUrl)
	if err != nil {
		fmt.Println("error creating instance:%s", err)
		return ""
		//log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func GetUserData(number string) trace.WorkRecordWorkDetails {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	var details trace.WorkRecordWorkDetails
	details, err = instance.GetRecord(nil, number)
	if err != nil {
		fmt.Println("error getting data:%s", err)
		//log.Fatal(err)
	}
	return details
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func BuildContract() (common.Address, uint64) {
	ok, addr, num := database.GetContracts()
	if !ok {
		auth := gentx("0x799bae9074c08e5445fa20f6a4b104ece28933d4dd4241f1fc89e35a335af17d", 0, 500000)
		res, tx, _, err := trace.DeployTrace(auth, config.Client)
		if err != nil {
			log.Fatalf("初次部署合约失败，err:%s", err)
		}
		num = blocknumber.GetBlockNumber(tx.Hash().Hex())
		database.InsertContracts(res, num)
		return res, num
	} else {
		return common.HexToAddress(addr), num
	}
}
