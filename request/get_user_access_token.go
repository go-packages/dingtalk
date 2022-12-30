package request

type GetUserAccessToken struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	AuthCode     string `json:"code,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	GrantType    string `json:"grantType" validate:"required"`
}

func NewGetUserAccessToken(grantType string, authCode string, refreshToken string, clientId string, clientSecret string) *GetUserAccessToken {
	return &GetUserAccessToken{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		AuthCode:     authCode,
		RefreshToken: refreshToken,
		GrantType:    grantType,
	}
}
