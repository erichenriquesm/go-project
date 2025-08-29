package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"os"
	"strconv"
)


func ConnectDB() (*sql.DB, error) {
	host      := 	os.Getenv("DB_HOST")
	port, err := 	strconv.Atoi(os.Getenv("DB_PORT"))
	user      := 	os.Getenv("DB_USER")
	password  :=	os.Getenv("DB_PASSWORD")
	dbName    := 	os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Printf("System connected with DB")

	return db, nil
}
