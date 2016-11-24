package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesRepostBizString : 转发一条微博信息.
// http://open.weibo.com/wiki/C/2/statuses/repost/biz
func StatusesRepostBizString(param apiParam.StatusesRepostBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.StatusesRepostBizURL, URLParam)
	return apiHelper.APIPost(apiURL.StatusesRepostBizURL, URLParam)
}

// StatusesRepostBizObject : 以对象的方式返回
func StatusesRepostBizObject(param apiParam.StatusesRepostBizParam, accessToken string) (*apiResp.StatusesRepostBizResp, *apiResp.ErrorResp) {
	body, error := StatusesRepostBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesRepostBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesRepostBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
