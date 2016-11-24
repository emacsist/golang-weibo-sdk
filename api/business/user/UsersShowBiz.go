package user

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// UsersShowBizString : 获取当前登录用户信息
// http://open.weibo.com/wiki/C/2/users/show/biz
func UsersShowBizString(accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(nil, apiURL.UsersShowBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// UsersShowBizObject : 以对象的方式返回
func UsersShowBizObject(accessToken string) (*apiResp.UsersShowBizResp, *apiResp.ErrorResp) {
	body, error := UsersShowBizString(accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.UsersShowBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "UsersShowBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
