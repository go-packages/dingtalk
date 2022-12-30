package dingtalk

import (
	"net/http"

	"github.com/go-packages/dingtalk/v2/constant"
	"github.com/go-packages/dingtalk/v2/request"
	"github.com/go-packages/dingtalk/v2/response"
)

// GetUserAccessToken 获取用户token
func (ding *DingTalk) GetUserAccessToken(grantType string, authCode string, refreshToken string) (req response.GetUserAccessToken, err error) {
	return req, ding.Request(
		http.MethodPost,
		constant.GetUserAccessTokenKey,
		nil,
		request.NewGetUserAccessToken(grantType, authCode, refreshToken, ding.key, ding.secret),
		&req,
	)
}
