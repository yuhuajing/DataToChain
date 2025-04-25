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

type Assert struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   uint64 `json:"timestamp"`
	From        string `json:"from" gencodec:"required"`
	Nonce       uint64 `json:"nonce"`
	To          string `json:"to" gencodec:"required"`
	Amount      uint64 `json:"amount"`
	Types       string `json:"types"`
	Logindex    uint   `json:"logindex"`
	Txhash      string `json:"txhash" gencodec:"required"`
}
