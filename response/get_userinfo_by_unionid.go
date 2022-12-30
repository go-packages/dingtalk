package response

type GetUserInfoByUnionId struct {
	UnionId   string `json:"unionId"`
	Nick      string `json:"nick"`
	AvatarUrl string `json:"avatarUrl"`
	Mobile    string `json:"mobile"`
	OpenId    string `json:"openId"`
	Email     string `json:"email"`
	StateCode string `json:"stateCode"`
}

func (*GetUserInfoByUnionId) CheckError() error {

	return nil
}
