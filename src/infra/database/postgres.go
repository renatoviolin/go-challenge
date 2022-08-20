package database

import (
	"api-unico/infra/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgreSQLConnection struct {
	Db *sqlx.DB
}

func NewPostgreSQLConnection(connectionString string) postgreSQLConnection {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		logger.Logger.Fatal().Msg(err.Error())
	}

	return postgreSQLConnection{
		Db: db,
	}
}
