package friend

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// FriendshipsFollowersDestroyBizString : 移除当前登录用户的粉丝
// http://open.weibo.com/wiki/C/2/friendships/destroy/biz
func FriendshipsFollowersDestroyBizString(param apiParam.FriendshipsFollowersDestroyBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.FriendshipsFollowersDestroyBizURL, URLParam)
	return apiHelper.APIPost(apiURL.FriendshipsFollowersDestroyBizURL, URLParam)
}

// FriendshipsFollowersDestroyBizObject : 以对象的方式返回
func FriendshipsFollowersDestroyBizObject(param apiParam.FriendshipsFollowersDestroyBizParam, accessToken string) (*apiResp.FriendshipsFollowersDestroyBizResp, *apiResp.ErrorResp) {
	body, error := FriendshipsFollowersDestroyBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.FriendshipsFollowersDestroyBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "FriendshipsFollowersDestroyBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
