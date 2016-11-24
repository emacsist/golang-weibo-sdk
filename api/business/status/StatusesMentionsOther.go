package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesMentionsOtherString : 获取@某人的微博
// http://open.weibo.com/wiki/C/2/statuses/mentions/other
func StatusesMentionsOtherString(param apiParam.StatusesMentionsOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesMentionsOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesMentionsOtherObject : 以对象的方式返回
func StatusesMentionsOtherObject(param apiParam.StatusesMentionsOtherParam, accessToken string) (*apiResp.StatusesMentionsOtherResp, *apiResp.ErrorResp) {
	body, error := StatusesMentionsOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesMentionsOtherResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesMentionsOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
