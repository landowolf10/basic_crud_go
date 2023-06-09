package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "landowolf10"
	dbName := "notes"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
