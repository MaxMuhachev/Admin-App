package app

import (
	"content/src/config"
	"content/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connect struct {
	Mysql *sqlx.DB
}

var c *Connect

func NewConnect() *Connect {
	return c
}
func CreateConnect() *Connect {
	db, err := sqlx.Open("mysql", utils.FormatConnect(config.GetDnsConfig()))
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	if err != nil {
		panic(err.Error())
	}
	c = &Connect{db}

	return c
}

//func CloseConnect() {
//	defer func(Mysql *sqlx.DB) {
//		err := Mysql.Close()
//		if err != nil {
//			log.Println(err)
//		}
//		c = nil
//	}(c.Mysql)
//}
