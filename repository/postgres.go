package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const audioTable = "audio"

type Config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
	SSLMode  string
}

func NewPostgresDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname =%s host=%s port=%s sslmode=%s", c.User, c.Password, c.DBName, c.Host, c.Port, c.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, err
}
