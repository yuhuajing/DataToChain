package sendtx

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/core/blocknumber"
	"main/core/database"
	"main/trace"
	"math/big"
	"strings"
)

func gentx(key string, value *big.Int, gaslimit uint64) *bind.TransactOpts {

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
	auth.Value = value       // in wei
	auth.GasLimit = gaslimit // in units
	auth.GasPrice = gasPrice
	return auth
}

func Deposit(user, address, amount string, nonce uint64) string {
	_, addr, _ := database.GetContracts()
	instance, err := trace.NewTrace(common.HexToAddress(addr), config.Client)
	if err != nil {
		log.Fatalf("error creating callerinstance instance:%s", err)
	}
	key := database.GetAddressKey(user, common.HexToAddress(address))

	bigInt := new(big.Int)

	// 将字符串转换为 big.Int
	bigInt.SetString(amount, 10) // 10 表示十进制

	auth := gentx(key, bigInt, 100000)
	if auth == nil {
		return ""
	}
	tx, err := instance.Deposit(auth, nonce)
	if err != nil {
		fmt.Println("error creating instance:%s", err)
		return ""
		//log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func Faucet(address string) string {
	var ctx = context.Background()
	// Import the from address
	//0x96216849c49358B10257cb55b28eA603c874b05E
	key := "d56c09fa8fbf8134039d9b2c4c22edeb993c856451c94267c65f90dcb04e9cf4"
	// Decode the provided private key.
	if ok := strings.HasPrefix(key, "0x"); ok {
		key, _ = strings.CutPrefix(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Errors in parsing key %v", err)
	}
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	// Compute the Ethereum address of the signer from the public key.
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Retrieve the nonce for the signer's account, representing the transaction count.

	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the transaction fields, including the recipient address, value, and gas parameters.
	toAddr := common.HexToAddress(address)
	amount := "100000000000000000"
	bigInt := new(big.Int)

	// 将字符串转换为 big.Int
	bigInt.SetString(amount, 10) // 10 表示十进制

	chainID, err := config.Client.ChainID(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	gasPrice := SuggestedGasPrice(ctx)
	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      21000,
		To:       &toAddr,
		Value:    bigInt,
	}

	tx := types.NewTx(&txData)

	// Sign the transaction with the private key of the sender.
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Encode the signed transaction into RLP (Recursive Length Prefix) format for transmission.
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)

	if err != nil {
		log.Fatalln(err)
	}

	err = config.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx.Hash().Hex()

}

func SuggestedGasPrice(ctx context.Context) *big.Int {
	// Retrieve the currently suggested gas price for a new transaction.
	gasPrice, err := config.Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	return gasPrice
}

func Withdraw(user, address, amount string, nonce uint64) string {
	_, addr, _ := database.GetContracts()
	instance, err := trace.NewTrace(common.HexToAddress(addr), config.Client)
	if err != nil {
		log.Fatalf("error creating instance instance:%s", err)
	}
	key := database.GetAddressKey(user, common.HexToAddress(address))

	auth := gentx(key, big.NewInt(0), 100000)
	if auth == nil {
		return ""
	}

	bigInt := new(big.Int)

	// 将字符串转换为 big.Int
	bigInt.SetString(amount, 10) // 10 表示十进制

	tx, err := instance.Withdrawl(auth, bigInt.Uint64(), nonce)
	if err != nil {
		fmt.Println("error creating instance:%s", err)
		return ""
		//log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func BuildContract() (common.Address, uint64) {
	ok, addr, num := database.GetContracts()
	if !ok {
		auth := gentx("0x799bae9074c08e5445fa20f6a4b104ece28933d4dd4241f1fc89e35a335af17d", big.NewInt(0), 700000)
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

func GetDepositBalance(address string) string {
	addrs := common.HexToAddress(address)
	ok, addr, _ := database.GetContracts()
	if ok {
		contractAddr := common.HexToAddress(addr)
		instance, _ := trace.NewTraceCaller(contractAddr, config.Client)
		balance, _ := instance.Locked(nil, addrs)
		return balance.String()
	} else {
		BuildContract()
	}
	return ""
}
