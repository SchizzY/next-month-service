package main

import (
	"database/sql"
	"log"
	"next-month/service/helper"
	"next-month/service/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	helper.SetEnvironment()
	app := gin.Default()

	connStr := helper.NeonConn()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}

	routes.Get(app, db)
	app.Run()
}
