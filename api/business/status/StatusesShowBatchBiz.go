package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesShowBatchBizString : 根据微博ID批量获取微博信息
// http://open.weibo.com/wiki/C/2/statuses/show_batch/biz
func StatusesShowBatchBizString(param apiParam.StatusesShowBatchBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesShowBatchBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesShowBatchBizObject : 以对象的方式返回
func StatusesShowBatchBizObject(param apiParam.StatusesShowBatchBizParam, accessToken string) (*apiResp.StatusesShowBatchBizResp, *apiResp.ErrorResp) {
	body, error := StatusesShowBatchBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesShowBatchBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		logrus.Warnf("response body is not json => %v\n", string(body))
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesShowBatchBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
