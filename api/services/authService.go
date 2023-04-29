package api

import (
	"net/http"

	apiExceptions "auth/api/exceptions"
	"auth/api/services/dtos"
	"auth/dataservice/models"
	dataservice "auth/dataservice/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repos *dataservice.Repositories
}

func (s AuthService) Authenticate(dto *dtos.AuthInput) (*models.User, *apiExceptions.HTTPError) {
	candidate, err := s.Repos.UserRepo.GetByEmail(dto.Email)

	if err != nil {
		return nil, apiExceptions.Error500(err.Error())
	} else if candidate != nil {
		return nil, &apiExceptions.HTTPError{
			Message:    "User with this email already exists",
			StatusCode: http.StatusBadRequest,
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := models.User{
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	userErr := s.Repos.UserRepo.CreateUser(&user)

	if userErr != nil {
		return nil, apiExceptions.Error500(err.Error())
	}

	return &user, nil
}

func (s AuthService) Login(dto *dtos.AuthInput) (*models.User, *apiExceptions.HTTPError) {
	candidate, err := s.Repos.UserRepo.GetByEmail(dto.Email)

	if err != nil {
		return nil, apiExceptions.Error500(err.Error())
	} else if candidate == nil {
		return nil, &apiExceptions.HTTPError{
			Message:    "Provided credentials are invalid",
			StatusCode: http.StatusBadRequest,
		}
	}

	matchErr := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(dto.Password))

	if matchErr != nil {
		return nil, &apiExceptions.HTTPError{
			Message:    "Provided credentials are invalid",
			StatusCode: http.StatusBadRequest,
		}
	}

	return candidate, nil
}
