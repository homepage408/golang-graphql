package connection

import (
	"database/sql"
	"fmt"
	"log"

	conf "github.com/golang-graphql/config"
)

type postgresConnection struct {
	db *sql.DB
}

type Connection interface {
	DB() *sql.DB
}

func NewPostgresConnection(db *sql.DB, conf *conf.Config) Connection {
	return &postgresConnection{
		db: setConnection(conf),
	}
}

func (connection *postgresConnection) DB() *sql.DB {
	return connection.db
}

func setConnection(conf *conf.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.DBName)

	db, err := sql.Open("Postgres", psqlInfo)
	if err != nil {
		log.Println("Error establishing connection")
		panic(err)
	}

	return db
}
