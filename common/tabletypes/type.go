package tabletypes

type AddressKey struct {
	ID        uint   `gorm:"primary_key"`
	Shuser    string `gorm:"shuser"`
	Shaddress string `json:"shaddress"`
	Shkey     string `json:"shkey"`
}

type Contracts struct {
	ID        uint   `gorm:"primary_key"`
	Shaddress string `json:"shaddress"`
	Shblock   uint64 `json:"shblock"`
}

type Assert struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Shtimestamp uint64 `json:"shtimestamp"`
	Shfrom      string `json:"shfrom" gencodec:"required"`
	Shnonce     uint64 `json:"shnonce"`
	Shto        string `json:"shto" gencodec:"required"`
	Shamount    uint64 `json:"shamount"`
	Shtypes     string `json:"shtypes"`
	Shlogindex  uint   `json:"shlogindex"`
	Txhash      string `json:"txhash" gencodec:"required"`
}
