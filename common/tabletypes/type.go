package tabletypes

type AddressKey struct {
	ID      uint   `gorm:"primary_key"`
	User    string `gorm:"user"`
	Address string `json:"address"`
	Key     string `json:"key"`
}

type Contracts struct {
	ID      uint   `gorm:"primary_key"`
	Address string `json:"address"`
	Block   uint64 `json:"block"`
}

type AssertOperation struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   string `json:"timestamp"`
	Address     string `json:"address"`
	From        string `json:"from" gencodec:"required"`
	To          string `json:"to" gencodec:"required"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Amount      uint64 `json:"amount"`
	Types       string `json:"types"`
}
