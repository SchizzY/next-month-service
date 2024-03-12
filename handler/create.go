package handler

import (
	"database/sql"
)

func CreateTest(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test (id serial PRIMARY KEY, name VARCHAR(50))")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO test (name) VALUES ('test')")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO test (name) VALUES ('test2')")

	if err != nil {
		panic(err)
	}
}

func ClearTest(db *sql.DB) {
	_, err := db.Exec("DROP TABLE test")
	if err != nil {
		panic(err)
	}
}
