package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesFilterCreateBizString : 删除一条微博
// http://open.weibo.com/wiki/C/2/statuses/destroy/biz
func StatusesFilterCreateBizString(param apiParam.StatusesFilterCreateBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.StatusesFilterCreateBizURL, URLParam)
	return apiHelper.APIPost(apiURL.StatusesFilterCreateBizURL, URLParam)
}

// StatusesFilterCreateBizObject : 以对象的方式返回
func StatusesFilterCreateBizObject(param apiParam.StatusesFilterCreateBizParam, accessToken string) (*apiResp.StatusesFilterCreateBizResp, *apiResp.ErrorResp) {
	body, error := StatusesFilterCreateBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.StatusesFilterCreateBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesFilterCreateBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
