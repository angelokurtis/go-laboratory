package mysql

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func NewDB() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, func() {}, errors.WithStack(err)
	}

	cleanup := func() {
		if cerr := db.Close(); cerr != nil {
			slog.Warn("an error happened while cleaning up",
				slog.String("component", fmt.Sprintf("%T", db)),
				slog.String("err", cerr.Error()),
			)
		}
	}

	return db, cleanup, nil
}
