package database

import (
	"conversion-currency/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionString)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
