package config

import (
	"database/sql"

	handlers "auth/api/handlers"
	api "auth/api/services"
	"auth/dataservice/models"
	dataservice "auth/dataservice/repositories"
)

type API struct {
	Handlers *handlers.Handlers
	DB       *DataService
	Services *api.Services
}

type DataService struct {
	DB           *sql.DB
	Models       models.Models
	Repositories *dataservice.Repositories
}
