package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("user not exist")

type Dao struct {
	db *sql.DB
}

func NewDao(url string) *Dao {
	d := new(Dao)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	d.db = db

	return d
}

func (d *Dao) GetUserById(id string) (string, error) {
	var name string
	err := d.db.QueryRow("SELECT name FROM users WHERE id=?", id).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.Wrap(ErrNotFound, "")
		}

		return "", err
	}

	return name, nil
}