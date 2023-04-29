package dtos

type JWTTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}

type JWTPayload struct {
	Sub int
}
