package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesDestroyBizString : 删除一条微博
// http://open.weibo.com/wiki/C/2/statuses/destroy/biz
func StatusesDestroyBizString(param apiParam.StatusesDestroyBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.StatusesDestroyBizURL, URLParam)
	return apiHelper.APIPost(apiURL.StatusesDestroyBizURL, URLParam)
}

// StatusesDestroyBizObject : 以对象的方式返回
func StatusesDestroyBizObject(param apiParam.StatusesDestroyBizParam, accessToken string) (*apiResp.StatusesDestroyBizResp, *apiResp.ErrorResp) {
	body, error := StatusesDestroyBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.StatusesDestroyBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesDestroyBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
