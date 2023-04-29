package api

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	apiExceptions "auth/api/exceptions"
	"auth/api/services/dtos"
	"auth/dataservice/models"
	dataservice "auth/dataservice/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType int

const (
	AccessTokenType  TokenType = 1
	RefreshTokenType TokenType = 2
)

type TokenService struct {
	Repos *dataservice.Repositories
}

func (s TokenService) GenerateTokens(p *dtos.JWTPayload) (
	*dtos.JWTTokens,
	*apiExceptions.HTTPError,
) {
	var refreshSecret = os.Getenv("JWT_REFRESH_SECRET")
	var accessSecret = os.Getenv("JWT_ACCESS_SECRET")

	accessToken, err := getDefaultAccessToken(*p).SignedString([]byte(accessSecret))
	if err != nil {
		return nil, apiExceptions.Error500(err.Error())
	}

	refreshToken, err := getDefaultRefreshToken(*p).SignedString([]byte(refreshSecret))

	if err != nil {
		return nil, apiExceptions.Error500(err.Error())
	}

	return &dtos.JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s TokenService) Save(token models.Token) *apiExceptions.HTTPError {
	err := s.Repos.TokenRepo.SaveToken(&token)

	if err != nil {
		return apiExceptions.Error400(err.Error())
	}

	return nil
}

func (s TokenService) Delete(refreshToken string) *apiExceptions.HTTPError {
	tokenToDelete := &models.Token{RefreshToken: refreshToken}

	err := s.Repos.TokenRepo.DeleteToken(tokenToDelete)

	if err != nil {
		return apiExceptions.Error400(err.Error())
	}

	return nil
}

func (s TokenService) Validate(t string, tokenType TokenType) (
	*dtos.JWTPayload,
	*apiExceptions.HTTPError,
) {
	var token *jwt.Token
	var err error

	if tokenType == AccessTokenType {
		token, err = parseAccessToken(t)
	} else {
		token, err = parseRefreshToken(t)
	}

	claims := token.Claims.(jwt.MapClaims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, apiExceptions.Error401()
		}

		return nil, apiExceptions.Error500(err.Error())
	}
	if !token.Valid {
		return nil, apiExceptions.Error401()
	}

	sub := int(claims["sub"].(float64))

	return &dtos.JWTPayload{Sub: sub}, nil
}

func (s TokenService) GetNewTokens(payload *dtos.JWTPayload) (
	*dtos.JWTTokens,
	*apiExceptions.HTTPError,
) {
	tokens, err := s.GenerateTokens(payload)

	if err != nil {
		return nil, err
	}

	refreshToken := models.Token{UserID: payload.Sub, RefreshToken: tokens.RefreshToken}
	saveErr := s.Save(refreshToken)

	if saveErr != nil {
		return nil, saveErr
	}

	return tokens, nil
}

func (s TokenService) Refresh(refreshToken string) (*dtos.JWTTokens, *apiExceptions.HTTPError) {
	payload, err := s.Validate(refreshToken, RefreshTokenType)

	if err != nil {
		return nil, err
	}

	_, dbErr := s.Repos.TokenRepo.GetTokenByUserID(payload.Sub)

	if dbErr != nil {
		return nil, apiExceptions.Error401()
	}

	tokens, tErr := s.GetNewTokens(payload)

	if tErr != nil {
		return nil, tErr
	}

	return tokens, nil
}

func parseRefreshToken(refreshToken string) (*jwt.Token, error) {
	var refreshSecret = os.Getenv("JWT_REFRESH_SECRET")

	token, err := jwt.Parse(
		refreshToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(refreshSecret), nil
		},
	)

	return token, err
}

func parseAccessToken(accessToken string) (*jwt.Token, error) {
	var accessSecret = os.Getenv("JWT_ACCESS_SECRET")

	token, err := jwt.Parse(
		accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil
		},
	)

	return token, err
}

func getDefaultAccessToken(payload dtos.JWTPayload) *jwt.Token {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": json.Number(
				strconv.FormatInt(
					time.Now().Add(time.Hour*time.Duration(1)).Unix(),
					10,
				),
			),
			"iat": json.Number(
				strconv.FormatInt(
					time.Now().Unix(),
					10,
				),
			),
			"sub": payload.Sub,
		},
	)
	return token
}

func getDefaultRefreshToken(payload dtos.JWTPayload) *jwt.Token {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": json.Number(
				strconv.FormatInt(
					time.Now().Add(7*24*time.Hour*time.Duration(1)).Unix(),
					10,
				),
			),
			"iat": json.Number(
				strconv.FormatInt(
					time.Now().Unix(),
					10,
				),
			),
			"sub": payload.Sub,
		},
	)

	return token
}
