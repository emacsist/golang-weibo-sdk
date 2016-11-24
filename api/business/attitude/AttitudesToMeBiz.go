package attitude

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// AttitudesToMeBizString : 收到的赞列表
// http://open.weibo.com/wiki/C/2/attitudes/to_me/biz
func AttitudesToMeBizString(param apiParam.AttitudesToMeBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.AttitudesToMeBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// AttitudesToMeBizObject : 以对象的方式返回
func AttitudesToMeBizObject(param apiParam.AttitudesToMeBizParam, accessToken string) (*apiResp.AttitudesToMeBizResp, *apiResp.ErrorResp) {
	body, error := AttitudesToMeBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.AttitudesToMeBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "AttitudesToMeBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
