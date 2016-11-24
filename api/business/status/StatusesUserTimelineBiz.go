package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesUserTimelineBizString : 获取当前登录用户发布的微博.
// http://open.weibo.com/wiki/C/2/statuses/user_timeline/biz
func StatusesUserTimelineBizString(param apiParam.StatusesUserTimelineBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesUserTimelineBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesUserTimelineBizObject : 以对象的方式返回
func StatusesUserTimelineBizObject(param apiParam.StatusesUserTimelineBizParam, accessToken string) (*apiResp.StatusesUserTimelineBizResp, *apiResp.ErrorResp) {
	body, error := StatusesUserTimelineBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesUserTimelineBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesUserTimelineBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
