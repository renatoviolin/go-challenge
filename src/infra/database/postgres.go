package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type postgreSQLConnection struct {
	Db *sqlx.DB
}

func NewPostgreSQLConnection(connectionString string) postgreSQLConnection {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	return postgreSQLConnection{
		Db: db,
	}
}
