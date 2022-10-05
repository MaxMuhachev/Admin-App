package app

import (
	"content/src/config"
	"content/src/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connect struct {
	Mysql *sqlx.DB
}

func NewConnect() *Connect {
	db, err := sqlx.Open("mysql", utils.FormatConnect(config.GetDnsConfig()))

	if err != nil {
		panic(err.Error())
	}

	return &Connect{db}
}

func CloseConnect(s *Connect) {
	if err := s.Mysql.Close(); err != nil {
		_ = fmt.Errorf("could not close connection MySQL:%v", err)
	}
}
