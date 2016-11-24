package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesMentionsShieldBizString : 转发一条微博信息.
// http://open.weibo.com/wiki/C/2/statuses/repost/biz
func StatusesMentionsShieldBizString(param apiParam.StatusesMentionsShieldBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.StatusesMentionsShieldBizURL, URLParam)
	return apiHelper.APIPost(apiURL.StatusesMentionsShieldBizURL, URLParam)
}

// StatusesMentionsShieldBizObject : 以对象的方式返回
func StatusesMentionsShieldBizObject(param apiParam.StatusesMentionsShieldBizParam, accessToken string) (*apiResp.StatusesMentionsShieldBizResp, *apiResp.ErrorResp) {
	body, error := StatusesMentionsShieldBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesMentionsShieldBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesMentionsShieldBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
