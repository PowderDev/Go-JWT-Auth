package config

import (
	"database/sql"

	handlers "auth/api/handlers"
	services "auth/api/services"
	"auth/dataservice/models"
	dataservice "auth/dataservice/repositories"
	"gorm.io/gorm"
)

type Config struct {
	API API
}

func SetupConfig(dbConn *sql.DB, db *gorm.DB) *Config {
	m := models.New(dbConn)
	repos := dataservice.GetRepositories(db)
	s := services.GetServices(repos)

	h := handlers.GetHandlers(s)

	ds := &DataService{dbConn, m, repos}

	api := API{
		h,
		ds,
		s,
	}

	return &Config{api}
}
