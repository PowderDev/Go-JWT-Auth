package api

import (
	dataservice "auth/dataservice/repositories"
)

type Services struct {
	AuthService
	TokenService
}

func GetServices(repos *dataservice.Repositories) *Services {
	return &Services{
		AuthService{repos},
		TokenService{repos},
	}
}
