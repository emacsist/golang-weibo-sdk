package friend

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// FriendshipsCreateBizString : 关注某用户
// http://open.weibo.com/wiki/C/2/friendships/create/biz
func FriendshipsCreateBizString(param apiParam.FriendshipsCreateBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.FriendshipsCreateBizURL, URLParam)
	return apiHelper.APIPost(apiURL.FriendshipsCreateBizURL, URLParam)
}

// FriendshipsCreateBizObject : 以对象的方式返回
func FriendshipsCreateBizObject(param apiParam.FriendshipsCreateBizParam, accessToken string) (*apiResp.FriendshipsCreateBizResp, *apiResp.ErrorResp) {
	body, error := FriendshipsCreateBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.FriendshipsCreateBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "FriendshipsCreateBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
