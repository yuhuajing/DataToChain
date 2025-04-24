package explorer

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	fiber "github.com/gofiber/fiber/v2"

	"log"
	"main/core/database"
	"main/core/sendtx"
)

func Explorer() {
	app := fiber.New()
	app.Post("/chainservice/generateaddress", generateaddress)
	app.Post("/chainservice/uploadchain", uploadchain)
	app.Post("/chainservice/updatechain", uploadchain)
	app.Post("/chainservice/getchain", getchain)
	log.Fatal(app.Listen(":4000"))
}

type Info struct {
	Number       string `json:"number"`
	Worklocation string `json:"worklocation"`
	Workcontent  string `json:"workcontent"`
	Persion      string `json:"persion"`
	Advice       string `json:"advice"`
	Worktime     string `json:"worktime"`
	Remarks      string `json:"remarks"`
	Imagesurl    string `json:"imagesurl"`
}

type GetInfo struct {
	Number string `json:"number"`
}

type GetUserName struct {
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func uploadchain(c *fiber.Ctx) error {
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
	hash := sendtx.AddOrUpdateUserData(payload.Number, payload.Worklocation, payload.Workcontent, payload.Persion, payload.Advice, payload.Worktime, payload.Remarks, payload.Imagesurl)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}

func getchain(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	//fmt.Println(string(c.Body()))
	payload := &GetInfo{}
	if err := c.BodyParser(payload); err != nil {
		//fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	details := sendtx.GetUserData(payload.Number)
	data, _ := json.Marshal(details)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: string(data)})
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
