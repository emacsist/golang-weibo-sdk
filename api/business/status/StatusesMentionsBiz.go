package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesMentionsBizString : 获取@当前登录用户的微博列表
// http://open.weibo.com/wiki/C/2/statuses/mentions/biz
func StatusesMentionsBizString(param apiParam.StatusesMentionsBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesMentionsBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesMentionsBizObject : 以对象的方式返回
func StatusesMentionsBizObject(param apiParam.StatusesMentionsBizParam, accessToken string) (*apiResp.StatusesMentionsBizResp, *apiResp.ErrorResp) {
	body, error := StatusesMentionsBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesMentionsBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesMentionsBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
