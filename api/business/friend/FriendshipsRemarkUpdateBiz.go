package friend

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// FriendshipsRemarkUpdateBizString : 更新关注人备注
// http://open.weibo.com/wiki/C/2/friendships/remark/update/biz
func FriendshipsRemarkUpdateBizString(param apiParam.FriendshipsRemarkUpdateBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.FriendshipsRemarkUpdateBizURL, URLParam)
	return apiHelper.APIPost(apiURL.FriendshipsRemarkUpdateBizURL, URLParam)
}

// FriendshipsRemarkUpdateBizObject : 以对象的方式返回
func FriendshipsRemarkUpdateBizObject(param apiParam.FriendshipsRemarkUpdateBizParam, accessToken string) (*apiResp.FriendshipsRemarkUpdateBizResp, *apiResp.ErrorResp) {
	body, error := FriendshipsRemarkUpdateBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.FriendshipsRemarkUpdateBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "FriendshipsRemarkUpdateBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
