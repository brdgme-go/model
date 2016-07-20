// +build integration

package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	pgOpts := os.Getenv("PG_OPTS")
	if pgOpts == "" {
		panic("the PG_OPTS environment variable must be set")
	}
	var err error
	db, err = sql.Open("postgres", pgOpts)
	if err != nil {
		panic(fmt.Sprintf("unable to connect to database, %s", err))
	}
}
