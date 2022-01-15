package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection sqlx.DB

func NewConnection() (*Connection, error) {
	db, err := sqlx.Connect("mysql",
		"sorokin:1234@/fin_manager")
	if err != nil {
		return nil, err
	}
	return (*Connection)(db), nil
}
