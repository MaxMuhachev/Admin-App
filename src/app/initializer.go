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

var Conn *Connect

func NewConnect() {
	db, err := sqlx.Open("mysql", utils.FormatConnect(config.GetDnsConfig()))

	if err != nil {
		panic(err.Error())
	}
	Conn = &Connect{db}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	Conn = &Connect{db}

}

//func  CloseConnect() {
//	defer func(Mysql *sqlx.DB) {
//		err := Mysql.Close()
//		if err != nil {
//			log.Println(fmt.Errorf("could not close connection MySQL:%v", err))
//		}
//	}(c.Mysql)
//}
