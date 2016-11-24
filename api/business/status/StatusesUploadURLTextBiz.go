package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesUploadURLTextBizString : 发布一条微博同时指定上传的图片或图片url
// http://open.weibo.com/wiki/C/2/statuses/upload_url_text/biz
func StatusesUploadURLTextBizString(param apiParam.StatusesUploadURLTextBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)

	logrus.Infof("api invoke [url=%v, param=%v, token=%v]\n", apiURL.StatusesUploadURLTextBizURL, param, accessToken)
	return apiHelper.APIPost(apiURL.StatusesUploadURLTextBizURL, URLParam)
}

// StatusesUploadURLTextBizObject : 以对象的方式返回
func StatusesUploadURLTextBizObject(param apiParam.StatusesUploadURLTextBizParam, accessToken string) (*apiResp.StatusesUploadURLTextBizResp, *apiResp.ErrorResp) {
	body, error := StatusesUploadURLTextBizString(param, accessToken)
	if error != nil {
		return nil, error
	}
	var statusesResp apiResp.StatusesUploadURLTextBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesUploadURLTextBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
