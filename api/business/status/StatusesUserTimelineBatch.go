package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesUserTimelineBatchString : 批量获取用户个人微博列表.
// http://open.weibo.com/wiki/C/2/statuses/user_timeline_batch
func StatusesUserTimelineBatchString(param apiParam.StatusesUserTimelineBatchParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesUserTimelineBatchURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesUserTimelineBatchObject : 以对象的方式返回
func StatusesUserTimelineBatchObject(param apiParam.StatusesUserTimelineBatchParam, accessToken string) (*apiResp.StatusesUserTimelineBatchResp, *apiResp.ErrorResp) {
	body, error := StatusesUserTimelineBatchString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesUserTimelineBatchResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesUserTimelineBatchObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
