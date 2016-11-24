package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesCountBizString : 批量获取指定微博的转发数评论数喜欢数
// http://open.weibo.com/wiki/C/2/statuses/count/biz
func StatusesCountBizString(param apiParam.StatusesCountBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesCountBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesCountBizObject : 以对象的方式返回
func StatusesCountBizObject(param apiParam.StatusesCountBizParam, accessToken string) (*[]apiResp.StatusesCountBizResp, *apiResp.ErrorResp) {
	body, error := StatusesCountBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp []apiResp.StatusesCountBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesCountBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
