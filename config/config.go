package config

import (
	"os"
	"strconv"
)

type Config struct {
	Service  Service
	Database Database
}

type Service struct {
	Name    string
	Port    int
	Version string
}

type Database struct {
	Username string
	Password string
	Port     int
	DBName   string
	Driver   string
	Host     string
}

func (c *Config) Init() {
	dbPort, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	svcPort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))

	c.Service = Service{
		Name:    os.Getenv("SERVICE_NAME"),
		Port:    svcPort,
		Version: os.Getenv("SERVICE_VERSION"),
	}

	c.Database = Database{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     dbPort,
		DBName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
	}
}
