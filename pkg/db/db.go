package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Connection sqlx.DB

func NewConnection() (*Connection, error){
	db, err := sqlx.Connect("mysql", os.Getenv("DB_NAME")+"@'localhost'")
	if err != nil {
		return nil, err
	}
	return (*Connection)(db), nil
}

