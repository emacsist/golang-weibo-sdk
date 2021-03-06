package friend

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// FriendshipsFriendsBizString : 获取当前登录用户的关注列表
// http://open.weibo.com/wiki/C/2/friendships/friends/biz
func FriendshipsFriendsBizString(param apiParam.FriendshipsFriendsBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.FriendshipsFriendsBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// FriendshipsFriendsBizObject : 以对象的方式返回
func FriendshipsFriendsBizObject(param apiParam.FriendshipsFriendsBizParam, accessToken string) (*apiResp.FriendshipsFriendsBizResp, *apiResp.ErrorResp) {
	body, error := FriendshipsFriendsBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.FriendshipsFriendsBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "FriendshipsFriendsBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
