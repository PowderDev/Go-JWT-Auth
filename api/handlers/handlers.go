package api

import (
	services "auth/api/services"
)

type Handlers struct {
	AuthHandler
	MiddlewareHandler
}

func GetHandlers(s *services.Services) *Handlers {
	return &Handlers{
		AuthHandler{s},
		MiddlewareHandler{s},
	}
}
