package user

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// UsersCountsBatchOtherString : 批量获取用户标签
// http://open.weibo.com/wiki/C/2/tags/tags_batch/other
func UsersCountsBatchOtherString(param apiParam.UsersCountsBatchOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.UsersCountsBatchOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// UsersCountsBatchOtherObject : 以对象的方式返回
func UsersCountsBatchOtherObject(param apiParam.UsersCountsBatchOtherParam, accessToken string) (*[]apiResp.UsersCountsBatchOtherResp, *apiResp.ErrorResp) {
	body, error := UsersCountsBatchOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp []apiResp.UsersCountsBatchOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "UsersCountsBatchOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
