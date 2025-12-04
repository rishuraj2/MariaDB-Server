package utility

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CheckDbConnection(db *sql.DB) error {
	var err error

	for dbConnectionTries := 0; dbConnectionTries < 60; dbConnectionTries++ {
		if err = db.Ping(); err == nil {
			return nil
		}

		log.Printf("Database is not ready, waiting.... attempt %d/60", dbConnectionTries+1)
		time.Sleep(1 * time.Second)
	}

	return err
}