package repositories

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectToDatabase() (Db *sqlx.DB) {
	databaseUrl := "host=localhost port=5432 user=postgres password=934007717 dbname=courses sslmode=disable"

	var err error
	Db, err = sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		log.Fatalln("Connection error", err)
	}

	fmt.Println("Success to Connect!")
	return Db
}
