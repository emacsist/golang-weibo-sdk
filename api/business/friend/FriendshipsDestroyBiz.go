package friend

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// FriendshipsDestroyBizString : 取消关注某用户
// http://open.weibo.com/wiki/C/2/friendships/destroy/biz
func FriendshipsDestroyBizString(param apiParam.FriendshipsDestroyBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.FriendshipsDestroyBizURL, URLParam)
	return apiHelper.APIPost(apiURL.FriendshipsDestroyBizURL, URLParam)
}

// FriendshipsDestroyBizObject : 以对象的方式返回
func FriendshipsDestroyBizObject(param apiParam.FriendshipsDestroyBizParam, accessToken string) (*apiResp.FriendshipsDestroyBizResp, *apiResp.ErrorResp) {
	body, error := FriendshipsDestroyBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.FriendshipsDestroyBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "FriendshipsDestroyBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
