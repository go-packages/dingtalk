package response

type GetUserAccessToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireIn     int    `json:"expireIn"`
	CorpId       string `json:"coreId"`
}

func (*GetUserAccessToken) CheckError() error {

	return nil
}
