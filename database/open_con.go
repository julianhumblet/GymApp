package database

import (
	"database/sql"
	"fmt"
)

func OpenDBCon(dbConCreds string) (*sql.DB, error) {
	dbCon, err := sql.Open("mysql", dbConCreds)
	if err != nil {
		return nil, fmt.Errorf("error making database connection: %s", err)
	}

	err = dbCon.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %s", err)
	}

	return dbCon, nil
}
