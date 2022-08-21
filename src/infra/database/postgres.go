package database

import (
	"api-unico/infra/logger"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgreSQLConnection struct {
	Db *sqlx.DB
}

func NewPostgreSQLConnection(connectionString string) postgreSQLConnection {
	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		db, err = sqlx.Connect("postgres", connectionString)
		if err != nil {
			logger.LogError("NewPostgreSQLConnection", fmt.Sprintf("retrying connection in database %d", i+1))
		} else {
			break
		}
	}

	if db == nil {
		logger.LogFatal("NewPostgreSQLConnection", "unable to connect at database")
	}
	return postgreSQLConnection{
		Db: db,
	}
}
