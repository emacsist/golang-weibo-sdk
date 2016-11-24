package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesRepostTimelineAllString : 返回一条微博的全部转发微博列表.
// http://open.weibo.com/wiki/C/2/statuses/repost_timeline/all
func StatusesRepostTimelineAllString(param apiParam.StatusesRepostTimelineAllParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesRepostTimelineAllURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesRepostTimelineAllObject : 以对象的方式返回
func StatusesRepostTimelineAllObject(param apiParam.StatusesRepostTimelineAllParam, accessToken string) (*apiResp.StatusesRepostTimelineAllResp, *apiResp.ErrorResp) {
	body, error := StatusesRepostTimelineAllString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesRepostTimelineAllResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesRepostTimelineAllObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
