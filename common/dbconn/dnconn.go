package dbconn

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"main/common/tabletypes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConFig struct {
	Addr            string
	Port            int
	Db              string
	Username        string
	Password        string
	MaxIdealConn    int
	MaxOpenConn     int
	ConnMaxLifetime int
}

var MysqlCon = MysqlConFig{
	"127.0.0.1",
	3306,
	"eventlog",
	"root",
	"123456",
	10,
	256,
	600,
}

func Buildconnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		MysqlCon.Username, MysqlCon.Password, MysqlCon.Addr, MysqlCon.Port, MysqlCon.Db, "10s")
	//mysql connection
	dba, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Connect database error:%s\n", err)
	}

	dba.AutoMigrate(tabletypes.AddressKey{})
	dba.AutoMigrate(tabletypes.Contracts{})
	dba.AutoMigrate(tabletypes.Assert{})
	return dba
}
