package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver string
	Url string
}

func NewMySQL() *sql.DB {
	s, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	return s
}