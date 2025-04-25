package explorer

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"main/core/blocknumber"
	"main/core/database"
	"main/core/sendtx"
)

func Explorer() {
	app := fiber.New()
	app.Post("/chainservice/generateaddress", generateaddress)
	app.Post("/chainservice/faucet", faucet)
	app.Post("/chainservice/balance", balance)
	app.Post("/chainservice/depositbalance", depositbalance)
	app.Post("/chainservice/deposit", deposit)
	app.Post("/chainservice/withdraw", withdraw)
	log.Fatal(app.Listen(":4000"))
}

type Info struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Nonce    uint64 `json:"nonce"`
}

type GetInfo struct {
	Number string `json:"number"`
}

type AddressInfo struct {
	Address string `json:"address"`
}

type GetUserName struct {
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func faucet(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &AddressInfo{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.Faucet(payload.Address)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}

func withdraw(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &Info{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.Withdraw(payload.Username, payload.Address, payload.Amount, payload.Nonce)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}

func deposit(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &Info{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.Deposit(payload.Username, payload.Address, payload.Amount, payload.Nonce)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}

func balance(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &AddressInfo{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	details := blocknumber.GetAddressBalance(payload.Address)

	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: details})
}

func depositbalance(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &AddressInfo{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	details := sendtx.GetDepositBalance(payload.Address)

	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: details})
}

func generateaddress(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	payload := &GetUserName{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	ecdsaPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	newAddress := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	pbytes := crypto.FromECDSA(ecdsaPrivateKey)
	database.InsertAddressKey(payload.Username, newAddress, hexutil.Encode(pbytes))
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: newAddress.String()})
}

func EcdsaAddressFromPrivateKey(ecdsaPrivateKey *ecdsa.PrivateKey) common.Address {
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := crypto.PubkeyToAddress(*pub)
	return rAddress
}
