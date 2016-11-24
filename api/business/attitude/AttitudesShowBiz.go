package attitude

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// AttitudesShowBizString : 根据微博ID返回某条微博的赞列表
// http://open.weibo.com/wiki/C/2/attitudes/show/biz
func AttitudesShowBizString(param apiParam.AttitudesShowBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.AttitudesShowBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// AttitudesShowBizObject : 以对象的方式返回
func AttitudesShowBizObject(param apiParam.AttitudesShowBizParam, accessToken string) (*apiResp.AttitudesShowBizResp, *apiResp.ErrorResp) {
	body, error := AttitudesShowBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.AttitudesShowBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "AttitudesShowBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
