package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesUploadBizString : 上传图片并发布一条微博
// http://open.weibo.com/wiki/C/2/statuses/upload/biz
func StatusesUploadBizString(param apiParam.StatusesUploadBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	req, err := apiHelper.BuildPostMultiPartQuery(&param, apiURL.StatusesUploadBizURL, accessToken, "pic", param.Pic)
	if err != nil {
		return "", &apiResp.ErrorResp{Error: err.Error(), ErrorCode: apiResp.IOErrorCode, Request: "StatusesUploadBizString"}
	}

	logrus.Infof("api invoke [url=%v, param=%v, token=%v]\n", apiURL.StatusesUploadBizURL, param.String(), accessToken)
	return apiHelper.APIPostMultiPart(req)
}

// StatusesUploadBizObject : 以对象的方式返回
func StatusesUploadBizObject(param apiParam.StatusesUploadBizParam, accessToken string) (*apiResp.StatusesUploadBizResp, *apiResp.ErrorResp) {
	body, error := StatusesUploadBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesUploadBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesUploadBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
