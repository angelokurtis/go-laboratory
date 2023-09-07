package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
