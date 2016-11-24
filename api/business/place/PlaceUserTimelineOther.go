package place

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// PlaceUserTimelineOtherString : 批量获取用户标签
// http://open.weibo.com/wiki/C/2/tags/tags_batch/other
func PlaceUserTimelineOtherString(param apiParam.PlaceUserTimelineOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.PlaceUserTimelineOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// PlaceUserTimelineOtherObject : 以对象的方式返回
func PlaceUserTimelineOtherObject(param apiParam.PlaceUserTimelineOtherParam, accessToken string) (*apiResp.PlaceUserTimelineOtherResp, *apiResp.ErrorResp) {
	body, error := PlaceUserTimelineOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.PlaceUserTimelineOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "PlaceUserTimelineOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
