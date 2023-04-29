package config

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*sql.DB, *gorm.DB, error) {
	var dbConn *sql.DB
	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		dbConn, err = db.DB()

		if i == 10 {
			return nil, nil, err
		}

		if err != nil {
			log.Printf("Could not connect to DB. Retry #%d\n", i)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	return dbConn, db, nil
}
