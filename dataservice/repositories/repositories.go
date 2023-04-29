package dataservice

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo
	TokenRepo
}

func GetRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo{db},
		TokenRepo{db},
	}
}
