package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsShowBizString : 返回一条微博的全部评论列表.
// http://open.weibo.com/wiki/C/2/comments/show/biz
func CommentsShowBizString(param apiParam.CommentsShowBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsShowBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsShowBizObject : 以对象的方式返回
func CommentsShowBizObject(param apiParam.CommentsShowBizParam, accessToken string) (*apiResp.CommentsShowBizResp, *apiResp.ErrorResp) {
	body, error := CommentsShowBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsShowBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsShowBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
