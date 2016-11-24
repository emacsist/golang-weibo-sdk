package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesRepostTimelineBizString : 批量获取用户个人微博列表.
// http://open.weibo.com/wiki/C/2/statuses/user_timeline_batch
func StatusesRepostTimelineBizString(param apiParam.StatusesRepostTimelineBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesRepostTimelineBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesRepostTimelineBizObject : 以对象的方式返回
func StatusesRepostTimelineBizObject(param apiParam.StatusesRepostTimelineBizParam, accessToken string) (*apiResp.StatusesRepostTimelineBizResp, *apiResp.ErrorResp) {
	body, error := StatusesRepostTimelineBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesRepostTimelineBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		logrus.Warnf("response body is not json => %v\n", string(body))
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesRepostTimelineBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
