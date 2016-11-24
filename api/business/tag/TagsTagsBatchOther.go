package tag

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// TagsTagsBatchOtherString : 批量获取用户标签
// http://open.weibo.com/wiki/C/2/tags/tags_batch/other
func TagsTagsBatchOtherString(param apiParam.TagsTagsBatchOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.TagsTagsBatchOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// TagsTagsBatchOtherObject : 以对象的方式返回
func TagsTagsBatchOtherObject(param apiParam.TagsTagsBatchOtherParam, accessToken string) (*[]apiResp.TagsTagsBatchOtherResp, *apiResp.ErrorResp) {
	body, error := TagsTagsBatchOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp []apiResp.TagsTagsBatchOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "TagsTagsBatchOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
