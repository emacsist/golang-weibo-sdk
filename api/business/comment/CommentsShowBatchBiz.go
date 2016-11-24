package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsShowBatchBizString : 批量获取评论内容.
// http://open.weibo.com/wiki/C/2/comments/show_batch/biz
func CommentsShowBatchBizString(param apiParam.CommentsShowBatchBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsShowBatchBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsShowBatchBizObject : 以对象的方式返回
func CommentsShowBatchBizObject(param apiParam.CommentsShowBatchBizParam, accessToken string) (*[]apiResp.CommentsShowBatchBizResp, *apiResp.ErrorResp) {
	body, error := CommentsShowBatchBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp []apiResp.CommentsShowBatchBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsShowBatchBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
