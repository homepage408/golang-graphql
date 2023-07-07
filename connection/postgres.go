package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type postgresConnection struct {
	db *sql.DB
}

type Connection interface {
	DB() *sql.DB
}

func NewPostgresConnection() Connection {
	return &postgresConnection{
		db: setConnection(),
	}
}

func (connection *postgresConnection) DB() *sql.DB {
	return connection.db
}

func setConnection() *sql.DB {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error establishing connection")
		panic(err)
	}

	return db
}
