package pkg

import (
	"database/sql"
	"log"
	"net/url"
)

func InitDB(dbServer string) *sql.DB {
	u, err := url.Parse(dbServer)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
