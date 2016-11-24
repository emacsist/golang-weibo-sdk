package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesPublicTimelineBizString : 返回一条微博的全部转发微博列表.
// http://open.weibo.com/wiki/C/2/statuses/public_timeline/biz
func StatusesPublicTimelineBizString(param apiParam.StatusesPublicTimelineBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesPublicTimelineBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesPublicTimelineBizObject : 以对象的方式返回
func StatusesPublicTimelineBizObject(param apiParam.StatusesPublicTimelineBizParam, accessToken string) (*apiResp.StatusesPublicTimelineBizResp, *apiResp.ErrorResp) {
	body, error := StatusesPublicTimelineBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesPublicTimelineBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesPublicTimelineBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
